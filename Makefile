postgres:
	docker run --name postgresloggertg -p 5436:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=loggertg -d postgres

createdb:
	docker exec -it postgresloggertg createdb --username=root --owner=root Webloggertg

dropdb:
	docker exec -it postgresloggertg dropdb Webloggertg

migrateup:
	migrate -path pkg/common/database/migrations -database "postgresql://root:1234@localhost:5436/Webloggertg?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/common/database/migrations -database "postgresql://root:1234@localhost:5436/Webloggertg?sslmode=disable" -verbose down

# .PHONY: test
# test:
# 	go test ./pkg/common/database/sqlc -cover


# redis:
# 	docker run -d -p 6379:6379 --name market-place redis

# redisstart:
# 	docker start market-place

# redisstop:
# 	docker stop market-place

# run:
# 	go run cmd/app/main.go