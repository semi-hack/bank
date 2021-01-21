postgres:
  docker run --name some-postgres12 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
  docker exec -it some-postgres12 createdb --username=postgres --owner=postgres bank

.PHONY: createdb