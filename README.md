Запуск локальных миграций:
```
migrate -source file://schema -database postgres://postgres:{password}@localhost:5436/postgres?sslmode=disable up
```

Старт контейнера с PG:
```
docker run --name todo-app-pg -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres:12.2-alpine
```