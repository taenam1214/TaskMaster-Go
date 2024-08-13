package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "go-task-api/models"
    "go-task-api/utils"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

var jwtKey = []byte("your_secret_key")

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

    user := models.User{Username: creds.Username, Password: string(hashedPassword)}

    if err := utils.DB.Create(&user).Error; err != nil {
        if gorm.ErrDuplicateKey == err {
            http.Error(w, "Username already taken", http.StatusBadRequest)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)

    var user models.User
    utils.DB.Where("username = ?", creds.Username).First(&user)

    if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(time.Hour * 24)
    claims := &models.Claims{
        Username: creds.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString(jwtKey)

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })
}

func Home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to the Task Manager API!"))
}
