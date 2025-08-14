# Golang Books CRUD API

This is a simple CRUD API built with Go (Gin framework) and MySQL.

## Features
- Create a new book
- Retrieve all books
- Retrieve a book by ID
- Update a book
- Delete a book

## Requirements
- Go 1.20+
- MySQL 8+
- Git

## Setup
1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/golang-books-crud.git (TODO: change later)
   cd golang-books-crud
   ```


## Run app locally
```
go mod tidy
go run main.go
```

## Sample CURL
1. GET books
    ```
    curl http://localhost:3000/books
    ```
2. POST books
    ```
    curl -X POST http://localhost:3000/books ^
    -H "Content-Type: application/json" ^
    -d "{\"title\":\"My First Book\",\"author\":\"John Doe\",\"genre\":\"Fiction\",\"description\":\"A great story\",\"isbn\":\"1234567890\",\"image\":\"http://example.com/image.jpg\",\"published\":\"2024-01-01\",\"publisher\":\"My Publisher\"}"
    ```
3. PUT books
    ```
    curl -X PUT http://localhost:3000/books/1 ^
    -H "Content-Type: application/json" ^
    -d "{\"title\":\"Updated Book\",\"author\":\"Jane Smith\",\"genre\":\"Drama\",\"description\":\"Updated description\",\"isbn\":\"9876543210\",\"image\":\"http://example.com/new.jpg\",\"published\":\"2025-01-01\",\"publisher\":\"Updated Publisher\"}"
    ```
4. GET books by id
    ```
    curl http://localhost:3000/books/1
    ```
5. DELETE books
    ```
    curl -X DELETE http://localhost:3000/books/1
    ```


