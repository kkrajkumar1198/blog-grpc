# Blog Post using gRPC API

## Overview

This project implements a basic blogging system using gRPC for communication. It consists of server and client components, handling operations such as creating, retrieving, and deleting blog posts. The project is structured with server-side, client-side, and shared proto files.

## Requirements

- Go 1.19+
- SQLite
- HTTP client: POSTMAN, Insomnia, cURL

## Running Locally

1. Clone the repository:

    ```bash
    git clone https://github.com/kkrajkumar1198/blog-grpc.git
    ```

2. Browse into the project directory:

    ```bash
    cd blog-grpc/
    ```

3. Download dependencies:

    ```bash
    go mod tidy
    ```

4. Set environment variables following the `.env` file and run:

    ```bash
    go run main.go
    ```

    As concurrency was added for the HTTP server, the above command can run the entire project.

## Project Structure

```bash
blog-grpc
├── internal
│   ├── blog
│   │   ├── data
│   │   │   └── dataIface.go
│   │   │   └── dataImplementation.go
│   │   ├── models
│   │   │   └── models.go
│   │   ├── protos
│   │   │   └── bin
│   │   │       ├── blog_grpc.pb.go
│   │   │       └── blog.pb.go
│   │   │   └── blog.proto
│   │   ├── service.go
│   │   └── validations.go
│   ├── databases
│   │   └── migrations.go
│   │   └── interface.go
│   │   └── SQLite.go
│   └── server
│       └── server.go
├── pkg
│   └── client
│       ├── client.go
│       ├── connector.go
│       └── http
│           ├── handlers.go
│           ├── router.go
│           └── server.go
│           └── response.go
├── tests
│   └── grpc_mocks.go
│   └── service_test.go.go
├── .env
├── README.md
└── go.mod
└── go.sum
└── data.db
└── main.go
```

## Documentation

API Documentation: Swagger API documentation can be accessed at [http://localhost:3000/api/v1/swagger/index.html](http://localhost:3000/api/v1/swagger/index.html).

## Testing

Run the tests using the following command:

```bash
go test ./tests
```