# 📘 Task-Manager API Documentation

## Overview

The Task-Manager API provides functionality to manage tasks and users. It supports creating, reading, updating, and deleting tasks, as well as managing users via signup, login, and administrative actions.

---

## 🔗 Base URL

```
http://localhost:8080/api
```

---

## 📦 Data Model

### Task Object

```json
{
  "id": "string",                     // Unique Task ID
  "title": "string",                 // Task title (required)
  "description": "string",           // Task description
  "due_date": "2025-07-25T15:00:00Z", // ISO 8601 format (RFC3339)
  "status": "Pending"                // One of: "Pending", "In Progress", "Completed"
}
```

---

## 📒 Task Endpoints

### 📋 List All Tasks

**GET** `/tasks`

Retrieve all tasks.

#### Response

```json
[
  {
    "id": "1",
    "title": "Finish project proposal",
    "description": "Write and submit the proposal to the client",
    "due_date": "2025-07-20T17:00:00Z",
    "status": "Pending"
  }
]
```

**Status Codes**

* `200 OK` – Success

---

### 🔍 Get Task by ID

**GET** `/tasks?id={id}`

Retrieve a task by its ID.

#### Query Parameters

| Name | Type   | Required | Description     |
| ---- | ------ | -------- | --------------- |
| id   | string | ✅        | Task identifier |

#### Response

```json
{
  "id": "1",
  "title": "Finish project proposal",
  "description": "Write and submit the proposal to the client",
  "due_date": "2025-07-20T17:00:00Z",
  "status": "Pending"
}
```

**Status Codes**

* `200 OK` – Task found
* `404 Not Found` – Task does not exist

---

### 📝 Create New Task

**POST** `/tasks`

Create a new task.

#### Request Body

```json
{
  "id": "6",
  "title": "Finish project proposal",
  "description": "Write and submit the proposal to the client",
  "due_date": "2025-07-20T17:00:00Z",
  "status": "Pending"
}
```

**Status Codes**

* `201 Created` – Task successfully created
* `400 Bad Request` – Invalid format or duplicate ID

  ```json
  { "error": "task ID 'id' already exists" }
  ```

---

### ✏️ Update Task

**PUT** `/tasks/{id}`

Update an existing task by ID.

#### Path Parameters

| Name | Type   | Required | Description     |
| ---- | ------ | -------- | --------------- |
| id   | string | ✅        | Task identifier |

#### Request Body

```json
{
  "id": "6",
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "2025-07-21T17:00:00Z",
  "status": "In Progress"
}
```

**Status Codes**

* `200 OK` – Task updated
* `400 Bad Request` – Invalid request
* `404 Not Found` – Task not found

---

### ❌ Delete Task

**DELETE** `/tasks/{id}`

Delete a task by ID.

#### Path Parameters

| Name | Type   | Required | Description     |
| ---- | ------ | -------- | --------------- |
| id   | string | ✅        | Task identifier |

#### Response

```json
{ "message": "Task deleted" }
```

**Status Codes**

* `200 OK` – Task deleted
* `404 Not Found` – Task not found

---

## 👤 User Endpoints

### 🔐 Signup

**POST** `/users/signup`

Register a new user.

#### Request Body

```json
{
  "first_name": "Segni",
  "last_name": "Girma",
  "username": "valid_result",
  "password": "12345678",
  "user_type": "USER"
}
```

**Status Codes**

* `201 Created` – User created
* `400 Bad Request` – Invalid input or admin required for first user
* `409 Conflict` – Username already exists
* `500 Internal Server Error` – Server error during signup

---

### 🔑 Login

**GET** `/users/login`

Authenticate a user.

**Status Codes**

* `200 OK` – Login successful
* `400 Bad Request` – Invalid input
* `403 Unauthorized` – Invalid credentials
* `500 Internal Server Error` – Token generation failure

---

### 👥 List All Users (Admin Only)

**GET** `/users`

Retrieve a list of all registered users.

**Status Codes**

* `200 OK` – Success
* `400 Bad Request` – Invalid request
* `403 Forbidden` – Unauthorized access (not ADMIN)
* `500 Internal Server Error` – Error retrieving users

---

### 🔍 Get User by ID

**GET** `/users/{user_id}`

Retrieve a specific user by their ID.

#### Path Parameters

| Name     | Type   | Required | Description    |
| -------- | ------ | -------- | -------------- |
| user\_id | string | ✅        | Unique user ID |

#### Example

```
GET /users/12345
```

#### Response

```json
{
  "id": "12345",
  "first_name": "Segni",
  "last_name": "Girma",
  "username": "valid_result",
  "user_type": "USER"
}
```

**Status Codes**

* `200 OK` – User found
* `400 Bad Request` – Invalid ID
* `403 Forbidden` – Unauthorized (non-admin or not self)
* `404 Not Found` – User not found
* `500 Internal Server Error` – Server error

---

## 📌 Notes

* All date fields (`due_date`) must be in **ISO 8601** (RFC3339) format.
* `status` values must be one of: `"Pending"`, `"In Progress"`, or `"Completed"`.
* No authentication headers are shown in the Postman collection, but in a real-world application, all protected routes should enforce token-based authentication (e.g., JWT).



[Postman](https://documenter.getpostman.com/view/46771916/2sB34ijf71)