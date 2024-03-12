# Тестовое задание для golang dev в Postgres Professional
## Инструкция

### 1. Запуск
- Склонировать репозиторий
- ```make postgresinit```: Start postgres in docker container
- ```make server```: Server
- ```make swag_ui```: Swagger

## 3. Схема работы

## 4. Дополнительный функционал
Добавил дополнительные 'ручки':  
- Получение списка всех команд
- Удаление команды
- 
Добавил mock тесты.

Добавил Swagger для наглядного и удобного тестирование API (```make swag_ui```): http://localhost:8000/swagger/index.html  

Добавил migration для удобного добавления и изменения таблиц.