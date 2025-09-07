# Go Project with DDD, JWT Auth, and File Management

This is a **Golang** project structured using **Domain-Driven Design (DDD)** principles. The project includes:

## Features

* **Categories**: Manage and organize items with category support.
* **Files**: Upload, store, and manage files efficiently.
* **Authentication**: Secure endpoints using **JWT-based authentication**.
* **DDD Architecture**: Clear separation of domain, application, and infrastructure layers for maintainability and scalability.
* Database migrations with golang-migrate
* PostgreSQL database
* Environment configuration via `.env`

---

## Project Structure

```
meditation-backend/
├── api/                     # API docs
│   └── openapi.yaml
├── cmd/                     # Main entry
│   └── server/main.go
├── internal/
│   ├── config/              # Configuration and environment setup
│   ├── domain/              # Models
│   ├── middleware/          # HTTP and gRPC middleware (auth, logging, etc.)
│   ├── pkg/                 # Reusable packages and utilities shared across the project
│   ├── repository/          # Database access layer using (pgx)
│   ├── usecase/             # Business logic / services
│   ├── tests/               # Unit and integration tests
│   ├── delivery/
│   │   └── http/            # HTTP Handlers & validators
│   │   └── grpc/            # gRPC Handlers & validators
│   ├── server/
│   │   └── router/router.go # Router setup and route definitions
│   ├── db/
│   │   └── migrations/      # SQL migration files
│   └── utils/               # Helper functions
├── .env                     # Environment variables
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module file
├── Makefile                 # Build, migrate, and deploy commands
├── Dockerfile               # Docker setup for the project
└── README.md                # Project documentation
```


---

## Environment Variables

Create a `.env` file at the project root:

```env
DB_USER=postgres
DB_PASS=secret
DB_HOST=localhost
DB_PORT=5432
DB_NAME=meditation_db

DATABASE_URL=postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
JWT_SECRET=your_secret_key
```

---

## Setup

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Install golang-migrate CLI

**Option 1: Homebrew (recommended)**

```bash
brew install golang-migrate
```

**Option 2: Install via Go**

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

> ⚠️ Note: Installing via Go may produce a "stub" CLI without database drivers.  Ensure `migrate -help` lists `postgres` or `postgresql` under **Database drivers**.

---

### 3. Run Migrations

Set your environment:

```bash
export DATABASE_URL="postgres://username:password@localhost:5432/meditation_db?sslmode=disable"
```

Run all migrations:

```bash
migrate -database "$DATABASE_URL" -path internal/db/migrations up
```

Rollback last migration:

```bash
migrate -database "$DATABASE_URL" -path internal/db/migrations down 1
```

Create a new migration:

```bash
migrate create -ext sql -dir internal/db/migrations create_users_table
```

---

### 4. Run the Server

```bash
go run cmd/server/main.go
```

Server will start with default `Gin` HTTP router and routes:

* `/health` → Health check
* `/api/v1/categories` → Category CRUD
* `/api/v1/files` → File CRUD
* `/api/v1/auth` → User registration & login

---

## API Endpoints

### Auth

* `POST /api/v1/auth/register` → Register new user
* `POST /api/v1/auth/login` → Login
* `GET /api/v1/auth/profile` → Get current user profile (JWT required)

### Categories

* `POST /api/v1/categories` → Create category
* `GET /api/v1/categories` → List categories
* `GET /api/v1/categories/:id` → Get category
* `PUT /api/v1/categories/:id` → Update category
* `DELETE /api/v1/categories/:id` → Delete category

### Files

* `POST /api/v1/files` → Upload file
* `GET /api/v1/files` → List files
* `GET /api/v1/files/:id` → Get file
* `PUT /api/v1/files/:id` → Update file
* `DELETE /api/v1/files/:id` → Delete file

---

## License

MIT License

---

## Notes

* Recommended to use Homebrew installation for golang-migrate to avoid "unknown driver" issues.
* `.env` variables must be loaded before running migrations or server.
* JWT secret must be kept safe for authentication endpoints.
