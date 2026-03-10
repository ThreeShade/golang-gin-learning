# 🐹 Go Backend Developer — 35-Day Study Plan

> **Stack:** Go · Gin · SQLC · PostgreSQL · RabbitMQ
> **Time:** 3–4 hours/day · **Goal:** Get hired as a Go backend developer
> **Background:** Strong in React / Next.js / JavaScript

---

## 📋 How to Use This

- Check off each task as you complete it
- Each day has a **morning block** (learn) and **evening block** (build)
- Every day ends with a concrete output — don't move on until you have it
- Push your code to GitHub **every day** — employers see your commit history

---

## 📅 Overview

| Week | Focus | Days |
|------|-------|------|
| Week 1 | Go Fundamentals | 1–7 |
| Week 2 | REST API with Gin | 8–14 |
| Week 3 | PostgreSQL + SQLC | 15–21 |
| Week 4 | RabbitMQ + Production Patterns | 22–28 |
| Week 5 | Capstone Project + Job Prep | 29–35 |

---

## 🟦 Week 1 — Go Fundamentals

> Your goal this week: get comfortable with Go syntax before touching any framework.

### Day 1 — Go Setup & Basics

**Morning**
- [x] Install Go from golang.org/dl and set up VS Code with the official Go extension
- [x] Complete `tour.golang.org` — Basics section (variables, types, functions)
- [x] Read: how Go differs from JS — static types, zero values, no `undefined`, no `null`

**Evening**
- [ ] Write a hello world program and run it with `go run`
- [ ] Practice variables, type conversions, and basic functions in 5 small programs
- [ ] Write a function with multiple return values: `func divide(a, b float64) (float64, error)`

📦 **Output:** 5 working `.go` files committed to GitHub

---

### Day 2 — Control Flow & Structs

**Morning**
- [ ] Go by Example: `if/else`, `for` loops (Go only has `for` — no `while`), `switch`
- [ ] Read about structs — think of them as JS objects with a fixed shape and type
- [ ] Understand methods on structs: `func (s Student) AverageGrade() float64`

**Evening**
- [ ] Build a `Student` struct with fields: `Name string`, `Age int`, `Grades []float64`
- [ ] Add methods: `AverageGrade()`, `IsPassing()`, `String()`
- [ ] Print a formatted report using `fmt.Printf`

📦 **Output:** Student struct mini-project

---

### Day 3 — Interfaces & Error Handling

**Morning**
- [ ] Read Go interfaces — duck typing without explicit `implements` keyword
- [ ] Study `io.Reader` and `io.Writer` as real-world interface examples
- [ ] Understand the `(value, error)` return pattern — this is everywhere in Go

**Evening**
- [ ] Write 3 functions that return `(result, error)`
- [ ] Create a custom error type: `type ValidationError struct { Field string; Message string }`
- [ ] Practice wrapping errors: `fmt.Errorf("getUserById: %w", err)`

📦 **Output:** Error handling exercise

---

### Day 4 — Slices, Maps & Pointers

**Morning**
- [ ] Slices vs arrays — compare with JS arrays (`append`, `len`, `cap`)
- [ ] Maps — compare with JS objects (`make`, `delete`, checking if key exists with `ok`)
- [ ] Pointers basics: `&variable` gives address, `*pointer` dereferences — why Go needs them

**Evening**
- [ ] Build a shopping cart: `map[string]Item` where `Item` is a struct with `Name`, `Price`, `Qty`
- [ ] Add functions: `AddItem`, `RemoveItem`, `Total() float64`, `PrintCart()`
- [ ] Use a pointer receiver on a method so it modifies the cart in place

📦 **Output:** Shopping cart CLI app

---

### Day 5 — Packages & Modules

**Morning**
- [ ] Understand `go.mod` and `go.sum` — compare with `package.json` and `package-lock.json`
- [ ] One folder = one package. Exported = UpperCase, unexported = lowerCase
- [ ] Read how to import local packages vs external packages (`go get`)

**Evening**
- [ ] Create a multi-file project: `main.go` + `calculator/calc.go` as a separate package
- [ ] Build a CLI unit converter: reads from stdin, converts kg↔lb, km↔miles, °C↔°F
- [ ] Use `fmt`, `os`, `bufio`, `strconv` from the standard library

📦 **Output:** CLI unit converter (multi-file project)

---

### Day 6 — Review & Build

**Morning**
- [ ] Re-read any weak areas from Days 1–5
- [ ] Read Effective Go (official docs) — focus on the Errors and Interfaces sections
- [ ] Write down 5 things that confused you this week and look each one up

**Evening**
- [ ] Build a terminal quiz app that uses: structs, interfaces, maps, error handling, and packages
- [ ] Questions are stored in a slice of structs, quiz is scored at the end
- [ ] Clean it up and push to GitHub with a README

📦 **Output:** Quiz CLI app on GitHub

---

### Day 7 — Rest + Light Review

**Morning**
- [ ] Read Go Proverbs at `go-proverbs.github.io` — internalize the Go philosophy
- [ ] Skim the `strings`, `strconv`, and `math` standard library docs
- [ ] Watch one 30-min Go intro video on YouTube to reinforce the week

**Evening**
- [ ] No heavy coding — skim the Go FAQ on the official site
- [ ] Write 5 answers to: "How is Go different from JavaScript?"
- [ ] Plan Week 2 — Gin is like Express, this week moves fast for you

📦 **Output:** Notes file with Go vs JS comparison

---

## 🟩 Week 2 — REST API with Gin

> Gin feels like Express. Your frontend background makes this week fast.

### Day 8 — HTTP Basics + First Gin Server

**Morning**
- [ ] Build a raw HTTP server with `net/http` first — no framework, understand the base
- [ ] Understand `http.HandlerFunc`, `http.ResponseWriter`, `*http.Request`
- [ ] Install Gin: `go get github.com/gin-gonic/gin`

**Evening**
- [ ] Rewrite the server in Gin — compare it side-by-side with the raw version
- [ ] Add routes: `GET /ping`, `GET /hello/:name`, `POST /echo`
- [ ] Test every route with curl or Postman

📦 **Output:** Basic Gin server with 3 routes

---

### Day 9 — Routing, Params & Middleware

**Morning**
- [ ] Route groups: `r.Group("/api/v1")` — like Express Router
- [ ] Path params: `c.Param("id")`, query params: `c.Query("page")`
- [ ] Study how Gin middleware works — same concept as Express, different syntax

**Evening**
- [ ] Write a logger middleware: logs method, path, status code, duration
- [ ] Write an auth middleware: checks for `Authorization` header, returns 401 if missing
- [ ] Apply both to a route group and test

📦 **Output:** Gin app with 2 custom middlewares

---

### Day 10 — JSON Binding & Validation

**Morning**
- [ ] Struct tags: `json:"name" binding:"required,min=3"` — how they work
- [ ] `c.ShouldBindJSON` vs `c.BindJSON` — know the difference (one aborts, one doesn't)
- [ ] Returning consistent error responses: `c.JSON(400, gin.H{"error": "..."})`

**Evening**
- [ ] Build complete in-memory CRUD for `/users`: GET list, GET one, POST, PUT, DELETE
- [ ] Use a struct with validation tags for the request body
- [ ] Test all routes with valid and invalid data in Postman

📦 **Output:** Full in-memory CRUD API

---

### Day 11 — Error Handling in APIs

**Morning**
- [ ] Design a consistent error response format: `{ "error": "message", "code": 400 }`
- [ ] Custom error types: `type APIError struct { Message string; Status int }`
- [ ] Read about Gin's built-in recovery middleware for handling panics

**Evening**
- [ ] Refactor Day 10 CRUD to use centralized error handling
- [ ] Every route returns the same error format — no inconsistent responses
- [ ] Test edge cases: invalid IDs, missing fields, not found, wrong types

📦 **Output:** Polished, error-handled CRUD API

---

### Day 12 — Project: Todo API

**Morning**
- [ ] Plan on paper: routes, structs, and data flow for a Todo API
- [ ] Entities: `User`, `Task`. Define all JSON request/response shapes
- [ ] Decide which routes need auth middleware and which are public

**Evening**
- [ ] Implement the full Todo API with Gin (no DB yet, in-memory store)
- [ ] Routes: `GET /tasks`, `POST /tasks`, `PUT /tasks/:id`, `DELETE /tasks/:id`
- [ ] Add pagination: `?page=1&limit=10`

📦 **Output:** Todo API with pagination

---

### Day 13 — Refactor & Polish

**Morning**
- [ ] Clean folder structure: separate `handlers/`, `routes/`, `models/` packages
- [ ] Read about the Repository pattern in Go — how to decouple data from handlers
- [ ] Every input to POST/PUT must be validated — no unvalidated writes

**Evening**
- [ ] Refactor the Todo API into a clean folder structure
- [ ] Write a reusable paginated response helper function
- [ ] Test every edge case: empty list, invalid ID, missing required fields

📦 **Output:** Clean, structured Gin project

---

### Day 14 — Rest + Interview Prep

**Morning**
- [ ] Study: What is middleware? What is a handler? How does Gin differ from `net/http`?
- [ ] Read: how Go handles concurrency in HTTP servers (each request gets a goroutine)
- [ ] Review your Week 2 code — can you explain every line out loud?

**Evening**
- [ ] Write answers to these 5 questions in a notes file:
  - [ ] What is the difference between `net/http` and Gin?
  - [ ] How does middleware work in Gin?
  - [ ] How do you validate request bodies in Go?
  - [ ] What are struct tags and how are they used with JSON?
  - [ ] How do you handle errors globally in a Gin API?

📦 **Output:** Interview notes file

---

## 🟨 Week 3 — PostgreSQL + SQLC

> You know basic SQL. This week is about wiring it into Go the right way.

### Day 15 — Docker + PostgreSQL Setup

**Morning**
- [ ] Install Docker Desktop if not already installed
- [ ] Write a `docker-compose.yml` with a Postgres service
- [ ] Connect to it with `psql` and verify it works

**Evening**
- [ ] Practice SQL you'll actually use: `CREATE TABLE`, `INSERT`, `SELECT`, `UPDATE`, `DELETE`
- [ ] Practice JOINs: write a query that joins two tables
- [ ] Create your first real database: `todoapp` with `users` and `tasks` tables

📦 **Output:** Postgres running in Docker with a real schema

---

### Day 16 — Database Migrations

**Morning**
- [ ] Install `golang-migrate`: `go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- [ ] Understand up/down migrations — up creates, down reverts
- [ ] Read why migrations matter: version-controlled schema changes

**Evening**
- [ ] Write `000001_init_schema.up.sql` with `users` and `tasks` tables
- [ ] Write the corresponding `.down.sql` that drops both tables
- [ ] Run `migrate up`, verify in psql, then run `migrate down` and back up

📦 **Output:** Schema with working migrations

---

### Day 17 — SQLC Introduction

**Morning**
- [ ] Install sqlc: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- [ ] Read how SQLC works: you write SQL → it generates type-safe Go code
- [ ] Configure `sqlc.yaml` for your todo project

**Evening**
- [ ] Write 4 queries in `queries/users.sql`: `GetUser`, `CreateUser`, `ListUsers`, `DeleteUser`
- [ ] Write 4 queries in `queries/tasks.sql`: `GetTask`, `CreateTask`, `ListTasks`, `UpdateTask`
- [ ] Run `sqlc generate` and read every generated file — understand what was created

📦 **Output:** Generated Go DB code from SQLC

---

### Day 18 — Connect Gin + SQLC

**Morning**
- [ ] Install pgx driver: `go get github.com/jackc/pgx/v5`
- [ ] Create a DB connection pool: `pgxpool.New(ctx, connString)`
- [ ] Pass the DB pool to handlers via dependency injection (not global variables)

**Evening**
- [ ] Replace the in-memory store in your Todo API with real Postgres queries via SQLC
- [ ] Test every endpoint — it should work exactly the same, but now data persists
- [ ] Try restarting the server — confirm data is still there

📦 **Output:** Todo API backed by real PostgreSQL

---

### Day 19 — Advanced Queries & Auth Prep

**Morning**
- [ ] Write SQLC queries with filters: `WHERE user_id = $1 AND completed = $2`
- [ ] Add pagination at the DB level: `LIMIT $1 OFFSET $2`
- [ ] Read about password hashing: install `golang.org/x/crypto/bcrypt`

**Evening**
- [ ] Add `password_hash` to the `users` table via a new migration
- [ ] Write `CreateUser` that hashes the password before storing
- [ ] Write `AuthenticateUser` that fetches user and compares hash with `bcrypt.CompareHashAndPassword`

📦 **Output:** Auth-ready user queries

---

### Day 20 — Database Transactions

**Morning**
- [ ] What are transactions and when to use them: ACID properties explained simply
- [ ] How to use transactions with pgx: `pool.Begin(ctx)`, `tx.Commit()`, `tx.Rollback()`
- [ ] SQLC transactions: passing `*sql.Tx` instead of `*sql.DB`

**Evening**
- [ ] Implement a feature that needs a transaction: e.g. "complete task and update user score"
- [ ] Make it fail partway through and verify the rollback works
- [ ] Write a `Store` struct that wraps the SQLC queries and adds transaction support

📦 **Output:** Transaction example in your app

---

### Day 21 — Rest + DB Interview Prep

**Morning**
- [ ] Review your whole Week 3 codebase — understand every file
- [ ] Re-read the SQLC docs for anything you skipped
- [ ] Study: what is N+1 query problem and how to avoid it

**Evening**
- [ ] Write answers to these 5 questions:
  - [ ] What is SQLC? How is it different from an ORM like Prisma or GORM?
  - [ ] What are database migrations and why are they important?
  - [ ] What is a connection pool and why does it matter?
  - [ ] When would you use a database transaction?
  - [ ] What is the difference between `pgx` and `database/sql`?

📦 **Output:** DB interview notes

---

## 🟪 Week 4 — RabbitMQ + Production Patterns

> This week makes your app production-ready and event-driven.

### Day 22 — JWT Authentication

**Morning**
- [ ] How JWT works: header.payload.signature — what each part contains
- [ ] Install: `go get github.com/golang-jwt/jwt/v5`
- [ ] JWT in Gin: generate on login, validate in middleware, extract claims

**Evening**
- [ ] Implement `POST /login` — validates credentials, returns JWT
- [ ] Write a JWT middleware that validates the token and sets user ID in context
- [ ] Protect all `/tasks` routes with the JWT middleware and test with Postman

📦 **Output:** JWT auth working end-to-end

---

### Day 23 — RabbitMQ Setup & Basics

**Morning**
- [ ] Add RabbitMQ to your `docker-compose.yml` (use the `rabbitmq:3-management` image)
- [ ] Install the Go client: `go get github.com/rabbitmq/amqp091-go`
- [ ] Understand: Producer → Exchange → Queue → Consumer

**Evening**
- [ ] Write a basic producer: connects to RabbitMQ, declares a queue, publishes "hello world"
- [ ] Write a basic consumer: connects, listens on the queue, prints each message
- [ ] Run both — see the message flow in real time and in the RabbitMQ management UI

📦 **Output:** Working producer and consumer

---

### Day 24 — RabbitMQ + API Integration

**Morning**
- [ ] How to publish a message from inside a Gin handler
- [ ] Message payload design: use JSON `{"event": "task.created", "taskId": "...", "userId": "..."}`
- [ ] Consumer as a separate goroutine running alongside your Gin server

**Evening**
- [ ] Publish a `task.created` event when a new task is created via `POST /tasks`
- [ ] Write a consumer goroutine that processes the event and logs: "Task X created by User Y"
- [ ] Start the consumer in `main.go` with `go startConsumer()` before the Gin server

📦 **Output:** Event-driven task creation flow

---

### Day 25 — Config & Environment Variables

**Morning**
- [ ] Why hardcoded strings are bad in production — secrets in code = security risk
- [ ] Install Viper: `go get github.com/spf13/viper` — or use the simpler `os.Getenv`
- [ ] `.env` file for local dev, environment variables for production

**Evening**
- [ ] Create a `config/config.go` that loads: `DB_URL`, `RABBITMQ_URL`, `JWT_SECRET`, `PORT`
- [ ] Replace every hardcoded string in your app with config values
- [ ] Add `.env` to `.gitignore` and create `.env.example` with placeholder values

📦 **Output:** Config-driven app with `.env.example`

---

### Day 26 — Testing in Go

**Morning**
- [ ] Go testing basics: `_test.go` files, `func TestXxx(t *testing.T)`
- [ ] Table-driven tests — the idiomatic Go way to test many cases
- [ ] Mocking in Go: interfaces make mocking easy — no extra library needed

**Evening**
- [ ] Write unit tests for your password hashing functions
- [ ] Write table-driven tests for your JWT generate/validate functions
- [ ] Write at least one integration test for a SQLC query using a test database

📦 **Output:** Test suite with unit and integration tests

---

### Day 27 — Project Structure & Swagger Docs

**Morning**
- [ ] Read about the standard Go project layout: `/cmd`, `/internal`, `/pkg`
- [ ] `/cmd/api/main.go` — entry point. `/internal/` — private app code. `/pkg/` — reusable utilities
- [ ] Install swaggo: `go install github.com/swaggo/swag/cmd/swag@latest`

**Evening**
- [ ] Refactor your project into the `/cmd` and `/internal` structure
- [ ] Add Swagger comments to all your Gin handlers
- [ ] Run `swag init` and serve the Swagger UI at `/swagger/*any`

📦 **Output:** Clean project structure + Swagger docs

---

### Day 28 — Rest + Interview Prep

**Morning**
- [ ] Review all 4 weeks — read through every file you've written
- [ ] List your weak spots and spend 30 minutes on each
- [ ] Study: how goroutines differ from async/await in JavaScript

**Evening**
- [ ] Write answers to these 5 questions:
  - [ ] What is a message broker? Why use one instead of direct API calls?
  - [ ] What happens if a RabbitMQ consumer crashes mid-processing?
  - [ ] How do goroutines differ from async/await in JavaScript?
  - [ ] How would you structure a Go project for a team?
  - [ ] What is JWT and how does it work?

📦 **Output:** Week 4 interview notes

---

## 🟥 Week 5 — Capstone Project + Job Prep

> Build one real, complete project to show employers.

### Day 29–30 — Project Planning

**Morning**
- [ ] Design a **Job Board API** — realistic, hirable, and uses all 5 technologies
- [ ] Entities: `User`, `Company`, `Job`, `Application`
- [ ] Write the full schema on paper with all foreign keys

**Evening**
- [ ] Define every API route: auth, jobs CRUD, applications, company profiles
- [ ] Plan which RabbitMQ events to publish (e.g. `application.submitted`)
- [ ] Create the GitHub repo, write the initial README, push the empty project

📦 **Output:** GitHub repo + project spec

---

### Day 31–33 — Build the API

**Morning (each day)**
- [ ] Day 31: Auth (register/login with JWT) + Company endpoints
- [ ] Day 32: Jobs CRUD with pagination and filters (by location, salary, type)
- [ ] Day 33: Applications endpoint + RabbitMQ event on new application

**Evening (each day)**
- [ ] Day 31: Write migrations and SQLC queries for users and companies
- [ ] Day 32: Write migrations and SQLC queries for jobs, add search/filter
- [ ] Day 33: Wire in RabbitMQ consumer that logs application events

📦 **Output:** Fully working backend API

---

### Day 34–35 — Polish & Ship

**Morning**
- [ ] Add JWT middleware to all protected routes
- [ ] Write tests for at least 3 critical functions
- [ ] Add Swagger docs to all endpoints

**Evening**
- [ ] Write a detailed README: what it does, tech stack, how to run locally, API endpoints
- [ ] Create a clean `docker-compose.yml` so anyone can run the whole stack with one command
- [ ] Deploy to Railway or Render (free tier) — get a live URL to put on your resume
- [ ] Push everything to GitHub with a clean commit history

📦 **Output:** Live, deployed project with public GitHub URL

---

## 🎯 Capstone Project: Job Board API — Feature Checklist

### Auth
- [ ] `POST /api/v1/register` — create user account
- [ ] `POST /api/v1/login` — returns JWT token
- [ ] JWT middleware protecting all non-public routes

### Companies
- [ ] `GET /api/v1/companies` — paginated list
- [ ] `POST /api/v1/companies` — create company (auth required)
- [ ] `GET /api/v1/companies/:id` — company detail

### Jobs
- [ ] `GET /api/v1/jobs` — paginated list with filters
- [ ] `POST /api/v1/jobs` — create job posting (auth required)
- [ ] `GET /api/v1/jobs/:id` — job detail
- [ ] `PUT /api/v1/jobs/:id` — update job (auth required)
- [ ] `DELETE /api/v1/jobs/:id` — delete job (auth required)

### Applications
- [ ] `POST /api/v1/jobs/:id/apply` — apply to a job (publishes RabbitMQ event)
- [ ] `GET /api/v1/applications` — list my applications (auth required)

### Infrastructure
- [ ] RabbitMQ consumer that processes `application.submitted` events
- [ ] All config via environment variables
- [ ] Swagger/OpenAPI docs at `/swagger/`
- [ ] `docker-compose.yml` with Postgres + RabbitMQ + the API
- [ ] `.env.example` file

### Quality
- [ ] Unit tests for auth and validation logic
- [ ] Integration test for at least one DB query
- [ ] Clean project structure (`/cmd`, `/internal`)
- [ ] Detailed README with setup instructions
- [ ] Deployed with a live URL

---

## 💼 Job Search Checklist

### GitHub Profile
- [ ] Capstone project pinned with a clear README
- [ ] Commit activity visible (daily commits from Week 1 pay off here)
- [ ] At least one smaller project from earlier in the plan also pinned
- [ ] `.env.example` present — no secrets committed

### Resume
- [ ] List: Go, Gin, SQLC, PostgreSQL, RabbitMQ, Docker, JWT
- [ ] Also list: React, Next.js, JavaScript — full-stack is a plus
- [ ] Link to GitHub and the live project URL
- [ ] 2–3 bullet points describing the capstone project and its scale

### Job Search
- [ ] Search `golang.cafe` — Go-specific job board
- [ ] Search LinkedIn: "Go developer", "Golang backend", "Go engineer"
- [ ] Target: Junior/Mid Backend Engineer (Go), Full-Stack (Go + React), API Engineer
- [ ] Companies that use Go: Docker, Cloudflare, HashiCorp, Twitch, Uber, Monzo

### Interview Prep
- [ ] Can explain all 5 technologies in plain English
- [ ] Can walk through the capstone project end-to-end
- [ ] Practiced answering: "Why Go over Node.js for backend?"
- [ ] Reviewed: goroutines, interfaces, error handling, REST design
- [ ] Did at least one mock interview (record yourself or use a friend)

---

## 🔗 Essential Resources

| Resource | Purpose |
|----------|---------|
| [tour.golang.org](https://tour.golang.org) | Complete this in Week 1 — the official Go tour |
| [gobyexample.com](https://gobyexample.com) | Daily reference throughout all 5 weeks |
| [go-proverbs.github.io](https://go-proverbs.github.io) | Read on Day 7 — internalize Go philosophy |
| [sqlc.dev](https://sqlc.dev) | SQLC docs and playground |
| [pkg.go.dev](https://pkg.go.dev) | Official Go package documentation |
| [golang.cafe](https://golang.cafe) | Go-specific job board |
| Air (`github.com/air-verse/air`) | Hot reload for Go — like nodemon |

---

## 💡 Tips

- **Code every single day** — even 30 minutes on slow days keeps the momentum
- **Don't understand everything before coding** — write broken code, debug it, learn from errors
- **Use Docker Compose from Day 1** — never install Postgres or RabbitMQ natively
- **Your React background is an asset** — many companies want Go devs who can talk to frontend teams
- **Commit to GitHub daily** — your commit history is visible to employers

---

*35 days. Start today.*
