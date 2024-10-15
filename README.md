```
.
├── backend/
├── frontend/
├── docker-compose.yml
├── README.md
├── setup.sql
└── setup_db.sh
```

# Godower
Warehouse management app
- Backend: Golang
- Frontend: Svelte
- Database: Sqlite
- Dockerised the application
- Added login and register pages

## Setup using Docker
Simply run:
```
$ docker compose up --build
```
Backend will be started at http://localhost:8080

Frontend will be started at http://localhost:5173

## Local Setup
### Backend
```
$ cd backend/
$ go run .
```
Backend will be started at http://localhost:8080

### Frontend
```
$ cd frontend/
$ npm install
$ npm run dev
```
Frontend will be started at http://localhost:5173

---

**Test user credentials**:
- Email: `test@email.com`
- Password: `password`

*New user can also be registered from the Register page.*

---

`setup_db.sh and setup.sql` are scripts to populate db from the json files.
