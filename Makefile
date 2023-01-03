server:
	DATABASE_URL=postgres://root:root@localhost:5432/assessment?sslmode=disable PORT=:2565 go run server.go

unittest:
	go test -v --tags=unit ./...

ittest:
	go test -v --tags=integration ./...

dbstart:
	docker start assessment-db-1

dbstop:
	docker stop assessment-db-1

composeup:
	docker-compose -f docker-composer.test.yml up --build --abort-on-container-exit --exit-code-from it_tests

composedown:
	docker-compose -f docker-composer.test.yml down