Go REST API (Clean Architecture + JWT Auth)

Aplikasi ini adalah backend REST API sederhana yang dibangun menggunakan Go dengan pendekatan Clean Architecture.

**Fokus utama project ini:**

- Struktur code yang scalable dan production-ready
- Implementasi authentication menggunakan JWT (access token & refresh token)
- Integrasi dengan PostgreSQL
- Containerized menggunakan Docker 

**Fitur Utama**<br>
- User Registration
- User Login (JWT Authentication)
- Access Token & Refresh Token
- Protected Endpoint (JWT Middleware)
- Role-based field (USER / ADMIN)
- REST API menggunakan Gin
- PostgreSQL sebagai database

**Arsitektur**<br>
Project ini menggunakan pendekatan layered clean architecture:<br>

cmd/            → entry point aplikasi<br>
router/         → routing (endpoint mapping)<br>
delivery/http   → HTTP handler (Gin)<br>
usecases/       → business logic<br>
repository/     → akses database<br>
entities/       → model / domain<br>
configs/        → koneksi database<br>

**Flow aplikasi:**<br>

Request → Router → Handler → Usecase → Repository → Database

**Tech Stack**<br>
Go (Golang)
Gin (HTTP Framework)
PostgreSQL
Docker & Docker Compose
JWT (Authentication)

**Endpoint Utama**<br>
**Register**<br>
POST /register

**Login**<br>
POST /login

**Get Users (Protected)** <br>
GET /users
Authorization: Bearer <token>

**Cara Menjalankan**
1. Jalankan dengan Docker
go run cmd/main.go
2. Akses API
http://localhost:8080

**Catatan**<br>
- Password disimpan dalam bentuk hash (bcrypt)
- Token memiliki masa berlaku (expired)
- Refresh token disimpan di database
  
**Tujuan Project**<br>

**Project ini dibuat sebagai:**<br>

- pembelajaran backend Go dari dasar hingga production style
- implementasi real authentication system (JWT + refresh token)
- fondasi untuk pengembangan API yang lebih kompleks (RBAC, OAuth2, dll)
