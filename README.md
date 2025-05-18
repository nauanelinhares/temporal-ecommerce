## 🚀 Getting Started

This project is a temporal e-commerce application built with Go.

### ⚙️ Technologies Used

- Go
- Fiber (Web Framework)
- Gorm (ORM)
- PostgreSQL (Database)
- Docker Compose
- Temporal (Workflow Engine - *Inferred from project name and structure*)

### Prerequisites

Make sure you have the following installed:

- Go (version 1.20 or higher)
- Docker
- Docker Compose

### Installation and Setup

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd temporal-ecommerce
   ```

2. Copy the example environment file and update it with your database credentials:

   ```bash
   cp .envrc.example .envrc 
   # Edit .envrc with your database configuration
   ```

3. Start the required services using Docker Compose (this will start the PostgreSQL database):

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

To run the Go application, you will typically use the `go run` command or build an executable. (Further instructions might be needed based on how the main application is structured in `main.go` and the `cmd` directory.)

```bash
# Example (may vary based on project structure)
go run main.go
```

### 🧹 Cleaning Up

To stop the Docker Compose services:

```bash
make deps/down
```

To revert database migrations:

```bash
make migrate/down
``` 