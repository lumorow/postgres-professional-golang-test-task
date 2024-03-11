# Тестовое задание для golang dev в Postgres Professional
## Инструкция
### 1. Установка механизма миграции
Установка для Linux:
- ```curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash```
- ```sudo apt-get update```
- ```sudo apt-get install migrate```

Установка для MacOS:
- ```brew install golang-migrate```

Установка для Windows:
- ```scoop install migrate```

*optionally*:
- ```go install github.com/golang/mock/mockgen@v1.6.0```


### 2. Запуск
- Склонировать репозиторий
- ```go mod tidy```: Установить зависимости
- ```make postgresinit```: Postgres Docker
- ```make migrationup```: Migrate table
- ```make server```: Server

## 3. Схема работы

## 4. Дополнительный функционал
- Добавил Swagger для наглядного и удобного тестирование API: http://localhost:8000/swagger/index.html