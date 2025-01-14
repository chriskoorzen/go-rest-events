# Go Events REST API

## Introduction
This project is a simple REST API built using Go. It demonstrates the basic principles of building a RESTful service, including handling HTTP requests, routing, authentication, and responding with JSON data.

## Summary
The Go Events REST API provides endpoints to manage a collection of events, creating new users, and registering users to "attend" events.

It supports the following operations:
- Create a new user
- Log in a user
- Create a new event ```Auth Required```
- Retrieve all events
- Retrieve a single event by ID
- Update an existing event by ID ```Auth Required``` ```Creator Only```
- Delete an event by ID ```Auth Required``` ```Creator Only```
- Register a logged in user for an event by ID ```Auth Required```
- Cancel registration of a logged in user for an event by ID ```Auth Required```
- Retrieve all registered users for an event by ID ```Auth Required``` ```Creator Only```

## Endpoints

### Create New User
- **URL:** `/users`
- **Method:** `POST`
- **Description:** Creates a new user.
- **Request Body:**
    ```json
    {
        "email": "user@example.com",
        "password": "password123"
    }
    ```
- **Response:**
    ```json
    {
        "message": "New user successfully created"
    }
    ```

### Login User
- **URL:** `/login`
- **Method:** `POST`
- **Description:** Logs in a user.
- **Request Body:**
    ```json
    {
        "email": "user@example.com",
        "password": "password123"
    }
    ```
- **Response:**
    ```json
    {
        "message": "Login successful",
        "token": "jwt-token"
    }
    ```

### Get All Events
- **URL:** `/events`
- **Method:** `GET`
- **Description:** Retrieves a list of all events.
- **Response:**
    ```json
    {
        "events": [
            {
                "title": "Event 1",
                "description": "Description for event 1",
                "location": "Location of for event 1",
                "datetime": "datetime as string in ISO8601 format"
            },
            ...
        ]
    }
    ```

### Get Event by ID
- **URL:** `/events/{id}`
- **Method:** `GET`
- **Description:** Retrieves a single event by its ID.
- **Response:**
    ```json
    {
        "event": {
            "title": "Event 1",
            "description": "Description for event 1",
            "location": "Location of for event 1",
            "datetime": "datetime as string in ISO8601 format"
        }
    }
    ```

### Create New Event
- **URL:** `/events`
- **Method:** `POST`
- **Authorization:** Required (JWT in "Authorization" header)
- **Authorization Level:** Logged In
- **Description:** Creates a new event.
- **Request Body:**
    ```json
    {
        "title": "Event 1",
        "description": "Description for event 1",
        "location": "Location of for event 1",
        "datetime": "datetime as string in ISO8601 format"
    }
    ```
- **Response:**
    ```json
    {
        "message": "POST successful",
        "event": {
            "title": "Event 1",
            "description": "Description for event 1",
            "location": "Location of for event 1",
            "datetime": "datetime as string in ISO8601 format"
        }
    }
    ```

### Update Event by ID
- **URL:** `/events/{id}`
- **Method:** `PUT`
- **Authorization:** Required (JWT in "Authorization" header)
- **Authorization Level:** Logged In, Creator Only
- **Description:** Updates an existing event by its ID.
- **Request Body:**
    ```json
    {
        "title": "Updated Event",
        "description": "Description for updated event",
        "location": "Location for updated event",
        "datetime": "Updated datetime as string in ISO8601 format"
    }
    ```
- **Response:**
    ```json
    {
        "message": "Successfully updated event"
    }
    ```

### Delete Event by ID
- **URL:** `/events/{id}`
- **Method:** `DELETE`
- **Authorization:** Required (JWT in "Authorization" header)
- **Authorization Level:** Logged In, Creator Only
- **Description:** Deletes an event by its ID.
- **Response:**
    ```json
    {
        "message": "Successfully deleted event"
    }
    ```

### Register User for Event
- **URL:** `/events/{id}/register`
- **Method:** `POST`
- **Authorization:** Required (JWT in "Authorization" header)
- **Authorization Level:** Logged In
- **Description:** Registers a logged in user for an event by its ID.
- **Request Body:**
    ```json
    { }
    ```
- **Response:**
    ```json
    {
        "message": "Successfully registered for event"
    }
    ```

### Cancel User Registration for Event
- **URL:** `/events/{id}/register`
- **Method:** `DELETE`
- **Authorization:** Required (JWT in "Authorization" header)
- **Authorization Level:** Logged In
- **Description:** Cancels the registration of a logged in user for an event by its ID.
- **Response:**
    ```json
    {
        "message": "Successfully cancelled registration for event"
    }
    ```

### Get All Registered Users for Event
- **URL:** `/events/{id}/register`
- **Method:** `GET`
- **Authorization:** Required (JWT in "Authorization" header)
- **Authorization Level:** Logged In, Creator Only
- **Description:** Retrieves all registered users for an event by its ID.
- **Response:**
    ```json
    {
        "registrations": [
            {
                "email": "user@example.com"
            },
            ...
        ]
    }
    ```

## How to Run
1. Clone the repository:
     ```sh
     git clone https://github.com/chriskoorzen/go-rest-events.git
     ```
2. Navigate to the project directory:
     ```sh
     cd go-rest-events
     ```
3. Build and run the application:
     ```sh
     go build
     ./go-rest-events
     ```
4. The API will be available at `http://localhost:8080`.

## Conclusion
This project serves as a basic template for building RESTful APIs in Go. Feel free to extend and modify it to suit your needs.