# Go Fiber Auth API

## Overview
Go Fiber Auth API is a lightweight and efficient authentication system built using Go Fiber, GORM, and PostgreSQL. It features JWT-based authentication and allows user management with role-based access control.

## Features
- **User Registration**
- **User Login (JWT Authentication)**
- **Retrieve All Users (Admin Only)**
- **Retrieve User by ID**
- **Update User Information (User & Admin)**
- **Delete User (Admin Only)**

## Prerequisites
- **Go** (>= 1.19)
- **Docker & Docker Compose**
- **PostgreSQL**
- **Postman or cURL** for API testing

## Installation & Setup

### 1. Clone the repository
```sh
$ git clone https://github.com/yourusername/go-fiber-auth.git
$ cd go-fiber-auth
```

### 2. Configure Environment Variables
Create a `.env` file in the root directory and add:
```env
DATABASE_URL=postgres://user:password@db:5432/gofiber_auth
JWT_SECRET=your_secret_key
```

### 3. Run with Docker
```sh
$ docker-compose up --build
```

### 4. Apply Database Migrations
If the database is not migrated, run:
```sh
$ docker exec -it go-fiber-auth-app go run migrate.go
```

## API Endpoints & Usage

### 1. Register User
**Endpoint:**
```http
POST /api/register
```
**Request Body:**
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "password123"
}
```
**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "johndoe@example.com",
  "is_admin": false
}
```

### 2. Login User
**Endpoint:**
```http
POST /api/login
```
**Request Body:**
```json
{
  "email": "johndoe@example.com",
  "password": "password123"
}
```
**Response:**
```json
{
  "token": "your_jwt_token"
}
```

### 3. Get All Users (Admin Only)
**Endpoint:**
```http
GET /api/users
```
**Headers:**
```json
{
  "Authorization": "Bearer your_jwt_token"
}
```

### 4. Get User by ID
**Endpoint:**
```http
GET /api/users/{id}
```
**Headers:**
```json
{
  "Authorization": "Bearer your_jwt_token"
}
```

### 5. Update User (User or Admin)
**Endpoint:**
```http
PUT /api/users/{id}
```
**Headers:**
```json
{
  "Authorization": "Bearer your_jwt_token"
}
```
**Request Body:**
```json
{
  "name": "Updated Name",
  "email": "updatedemail@example.com"
}
```

### 6. Delete User (Admin Only)
**Endpoint:**
```http
DELETE /api/users/{id}
```
**Headers:**
```json
{
  "Authorization": "Bearer your_jwt_token"
}
```

## Testing with Postman
1. **Import API Collection** into Postman.
2. Set the **Base URL** to your running API instance.
3. Register a user and log in to obtain a **JWT token**.
4. Use the **JWT token** in the `Authorization` header for protected routes.

## License
This project is licensed under the MIT License.

