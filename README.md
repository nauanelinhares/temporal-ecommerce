## 🚀 Getting Started

This project is a temporal e-commerce application built with Go, designed to showcase how Temporal can be used to manage complex, long-running, and reliable business processes in an e-commerce setting.

### ✨ Features (and Learnings)

*   **RESTful API**: Built using the [Fiber](https://gofiber.io/) web framework, known for its speed and Express.js-like API.
    *   `GET /`: Simple welcome message.
    *   `/health`: Health check endpoint.
    *   Entity management for Users, Products, and Orders (specific routes under `/users`, `/products`, `/orders` - to be detailed further).
*   **Workflow Management with Temporal**:
    *   Utilizes [Temporal.io](https://temporal.io/) to orchestrate e-commerce workflows such as order processing, payment, inventory updates, and notifications. This ensures reliability and fault tolerance for critical business operations.
    *   **Learning**: How to initialize and use the Temporal client in a Go application.
*   **Database Interaction**:
    *   Uses [GORM](https://gorm.io/) as an ORM for interacting with a PostgreSQL database.
    *   **Learning**: Setting up GORM, defining models (assumed), and connecting to PostgreSQL.
*   **Configuration Management**:
    *   Loads application configuration (e.g., database connection details) from an external source (likely `.envrc` as per setup instructions).
    *   **Learning**: Best practices for separating configuration from code.
*   **Structured Project Layout**:
    *   Organized into logical packages like `internal/config`, `src/web/handlers`, promoting maintainability.

### ⚙️ Technologies Used

- Go
- Fiber (Web Framework)
- Gorm (ORM)
- PostgreSQL (Database)
- Docker Compose
- Temporal (Workflow Engine)

### Prerequisites

Make sure you have the following installed:

- Go (version 1.20 or higher)
- Docker
- Docker Compose
- Access to a running Temporal Server (either local via Docker Compose or a cloud service)

### Installation and Setup

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd temporal-ecommerce
   ```

2. Create a .env file and put your database configuration

   ```bash
   # Example based on main.go:
   # DB_HOST=localhost
   # DB_USER=youruser
   # DB_PASSWORD=yourpassword
   # DB_NAME=yourdbname
   # DB_PORT=5432
   # DB_SSLMODE=disable 
   ```

3. Start the required services using Docker Compose (this will start the PostgreSQL database).
   You might also need to start a Temporal server instance if you don't have one running.
   (Consider adding a Temporal service to your `docker-compose.yml` for local development).

   ```bash
   make deps/up
   ```

   If port 5432 is already in use, you can try the alternative port:

   ```bash
   make deps/up-alt
   ```

4. Apply database migrations:

   ```bash
   make migrate/up
   ```

### ▶️ Running the Application

1. Ensure your PostgreSQL database and Temporal server are running and accessible.
2. Run the application from the project root:

```bash
go run main.go
```

The application will start and listen on port `3000`. You can access it at `http://localhost:3000`.
The root endpoint `GET /` should return "Hello, World!".

### 🛣️ API Endpoints

*   `GET /`: Returns a welcome message.
*   `/health`: Provides health status of the application.
*   User Endpoints: (e.g., `POST /users`, `GET /users/{id}`) - Managed by `UserHandler`.
*   Product Endpoints: (e.g., `POST /products`, `GET /products/{id}`) - Managed by `ProductHandler`.
*   Order Endpoints: (e.g., `POST /orders`, `GET /orders/{id}`) - Managed by `OrderHandler`, likely triggering Temporal workflows.

(Note: Specific paths and request/response formats for User, Product, and Order endpoints should be documented further as they are developed.)

###  workflows Workflows with Temporal

This project leverages Temporal to manage various e-commerce processes. Workflows are defined to handle operations like:

*   Order creation and processing
*   Payment authorization and capture
*   Inventory management
*   User notifications

These workflows ensure that even if parts of the system fail or take a long time, the overall process completes reliably. The handlers (e.g., `OrderHandler`) are responsible for initiating and interacting with these Temporal workflows.

### 🧹 Cleaning Up

To stop the Docker Compose services:

```bash
make deps/down
```

To revert database migrations:

```bash
make migrate/down
``` 