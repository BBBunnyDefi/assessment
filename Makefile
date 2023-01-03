server:
	DATABASE_URL=postgres://root:root@localhost:5432/assessment?sslmode=disable PORT=:2565 go run server.go

unittest:
	go test -v --tags=unit ./...

ittest:
	go test -v --tags=integration ./...