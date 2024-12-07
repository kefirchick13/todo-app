# Разворачивание приложения через docker compose 
docker compose up --build

# Поднятие миграций 
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5438/postgres?sslmode=disable' up


// Подключение к докер контейнеру
// On Windows CMD (not switching to bash):
// docker exec -it <container-id> /bin/sh
// On Windows CMD (after switching to bash):
// docker exec -it <container-id> //bin//sh
// or
// winpty docker exec -it <container-id> //bin//sh
// On Git Bash:
// winpty docker exec -it <container-id> //bin//sh


// Поднятие контейнера

// docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5438:5432 -d --rm postgres

// Поднятие баз

// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5438/postgres?sslmode=disable' up

// Дроп

// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5438/postgres?sslmode=disable' down

// запуск контейнера
// docker start 8a13bc9cbee328ac098df181c551f5ed6758132ec5ecca0c558f5e262aa43d84

// Просмотр контейнерров(всех)
// docker ps -a

// docker exec -it c96536d845bd80f7fb1cee2a244c5dd795b4921a6f088d99b782d856b80b7476 /bin/bash подключение к контейнеру базы

// psql -U postgres Вход в базу через утилиту

// \d Просмотр таблиц 

// exit выход из контейнера
