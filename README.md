# TaskMaster Go

**TaskMaster Go** is a RESTful API built with Go, designed to manage tasks with full CRUD (Create, Read, Update, Delete) functionality. This API features secure user authentication using JWT tokens, password hashing, and PostgreSQL for data persistence. The project is structured to be scalable, maintainable, and ready for deployment.

## Features

- **User Authentication**: Secure user registration and login with hashed passwords.
- **JWT-Based Authorization**: Ensures secure access to API endpoints.
- **Task Management**: Users can create, read, update, and delete their tasks.
- **Database Integration**: Persistent storage using PostgreSQL.
- **Modular Architecture**: Clean code structure following Go best practices.

## Tech Stack

- **Language**: Go
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **ORM**: GORM
- **Routing**: Gorilla Mux

## Project Structure

```
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
└── utils/
    ├── db.go          # Database connection setup
```

## Getting Started

### Prerequisites

Before running this project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/taskmaster-go.git
   cd taskmaster-go
   ```

2. **Set Up Environment Variables**:

   Create a `.env` file in the root of the project and add the following:

   ```env
   DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/tasks?sslmode=disable
   JWT_SECRET=your_secret_key
   ```

   - Replace `yourpassword` with the password for your PostgreSQL `postgres` user.
   - Replace `your_secret_key` with a secure key for JWT signing.

3. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

4. **Create the `tasks` Database**:

   Open the PostgreSQL command line interface:

   ```bash
   psql postgres
   ```

   Create the `tasks` database:

   ```sql
   CREATE DATABASE tasks OWNER postgres;
   ```

   Exit the `psql` CLI:

   ```sql
   \q
   ```

5. **Run Database Migrations**:

   The Go application will automatically create the required database tables when you run the application for the first time.

### Running the Application

To start the application, simply run:

```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Authentication

- **Register**: `POST /register`
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response: `201 Created` or `400 Bad Request`

- **Login**: `POST /login`
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response: `200 OK` with JWT token or `401 Unauthorized`

### Tasks

- **Get All Tasks**: `GET /tasks`
  - Headers: `Authorization: Bearer {token}`
  - Response: `200 OK` with a list of tasks

- **Create Task**: `POST /tasks`
  - Headers: `Authorization: Bearer {token}`
  - Request Body:
    ```json
    {
      "title": "string",
      "description": "string"
    }
    ```
  - Response: `201 Created` or `400 Bad Request`

- **Update Task**: `PUT /tasks/{id}`
  - Headers: `Authorization: Bearer {token}`
  - Request Body:
    ```json
    {
      "title": "string",
      "description": "string"
    }
    ```
  - Response: `200 OK` or `403 Forbidden`

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

  You'll receive a JWT token that you can use in the Authorization header for subsequent requests.

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

## Deployment

To deploy this application, you can follow these steps:

1. **Containerize the Application**:
   - Use Docker to create a containerized version of the app, which makes it easier to deploy on various platforms.

2. **Deploy to a Cloud Provider**:
   - Deploy the application on platforms like Heroku, AWS, Google Cloud, or DigitalOcean.
   - Ensure that the database is set up and accessible from the deployed environment.

3. **Use Environment Variables**:
   - Store sensitive information like `DATABASE_URL` and `JWT_SECRET` in environment variables.

4. **Configure HTTPS**:
   - If you're deploying in production, configure HTTPS for secure communication.

## Future Improvements

Here are some potential enhancements you can make to this project:

- **Pagination**: Implement pagination in the `GET /tasks` endpoint to handle large datasets.
- **Task Categories**: Add support for task categorization and tagging.
- **Task Deadlines**: Implement due dates for tasks with automatic notifications or reminders.
- **User Roles**: Extend the API to support different user roles (e.g., admin, regular user).
- **Frontend Integration**: Build a web or mobile interface that interacts with this API.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, feel free to open an issue or submit a pull request.