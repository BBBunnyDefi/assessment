server:
	DATABASE_URL=postgres://root:root@localhost:5432/assessment?sslmode=disable PORT=:2565 go run server.go