# Task Manager API Documentation

## Base URL

```
http://localhost:8080/api
```

---

## Data Model

```json
{
  "id": "string",                  // Unique identifier
  "title": "string",              // Task title (required)
  "description": "string",        // Detailed task description
  "due_date": "2025-07-25T15:00:00Z", // Due date in RFC3339 format (ISO8601)
  "status": "Pending"             // Task status: "Pending", "In Progress", "Completed"
}
```

---

## Endpoints

### 1. Get All Tasks

* **Method:** `GET`
* **URL:** `/tasks`
* **Description:** Retrieve a list of all tasks.
* **Response:**

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

* **Status Code:** `200 OK`

---

### 2. Get Task by ID

* **Method:** `GET`
* **URL:** `/tasks?id={id}`
* **Parameters:**

  * `id` (string) – Task ID
* **Description:** Retrieve a specific task by ID.
* **Response:**

```json
{
  "id": "1",
  "title": "Finish project proposal",
  "description": "Write and submit the proposal to the client",
  "due_date": "2025-07-20T17:00:00Z",
  "status": "Pending"
}
```

* **Status Codes:**

  * `200 OK` – Task found
  * `404 Not Found` – Task not found

---

### 3. Create a New Task

* **Method:** `POST`
* **URL:** `/tasks`
* **Description:** Create a new task
* **Body Example:**

```json
{
  "id": "6",
  "title": "Finish project proposal",
  "description": "Write and submit the proposal to the client",
  "due_date": "2025-07-20T17:00:00Z",
  "status": "Pending"
}
```

* **Status Codes:**

  * `201 Created` – Created successfully
  * `400 Bad Request` – Invalid data or task already exists

    * Example errors:

      ```json
      { "error": "failed to check task ID uniqueness ..." }
      { "error": "task ID 'id' already exists" }
      ```

---

### 4. Update an Existing Task

* **Method:** `PUT`
* **URL:** `/tasks/{id}`
* **Parameters:**

  * `id` (string) – Task ID
* **Description:** Update an existing task by ID
* **Body Example:**

```json
{
  "id": "6",
  "title": "Design System",
  "description": "Design End to End System",
  "due_date": "2025-07-21T17:00:00Z",
  "status": "Pending"
}
```

* **Status Codes:**

  * `200 OK` – Task updated successfully
  * `400 Bad Request` – Invalid data
  * `404 Not Found` – Task not found

---

### 5. Delete a Task

* **Method:** `DELETE`
* **URL:** `/tasks/{id}`
* **Parameters:**

  * `id` (string) – Task ID
* **Description:** Delete a task by ID
* **Response:**

```json
{
  "message": "Task deleted"
}
```

* **Status Codes:**

  * `200 OK` – Task deleted successfully
  * `404 Not Found` – Task not found


[Postman](https://documenter.getpostman.com/view/46771916/2sB34ijf71)