
# Task Manager API Documentation

**Base URL:** `http://localhost:8080/api`


## Data Model: Task

A task object includes the following fields:

| Field        | Type   |Description                                                          |
|--------------|--------|--------------------------------------------------------------------------------|
| `id`         | string | Unique task identifier                                               |
| `title`      | string | Title of the task                                                    |
| `description`| string | Detailed description of the task                                     |
| `due_date`   | string | Due date in ISO 8601 format (`YYYY-MM-DDTHH:MM:SSZ`)                 |
| `status`     | string | Task status: `"Pending"`, `"In Progress"`, or `"Completed"`          |

**Example `due_date`:** `"2025-07-25T15:00:00Z"`

---

## Create a New Task

**Endpoint:** `POST /tasks`  
**Description:** Create a new task.

### 🔹 Request

**Headers:**
```

Content-Type: application/json

````

**Body:**
```json
{
  "id": "6",
  "title": "Finish project proposal",
  "description": "Write and submit the proposal to the client",
  "due_date": "2025-07-20T17:00:00Z",
  "status": "Pending"
}
````

### 🔹 Response

| Status Code       | Description                          |
| ----------------- | ------------------------------------ |
| `201 Created`     | Task created successfully            |
| `400 Bad Request` | Invalid input or task already exists |

---

## Update an Existing Task

**Endpoint:** `PUT /tasks/{id}`
**Description:** Update an existing task by ID.

### 🔹 Parameters

| Name | In   | Type   | Required | Description    |
| ---- | ---- | ------ | -------- | -------------- |
| `id` | path | string | ✅        | ID of the task |

### 🔹 Request

**Headers:**

```
Content-Type: application/json
```

**Body:**

```json
{
  "id": "6",
  "title": "Design System",
  "description": "Design End to End System",
  "due_date": "2025-07-21T17:00:00Z",
  "status": "Pending"
}
```

### 🔹 Response

| Status Code       | Description                      |
| ----------------- | -------------------------------- |
| `200 OK`          | Task updated successfully        |
| `400 Bad Request` | Invalid input data               |
| `404 Not Found`   | Task with specified ID not found |

---

## Delete a Task

**Endpoint:** `DELETE /tasks/{id}`
**Description:** Delete a task by its ID.

### 🔹 Parameters

| Name | In   | Type   | Required | Description    |
| ---- | ---- | ------ | -------- | -------------- |
| `id` | path | string | ✅        | ID of the task |

### 🔹 Response

| Status Code     | Description               |
| --------------- | ------------------------- |
| `200 OK`        | Task deleted successfully |
| `404 Not Found` | Task not found            |

**Example Response Body:**

```json
{
  "message": "Task deleted"
}
```

---

## Date Format

* All dates must be in **ISO 8601** format (RFC 3339):

  ```
  YYYY-MM-DDTHH:MM:SSZ
  ```

Example: `"2025-07-25T15:00:00Z"`

[Postman](https://documenter.getpostman.com/view/46771916/2sB34ijf71)