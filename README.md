# GO USER API

## Description

Go User API is a RESTful API built with Golang that allows performing CRUD (Create, Read, Update, Delete) operations on a user resource. Authentication is implemented using JWT to ensure secure access.

## Features (Planned)

- User Registration
- User Login
- Change Password
- Find User by ID
- Delete User

## Technologies Used

- **Programming Language:** Go (Golang)
- **Framework:** Fiber
- **Database:** MySQL
- **Logging:** Logrus
- **Configuration Management:** Viper
- **Containerization:** Docker, Docker Compose

## Installation Guide

### 1. Clone the Repository

```bash
git clone https://github.com/nurmanhadi/go-user-api.git
```

### 2. Configure the Application

Edit the `config.json` file to set up the necessary configurations, such as database credentials and JWT settings.

### 3. Create the Database

Create a MySQL database with the name `go_user_api` (or any preferred name) and update the database URL in `config.json` under `mysql.url`.

### 4. Run Database Migrations

Execute the migration script using any database migration tool. If using Golang Migrate, run:

```bash
migrate up ./db/migrations/up.sql
```

### 5. Run the Application

#### Using Go Run

```bash
go run cmd/web/main.go
```

#### Using Docker Compose

If you prefer running the application with Docker Compose, execute the following command:

```bash
docker compose create
docker compose start
```

This will start the application along with the MySQL database container as defined in the `docker-compose.yml` file.

## API Documentation

### Postman Collection

The API collection is available in Postman for testing and integration purposes [here](https://www.postman.com/supply-observer-10980491/workspace/nurman-hadi/collection/40970638-05f05d8e-7aaf-4820-aaf9-aa17c7cb4d11?action=share&creator=40970638).

### Endpoints Overview

| Method  | Endpoint                | Description          |
|---------|-------------------------|----------------------|
| POST    | /api/v1/auth/register   | Register a new user |
| POST    | /api/v1/auth/login      | User login          |
| PUT     | /api/v1/auth/password   | Change password     |
| GET     | /api/v1/users/{id}      | Get user by ID      |
| DELETE  | /api/v1/users/{id}      | Delete user         |

### 1. Register a New User

**Endpoint:** `POST /api/v1/auth/register`

**Request Body:**
```json
{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "securepassword"
}
```

**Response (201 Created):**
```json
{
    "links": {
        "self": "/api/v1/auth/register"
    }
}
```

### 2. User Login

**Endpoint:** `POST /api/v1/auth/login`

**Request Body:**
```json
{
    "email": "john.doe@example.com",
    "password": "securepassword"
}
```

**Response (200 Success):**
```json
{
    "data": {
        "access_token": "<JWT_TOKEN>"
    },
    "links": {
        "self": "/api/v1/auth/login"
    }
}
```

### 3. Change Password

**Endpoint:** `PUT /api/v1/auth/password`

**Headers:**
- `Authorization: Bearer <JWT_TOKEN>`

**Request Body:**
```json
{
    "old_password": "securepassword",
    "new_password": "newsecurepassword"
}
```

**Response (200 Success):**
```json
{
    "links": {
        "self": "/api/v1/auth/password"
    }
}
```

### 4. Get User by ID

**Endpoint:** `GET /api/v1/users/{id}`

**Headers:**
- `Authorization: Bearer <JWT_TOKEN>`

**Response (200 Success):**
```json
{
    "data": {
        "id": "db1aede1-d8fd-4ddb-84be-324aecaa6a23",
        "name": "John Doe",
        "email": "john.doe@example.com",
        "created_at": "2025-03-04T07:50:50-05:00",
        "updated_at": "2025-03-04T07:50:50-05:00"
    },
    "links": {
        "self": "/api/v1/users/db1aede1-d8fd-4ddb-84be-324aecaa6a23"
    }
}
```

### 5. Delete User

**Endpoint:** `DELETE /api/v1/users/{id}`

**Headers:**
- `Authorization: Bearer <JWT_TOKEN>`

**Response (200 Success):**
```json
{
    "links": {
        "self": "/api/v1/users/{id}"
    }
}
```