Absolutely! Here's a comprehensive README for your project, covering all essential aspects:

---

# TaskMaster Go

**TaskMaster Go** is a lightweight and secure task management API built with Go. This API allows users to manage tasks efficiently, featuring user authentication with JWT (JSON Web Tokens) and CRUD (Create, Read, Update, Delete) operations for tasks. The API is backed by PostgreSQL for robust data persistence.

## Features

- **User Authentication**: Secure user registration and login with password hashing.
- **JWT-Based Authentication**: Secure API endpoints with JWT tokens.
- **Task Management**: Users can create, view, update, and delete tasks.
- **PostgreSQL Integration**: Persistent storage for user and task data.
- **RESTful API**: Follows REST principles for easy integration with front-end or other services.
- **Modular Design**: Clean and maintainable code structure with Go’s best practices.

## Tech Stack

- **Language**: Go
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **ORM**: GORM
- **Routing**: Gorilla Mux

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (1.16+ recommended)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/taenam1214/TaskMaster-Go.git
   cd taskmaster-go
   ```

2. **Set Up Environment Variables**:

   Create a `.env` file in the root of the project and add the following:

   ```bash
   DATABASE_URL=postgres://user:password@localhost:5432/tasks?sslmode=disable
   JWT_SECRET=your_secret_key
   ```

3. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

4. **Run Database Migrations**:

   ```bash
   go run main.go
   ```

   This command will automatically create the required database tables.

### Running the Application

To start the application, simply run:

```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Authentication

- **Register**: `POST /register`
  - Request Body: `{ "username": "string", "password": "string" }`
  - Response: `201 Created` or `400 Bad Request` (if username is taken)

- **Login**: `POST /login`
  - Request Body: `{ "username": "string", "password": "string" }`
  - Response: `200 OK` with JWT token or `401 Unauthorized`

### Tasks

- **Get All Tasks**: `GET /tasks`
  - Headers: `Authorization: Bearer {token}`
  - Response: `200 OK` with list of tasks

- **Create Task**: `POST /tasks`
  - Headers: `Authorization: Bearer {token}`
  - Request Body: `{ "title": "string", "description": "string" }`
  - Response: `201 Created` or `400 Bad Request`

- **Update Task**: `PUT /tasks/{id}`
  - Headers: `Authorization: Bearer {token}`
  - Request Body: `{ "title": "string", "description": "string" }`
  - Response: `200 OK` or `403 Forbidden` (if the task doesn't belong to the user)

- **Delete Task**: `DELETE /tasks/{id}`
  - Headers: `Authorization: Bearer {token}`
  - Response: `204 No Content` or `403 Forbidden`

## Testing the API

You can test the API using tools like [Postman](https://www.postman.com/) or `curl`.

### Example with `curl`:

- **Register**:
  ```bash
  curl -X POST http://localhost:8080/register -d '{"username":"testuser","password":"testpassword"}' -H "Content-Type: application/json"
  ```

- **Login**:
  ```bash
  curl -X POST http://localhost:8080/login -d '{"username":"testuser","password":"testpassword"}' -H "Content-Type: application/json"
  ```

  You'll receive a JWT token which you can use in the Authorization header for the following requests.

- **Create Task**:
  ```bash
  curl -X POST http://localhost:8080/tasks -d '{"title":"My First Task","description":"This is my first task."}' -H "Authorization: Bearer {your_token}" -H "Content-Type: application/json"
  ```

- **Get Tasks**:
  ```bash
  curl -X GET http://localhost:8080/tasks -H "Authorization: Bearer {your_token}"
  ```

- **Update Task**:
  ```bash
  curl -X PUT http://localhost:8080/tasks/{id} -d '{"title":"Updated Task","description":"Updated description."}' -H "Authorization: Bearer {your_token}" -H "Content-Type: application/json"
  ```

- **Delete Task**:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/{id} -H "Authorization: Bearer {your_token}"
  ```

## Project Structure

```plaintext
go-task-api/
│
├── main.go            # Entry point of the application
├── go.mod             # Go module file
├── go.sum             # Go dependencies
├── handlers/
│   ├── auth.go        # Authentication handlers
│   ├── task.go        # Task management handlers
├── models/
│   ├── jwt.go         # JWT Claims struct
│   ├── user.go        # User model
│   ├── task.go        # Task model
├── middleware/
│   └── auth.go        # Authentication middleware
├── utils/
│   ├── db.go          # Database connection setup
│   ├── hash.go        # Password hashing utilities

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you find a bug or have a suggestion.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

This README should provide comprehensive information to anyone looking to use, contribute to, or understand your project. You can customize it further based on specific needs or additional features you may add!