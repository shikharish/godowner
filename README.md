# Godower
Warehouse management app
- Backend: Golang
- Frontend: Svelte
- Database: Sqlite

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

New user can also be registerd from the register page.
