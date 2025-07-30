# SaaS Starter Go – MVP Edition

One-command starter kit to launch your SaaS today.

## ✨ What you get

- ✅ JWT auth (login / signup / roles)  
- ✅ HTML admin panel (Tailwind CDN)  
- ✅ User CRUD & role management (admin / user)  
- ✅ Password reset via email (token link)  
- ✅ Docker + hot-reload (`make dev`)  
- ✅ PostgreSQL + auto-migrations  
- ✅ Ready-to-sell MIT license  

## 🚀 Quick start

```bash
git clone https://github.com/nazgool97/startbase.git
cd startbase
cp .env.example .env
docker-compose up --build


Open browser:
Register: http://localhost:8080/signup
Login:   http://localhost:8080/admin/login
Users:   http://localhost:8080/admin/users


| Method | Endpoint         | Description      |
| ------ | ---------------- | ---------------- |
| POST   | /signup          | Create user      |
| POST   | /login           | JWT token        |
| GET    | /admin/login     | Admin login page |
| POST   | /admin/login     | Admin login form |
| GET    | /admin/users     | List users       |
| PUT    | /admin/users/:id | Change role      |
| POST   | /forgot-password | Reset email      |

Features
JWT auth & roles
HTML admin panel (Tailwind)
User CRUD & password reset
Docker + hot-reload
MIT License



**.gitignore** (Go + IDE):

```gitignore
bin/
tmp/
.env
.idea/
*.exe
.DS_Store
