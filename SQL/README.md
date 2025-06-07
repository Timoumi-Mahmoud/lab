##
## Intro:
This repo contains all SQL exercises from the famous website https://pgexercises.com/ which cover the following topics:
1. Basic sql
2. Joins and Subqueries
3. Modifying data
4. Aggregates 
5. Date
6. String
7. Recursive

Each sql file in this repo will contains my answer for each topic.



## set up postgres with docker
### using docker run
```bash
$ docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

$ docker start postgres15
$	docker exec -it postgres15 psql
$	docker exec -it postgres15 createdb --username=root --owner=root exercises
$ docker exec -it postgres15 dropdb exercises

$ docker cp clubdata.sql:/
$ docker exec -i postgres15 bash psql -u root -f /clubdata -d exercises

```

### using docker compose

```bash
$ Docker-compose -f DockerCompose.yaml  up -d
$ Docker-compose -f DockerCompose.yaml  down
```
## Execute queries

```bash
$ cat Part1.sql  | docker exec -i postgres15 psql -U root -d exercises

$ cat PartOne-Basic.sql | docker exec -i postgres15 psql -U root -d exercises
$ cat PartTwo-Joins-And-Subqueries.sql | docker exec -i postgres15 psql -U root -d exercises
$ cat PartThree-Modifying-Data.sql | docker exec -i postgres15 psql -U root -d exercises
```
