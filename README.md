# E-commerce API Microservice

## Project Overview

This project is a scalable e-commerce API microservice built using **Golang**. It features real-time order processing, efficient data caching, user authentication, and is designed for cloud deployment using **Docker**.

## Features

- **User Authentication & Authorization**: JWT-based authentication and role-based access control.
- **Product Management**: Endpoints to create, update, delete, and retrieve products.
- **Order Processing**: Real-time order handling using Goroutines and Channels.
- **Caching**: Uses **Redis** for caching frequently accessed data.
- **API Rate Limiting**: Prevents abuse using middleware.
- **Cloud Deployment**: Configured for deployment on **AWS EC2** or **Azure VM**.

## Technologies Used

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: PostgreSQL for primary storage, Redis for caching
- **Containerization**: Docker
- **Cloud Platform**: Deployable on AWS or Azure
- **Libraries**: JWT for authentication, Goroutines and Channels for concurrency

## Project Structure

```
ecommerce-api/
  |-- main.go                   // Entry point
  |-- go.mod                    // Go module file
  |-- config/
       |-- config.go            // Config settings
  |-- db/
       |-- db.go                // Database connection
  |-- routes/
       |-- routes.go            // Route initialization
  |-- controllers/
       |-- productController.go // Product request handlers
       |-- orderController.go   // Order request handlers
  |-- models/
       |-- product.go           // Product data model
       |-- order.go             // Order data model
  |-- services/
       |-- orderProcessing.go   // Order processing logic
  |-- middleware/
       |-- authMiddleware.go    // JWT authentication middleware
  |-- Dockerfile                // Docker configuration
  |-- README.md                 // Project documentation
```

## Installation and Setup

### Prerequisites

- **Docker**
- **Go** (if running locally)
- **PostgreSQL**
- **Redis**

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/Dibyendu-13/golang_ecommerce_project.git
   cd golang_ecommerce_project
   ```
2. Start the database:
   ```bash
   docker run --name postgres-db -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres
   ```
3. Update the `db.go` connection string if needed:
   ```go
   connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
   ```
4. Run the Go application:
   ```bash
   go run main.go
   ```

### Running with Docker

1. Build the Docker image:
   ```bash
   docker build -t ecommerce-api .
   ```
2. Run the container:
   ```bash
   docker run -p 8080:8080 ecommerce-api
   ```

### Running Both Containers on a Custom Network

1. Create a Docker network:
   ```bash
   docker network create mynetwork
   ```
2. Run PostgreSQL on the network:
   ```bash
   docker run --network mynetwork --name postgres-db -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres
   ```
3. Run the Go application container:
   ```bash
   docker run --network mynetwork -p 8080:8080 ecommerce-api
   ```

## API Endpoints

- **POST /products**: Create a new product.
- **GET /products/:id**: Retrieve a product by ID.
- **POST /orders**: Create a new order.

## Deployment

- **Containerization**: The application is ready for deployment with **Docker**.
- **Cloud Deployment**: Can be deployed on **AWS EC2** or **Azure VM**.

## Sample Dockerfile

```dockerfile
FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .
EXPOSE 8080
CMD ["./main"]
```

## Contributing

Feel free to submit pull requests or open issues for improvements and new features.

## License

This project is licensed under the MIT License.
