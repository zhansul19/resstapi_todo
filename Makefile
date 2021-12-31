postgres:	
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root todo
dropdb:
	docker exec -it postgres14 dropdb todo
.PHONY: postgres createdb dropdb 