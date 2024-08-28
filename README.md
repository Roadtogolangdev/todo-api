# Todo API
Проект REST API для управления задачами в формате to-do list. Реализованы создача, просмотр списка задач, обновление и удаление задач.

## Стек технологий

- Go 
- PostgreSQL 
- Docker (для упрощеной работы с базой данных)
- GORM 

## Установка и запуск

### 1. Клонируйте репозиторий

https://github.com/Roadtogolangdev/todo-api.git

### 2. Настройка Docker
Создание и запуск контейнеров

docker-compose up --build
На этом этапе у меня были трудности с SQL поэтому я вынес создание таблицы в отдельный init.sql файл.

### 3. Запуск приложения
Зайдите в корневую директорию проекта и запустите его командой: 
go run main.go

Сообщение "Успешное подключение к БД" подтвердит запуск на порту 8080. 

Конфигурация
Строка подключения к базе данных находится в database/database.go. 
connect := "user=postgres password=12345 dbname=basa sslmode=disable"

### Работа с API
Создание задачи
Метод: POST /tasks
![Screenshot_32](https://github.com/user-attachments/assets/05a72d1b-0f61-4c25-b2e8-e05df5373d16)

Просмотр списка задач
Метод: GET /tasks
![Screenshot_33](https://github.com/user-attachments/assets/68b84ce7-2945-4c62-b57d-4c94ef49b9e6)


Просмотр задачи
Метод: GET /tasks/{id}
![Screenshot_34](https://github.com/user-attachments/assets/8ed54d00-7bc8-4c64-b854-e530ca167f3d)


Обновление задачи
Метод: PUT /tasks/{id}
![Screenshot_35](https://github.com/user-attachments/assets/45d87e67-9900-4548-9f50-9391e5f754b1)

Удаление задачи
Метод: DELETE /tasks/{id}
![Screenshot_36](https://github.com/user-attachments/assets/4f278ee7-74b5-4c81-9946-9a384fded9cb)



Лицензия
Этот проект лицензирован под MIT License.
