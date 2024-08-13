package middleware

import (
    "context"
    "go-task-api/models"
    "go-task-api/utils"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

        claims := &models.Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        var user models.User
        utils.DB.Where("username = ?", claims.Username).First(&user)

        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func GetAuthenticatedUser(r *http.Request) models.User {
    return r.Context().Value("user").(models.User)
}
