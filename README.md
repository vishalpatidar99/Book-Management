# Book Management Backend API

This is a backend API for managing a library of books. It provides endpoints to perform various operations such as user authentication, adding, deleting, and retrieving books.

## Features

- User authentication using JWT tokens.
- Separate endpoints for regular and admin users.
- CRUD operations for managing books.
- Reading books from CSV files.

## Endpoints

### Authentication

- **POST /api/v1/users/login**: Logs in a user and returns a JWT token for authorization.

### User Actions

- **GET /api/v1/users/home**: Retrieves a list of books based on the user type (regular or admin).
- **POST /api/v1/books/addBook**: Adds a new book to the library. Only accessible to admin users.
- **DELETE /api/v1/books/deleteBook**: Deletes a book from the library. Only accessible to admin users.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/vishalpatidar99/Book-Management.git
    cd Book-Mangement
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Build the project:

    ```bash
    go build -o book-management-api
    ```

4. Run the server:

    ```bash
    ./book-management-api
    ```

The server will start running on `http://localhost:8080` by default.

## Usage

1. Make sure the server is running.
2. Use any API testing tool (such as Postman) to send requests to the endpoints.
3. Authenticate using the `/api/v1/users/login` endpoint to obtain a JWT token.
4. Use the obtained token in the `Authorization` header for accessing other endpoints.

## cURL Examples

- **Login**:
```bash
curl --location '127.0.0.1:8080/api/v1/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "vishupatidar10@gmail.com",
    "password": "Now@12345",
    "user_type": "Admin"
}'
```

- **home**:
```bash
curl --location '127.0.0.1:8080/api/v1/users/home' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZpc2h1cGF0aWRhcjEwQGdtYWlsLmNvbSIsInVzZXJfdHlwZSI6IkFkbWluIiwiZXhwIjoxNzE0ODQ2NzYyfQ.byivgV6d-xCj22lWVhMKD8oog_4e3C0AWoxTR-kGBdg' 
```

- **addBook**:
```bash
curl --location '127.0.0.1:8080/api/v1/books/addBook' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZpc2h1cGF0aWRhcjEwQGdtYWlsLmNvbSIsInVzZXJfdHlwZSI6IkFkbWluIiwiZXhwIjoxNzE0ODQ2MjczfQ.MIf98hrFNC9JvC_5EaJPTo6Dyw9h0v-0XCFow-UZe1E' \
--data '{
    "name": "Book7",
    "author": "Author7",
    "publish_year": 2024
}'
```

- **deleteBook**:
```bash
curl --location --request DELETE '127.0.0.1:8080/api/v1/books/deleteBook?name=Book7' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZpc2h1cGF0aWRhcjEwQGdtYWlsLmNvbSIsInVzZXJfdHlwZSI6IkFkbWluIiwiZXhwIjoxNzE0ODQ2NzYyfQ.byivgV6d-xCj22lWVhMKD8oog_4e3C0AWoxTR-kGBdg'
```