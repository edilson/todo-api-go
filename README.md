# Todo API in Go

A simple **RESTful Todo API** built with **Golang**, **PostgreSQL**, and **Docker**, demonstrating CRUD operations, database integration, and unit testing.

---

## 🏗️ Features

* Create, read, update, and delete todos
* Persistent storage with PostgreSQL
* RESTful API design
* Dockerized for easy setup and deployment
* Unit tests for repository and handlers

---

## 📁 Project Structure

```
todo-api/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go
├── config/
│   └── db.go
├── models/
│   └── todo.go
├── storage/
│   └── repository.go
├── handlers/
│   └── todo.go
└── routes/
    └── routes.go
```

---

## ⚡ Requirements

* [Go 1.25+](https://golang.org/dl/)
* [Docker](https://www.docker.com/get-started)
* [Docker Compose](https://docs.docker.com/compose/install/)

---

## 🔧 Setup

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/todo-api.git
cd todo-api
```

### 2. Build and run with Docker Compose

```bash
docker-compose up --build
```

* Go API: `http://localhost:8080`
* PostgreSQL: `localhost:5432` (username: `postgres`, password: `secret`, database: `todo_db`)

---

### 3. Environment Variables (Optional)

You can set environment variables in `.env` or in `docker-compose.yml`:

```
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=todo_db
DB_PORT=5432
```

---

## 🚀 API Endpoints

| Method | Endpoint      | Description       | Body Example                                       |
| ------ | ------------- | ----------------- | -------------------------------------------------- |
| GET    | `/todos`      | Get all todos     | —                                                  |
| POST   | `/todos`      | Create a new todo | `{ "title": "Learn Go" }`                          |
| PUT    | `/todos/{id}` | Update a todo     | `{ "title": "Read GORM docs", "completed": true }` |
| DELETE | `/todos/{id}` | Delete a todo     | —                                                  |

---

## 🧪 Testing

### 1. Unit tests for repository and handlers

```bash
# Run all tests
APP_ENV=test go test ./... -v
```

* Uses **in-memory SQLite** for isolation
* Covers **CRUD operations** for repository and handlers

### 2. (To Be Done) Integration tests with Docker

* Extend tests to run against the **real PostgreSQL container** for full integration coverage.

---

## 🐳 Docker Notes

* **Dockerfile** builds a lightweight Go binary
* **docker-compose.yml** orchestrates Go API + PostgreSQL
* Persistent data stored in `postgres-data` volume

---

## ⚙️ Auto Migration

Tables are automatically migrated on API startup:

```go
config.DB.AutoMigrate(&models.Todo{})
```

---

## 💡 Future Improvements

* Add authentication (JWT)
* Add request validation
* Add pagination and filtering
* Deploy to cloud (AWS/GCP/DigitalOcean)

---

## 📜 License

MIT License

---

If you want, I can also **add a badges section and curl examples for each endpoint** to make this README even more professional and user-friendly.

Do you want me to do that?
