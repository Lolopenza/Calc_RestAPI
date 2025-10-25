# 🧮 Calculator App — Backend

## 🧠 Overview

### The backend is developed in Go (Golang) and provides a REST API for evaluating mathematical expressions and managing history of calculations.

This backend supports:

- Evaluating expressions like: `2+3*10`
- Returning results in JSON format
- Full CRUD operations on stored expressions

---

## ⚙️ Technology Stack

### 🛠 Languages & Frameworks

| Library / Framework | Description | Role |
|-------------------|-------------|-----|
| **Go (Golang)** | High-performance backend language | Core API implementation |
| **Echo Framework**<br>(github.com/labstack/echo/v4) | Lightweight web framework | Routing & middleware |
| **Echo Middleware** | CORS + Logging support | Request logging & cross-origin support |
| **govaluate**<br>(github.com/Knetic/govaluate) | Math expression evaluation | Calculates results from strings |
| **google/uuid**<br>(github.com/google/uuid) | Universally unique IDs | Assigns IDs to each calculation |

---

## 📡 REST API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/calculations` | Get all calculations |
| `POST` | `/calculations` | Add new expression and calculate result |
| `PATCH` | `/calculations/:id` | Update existing calculation |
| `DELETE` | `/calculations/:id` | Delete calculation by ID |

---

## 💾 Data Storage

- Uses **in-memory array**: `var calculations []Calculation`
- Data is lost after server restart
- ✅ Future plan: PostgreSQL integration (via Docker)

---

## ▶️ Getting Started

```bash
go mod tidy
go run main.go
```
## Backend runs on:

```bash
http://localhost:8080
```

## Future Enhancements
- Database support (PostgreSQL via Docker)

- Dockerized backend deployment

- Improved validation / error handling

- Swagger API documentation