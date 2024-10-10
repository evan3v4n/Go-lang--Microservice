# GoLang Microservice for Order Management

This project is a microservice developed in Go, designed to manage orders with a focus on performance and scalability. It utilizes Redis as a persistent storage mechanism to store and retrieve order information efficiently.

## Features

- **Order Management**: Supports operations to create, list, retrieve, update, and delete orders.
- **Persistence**: Utilizes Redis, a key-value store, for fast and scalable storage of order data.
- **RESTful API**: Exposes a simple RESTful API to interact with the order system, making it easily integrable with other services or front-end applications.

## Project Structure

- **`application/`**: Core application setup including server initialization, configuration loading, and route definitions.
  - `app.go`: Initializes and starts the HTTP server and Redis client.
  - `config.go`: Manages configuration settings from environment variables.
  - `routes.go`: Defines HTTP routes and associates them with their handlers.
- **`handler/`**: Contains HTTP handlers that implement the logic for handling requests.
  - `order.go`: Implements order-related operations (CRUD).
- **`model/`**: Defines data structures related to orders.
  - `order.go`: Defines the order model and line item model.
- **`repository/order/`**: Implements the repository pattern for accessing order data stored in Redis.
  - `redis.go`: Contains methods for order data manipulation in Redis.
- `main.go`: The entry point of the microservice that sets up and starts the application.

## Quick Start

### Prerequisites

- Go (version 1.15 or newer)
- Redis server running on localhost:6379

### Running the Service

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Start the application by running:
   ```bash
   go run main.go

The service will start and listen for HTTP requests on the configured port (default: 3000).

## API Endpoints

- `POST /orders`: Create a new order.
- `GET /orders`: List all orders.
- `GET /orders/{id}`: Retrieve an order by its ID.
- `PUT /orders/{id}`: Update an order by its ID.
- `DELETE /orders/{id}`: Delete an order by its ID.

## Configuration

The service can be configured via environment variables:

- `REDIS_ADDR`: The address of the Redis server. Default is `localhost:6379`.
- `SERVER_PORT`: The port on which the server will listen. Default is `3000`.
