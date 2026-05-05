# E-Commerce API

A production-ready e-commerce backend API built with Go, designed with clean architecture principles and a modular service-oriented structure.

## Overview

This project is a fully-featured e-commerce API that provides comprehensive functionality for managing users, products, orders, and shopping carts. It's built with Go's performance and concurrency in mind, using a MySQL database for persistent storage.

## Features

- **User Management**: User registration, login, and JWT-based authentication
- **Product Management**: Create, retrieve, and manage products
- **Shopping Cart**: Add/remove items from cart, manage cart operations
- **Order Management**: Process orders and manage order items
- **Authentication**: JWT token-based secure authentication with password hashing
- **Database Migrations**: Version-controlled database schema management

## Tech Stack

| Component               | Technology                    |
| ----------------------- | ----------------------------- |
| **Language**            | Go 1.25.3                     |
| **Web Framework**       | Gorilla Mux (HTTP routing)    |
| **Database**            | MySQL                         |
| **Authentication**      | JWT (golang-jwt)              |
| **Password Hashing**    | bcrypt (golang.org/x/crypto)  |
| **Validation**          | Validator v10 (go-playground) |
| **Database Migrations** | golang-migrate v4             |
| **Environment Config**  | GoDotEnv                      |

## Project Structure

```
.
├── bin/                           # Compiled binaries
├── cmd/
│   ├── main.go                   # Application entry point
│   ├── api/
│   │   └── api.go               # API server setup and route registration
│   ├── migrate/
│   │   └── main.go              # Migration runner
│   └── migrations/              # SQL migration files
│       ├── *_add-user-table.sql
│       ├── *_add-products-table.sql
│       ├── *_add-orders-table.sql
│       └── *_add-order_items-table.sql
├── config/
│   └── env.go                    # Environment configuration loader
├── db/
│   └── db.go                     # Database connection setup
├── service/                       # Business logic layer
│   ├── auth/
│   │   ├── jwt.go               # JWT token generation and validation
│   │   └── password.go          # Password hashing and verification
│   ├── user/
│   │   ├── routes.go            # User routes (login, register)
│   │   ├── store.go             # User data access layer
│   │   └── routes_test.go       # User route tests
│   ├── product/
│   │   ├── routes.go            # Product routes
│   │   └── store.go             # Product data access layer
│   ├── order/
│   │   └── store.go             # Order data access layer
│   └── cart/
│       ├── routes.go            # Cart routes
│       └── service.go           # Cart business logic
├── types/
│   └── types.go                  # Core type definitions and interfaces
├── utils/
│   └── utils.go                  # Utility functions
├── go.mod                         # Go module definition
├── Makefile                       # Build and utility commands
└── .env                          # Environment variables (not tracked in git)
```

## Installation

### Prerequisites

- Go 1.25.3 or later
- MySQL 5.7 or later
- Make (for running build commands)

### Clone the Repository

```bash
git clone https://github.com/Eslam-Amin/ecommerce.git
cd ecommerce
```

### Install Dependencies

```bash
go mod download
```

### Environment Setup

Create a `.env` file in the project root with the following variables:

```env
# Database Configuration
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_ADDRESS=localhost:3306
DB_NAME=ecommerce

# Server Configuration
PUBLIC_HOST=http://localhost:8080
PORT=8080

# JWT Configuration
JWT_SECRET=your_secret_key_here
JWT_EXPIRATION_IN_SECONDS=3600
```

## Database Setup

### Run Migrations

```bash
# Run all pending migrations
make migrate-up

# Rollback the latest migration
make migrate-down

# Create a new migration
make migration name=your_migration_name
```

## Build & Run

### Build the Application

```bash
make build
```

### Run the Application

```bash
make run
```

The API server will start on `http://localhost:8080`.

### Run Tests

```bash
make test
```

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### User Routes

| Method | Endpoint    | Description             |
| ------ | ----------- | ----------------------- |
| `POST` | `/register` | Register a new user     |
| `POST` | `/login`    | Login and get JWT token |

### Product Routes

| Method | Endpoint    | Description           |
| ------ | ----------- | --------------------- |
| `GET`  | `/products` | Retrieve all products |
| `POST` | `/products` | Create a new product  |

### Cart Routes

| Method   | Endpoint        | Description           |
| -------- | --------------- | --------------------- |
| `GET`    | `/cart`         | Get current cart      |
| `POST`   | `/cart`         | Add item to cart      |
| `DELETE` | `/cart/:itemId` | Remove item from cart |

### Order Routes

| Method | Endpoint  | Description            |
| ------ | --------- | ---------------------- |
| `GET`  | `/orders` | Retrieve user's orders |
| `POST` | `/orders` | Create a new order     |

## Authentication

The API uses JWT (JSON Web Token) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

Tokens are obtained via the `/register` or `/login` endpoints and are configured to expire based on the `JWT_EXPIRATION_IN_SECONDS` environment variable.

## Development

### Code Organization

- **Handlers**: HTTP request handlers in each service's `routes.go` file
- **Stores**: Database queries and data operations in `store.go` files
- **Services**: Business logic in dedicated service packages
- **Types**: Interfaces and data structures defined in `types.go`

### Running Locally

```bash
# Terminal 1: Start MySQL
# (Ensure MySQL is running on the configured address)

# Terminal 2: Run the application
make run
```

The API will be available at `http://localhost:8080/api/v1`.

## Database Schema

The database includes the following tables:

- **users**: Stores user information and authentication credentials
- **products**: Product catalog with details and pricing
- **orders**: Customer orders
- **order_items**: Individual items within an order

## Error Handling

The API returns standard HTTP status codes and JSON error responses:

```json
{
  "error": "Error message description"
}
```

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## Author

**Eslam Amin**

- GitHub: [@Eslam-Amin](https://github.com/Eslam-Amin)
- Email: contact@example.com

## Support

For issues and questions, please open an issue on the GitHub repository.

---

**Last Updated**: May 2026
