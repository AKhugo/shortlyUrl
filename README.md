# 🚀 ShortlyURL – Modern URL Shortener in Go

A simple, fast, and secure URL shortening service written in Go.  
Generate short links, track clicks, and easily integrate into external tools via a RESTful API.

---

## 🔧 Features

- 🔗 Shorten URLs with a unique code
- 📊 Track clicks (IP, User-Agent, date)
- 🔐 JWT Authentication
- 🌐 RESTful API with documentation (Swagger)
- 💾 PostgreSQL database
- ⚡ Caching with Redis
- 📦 Dockerized for easy deployment
- 🧪 Unit and integration tests

---

## 🧱 Architecture

- **HTTP Server**: Go (`net/http` or Gin/Fiber)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: JWT + bcrypt
- **Logging**: Logrus
- **Deployment**: Docker, systemd

---

## 🛠️ Technologies Used

| Layer | Tools |
|------|-------|
| Backend | Go (Gin / Fiber) |
| Frontend | Next.js |
| Database | PostgreSQL |
| Cache | Redis |
| Authentication | JWT, bcrypt |
| Logging | Logrus |
| Deployment | Docker, Nginx |

---

## 📦 Installation

```bash
# Clone the project
git clone https://github.com/yourusername/shortlyurl.git 
cd shortlyurl

# Start using Docker
docker-compose up -d