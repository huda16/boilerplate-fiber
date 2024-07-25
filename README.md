# Your Project Name

## Description

This project is a Go-based web application utilizing the Fiber framework for HTTP handling and GORM for ORM. It connects to an Oracle database using the `go-ora` package for database operations.

## Project Structure

The project is organized into the following directories and files:

.
├── app
│ ├── controllers # Contains HTTP request handlers
│ ├── dto # Data Transfer Objects
│ ├── mappers # Functions to map models to DTOs and vice versa
│ ├── models # Database models
│ ├── repositories # Database interaction layer
│ ├── services # Business logic layer
│ └── utils # Utility functions
├── config # Configuration files and loaders
├── database # Database connection and migration logic
├── docs # Documentation files
├── middlewares # Middleware functions for request handling
├── routes # Route definitions
├── go.mod # Go module file
├── go.sum # Go module checksum file
└── main.go # Application entry point


## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/your_username/your_project_name.git
   cd your_project_name

2. **Set up environment variables**

Create a .env file in the root directory with the following content:

DB_HOST=localhost
DB_PORT=1521
DB_USER=your_user
DB_PASSWORD=your_password
DB_SERVICE=your_service

3. **Install dependencies**

Ensure you have Go installed. Run the following command to install the required Go modules:
go mod tidy

4. Run the application

Start the application using:
go run main.go
The server will start on port 3000 by default. You can change the port in the main.go file if needed.