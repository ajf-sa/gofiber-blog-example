
build:
	docker build . -t gofiber-blog
run:
	docker run --network ns2020-existing-network -p 3000:3000 gofiber-blog
createdb:
	docker exec -it db01 createdb --username=admin --owner=admin gofiber-blog-api-1
dropdb:
	docker exec db01 dropdb gofiber-blog-api-1 -U admin
migration:
	@read -p "Enter migration name " name; \
	migrate create -ext sql -dir server/db/migration -seq $$name

migrate:
	migrate -path server/db/migration -database "postgresql://admin:secret@localhost:5432/gofiber-blog-api-1?sslmode=disable" -verbose up
rollback:
	migrate -path server/db/migration -database "postgresql://admin:secret@localhost:5432/gofiber-blog-api-1?sslmode=disable" -verbose down
sqlc:
	sqlc generate
pg_dump:
	docker exec db01 pg_dump -U admin -F t gofiber-blog-api-1 | gzip >backups/my_db-$(shell date +%Y-%m-%d).tar.gz
.PHONY:	build run createdb dropdb migration migrate rollback sqlc gp_dump