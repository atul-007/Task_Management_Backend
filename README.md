# Task Management Backend

This is the backend application for managing tasks. It provides endpoints for user registration, authentication, and task management.

## Installation

To install and run the application, follow these steps:

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/atul-007/task_management_backend.git
    ```

2. Navigate to the project directory:

    ```bash
    cd task_management_backend
    ```

3. Build the application:

    ```bash
    go build -tags netgo -ldflags '-s -w' -o app
    ```

4. Run the application:

    ```bash
    ./app
    ```

The application will be accessible at `http://localhost:8080`.

## Usage

### Routes

#### User Registration

- **Route:** `POST /api/register`
- **Payload:**
    ```json
    {
        "username": "example_user",
        "password": "example_password"
    }
    ```
- **Description:** Registers a new user with the provided username and password.

#### User Login

- **Route:** `POST /api/login`
- **Payload:**
    ```json
    {
        "username": "example_user",
        "password": "example_password"
    }
    ```
- **Description:** Logs in a user with the provided username and password. Returns a JWT token upon successful login.

#### Get Tasks

- **Route:** `GET /api/tasks`
- **Description:** Retrieves tasks for the authenticated user.

#### Create Task

- **Route:** `POST /api/tasks`
- **Payload:**
    ```json
    {
        "title": "Example Task",
        "priority": "High",
        "due_date": "2024-05-15T12:00:00Z",
        "completed": false
    }
    ```
- **Description:** Creates a new task for the authenticated user.

#### Update Task

- **Route:** `PUT /api/tasks/:id`
- **Payload:**
    ```json
    {
        "title": "Updated Task",
        "priority": "Low",
        "due_date": "2024-05-20T12:00:00Z",
        "completed": true
    }
    ```
- **Description:** Updates an existing task for the authenticated user.

#### Delete Task

- **Route:** `DELETE /api/tasks/:id`
- **Description:** Deletes an existing task for the authenticated user.

Deployed Link:https://task-management-backend-3boh.onrender.com
