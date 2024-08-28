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
![Screenshot_32](https://github.com/user-attachments/assets/a39102db-e64c-4fa6-a324-2d6a43868fa1)

Просмотр списка задач
Метод: GET /tasks
![Screenshot_33](https://github.com/user-attachments/assets/cc9a52a8-4abb-4574-a96f-272dad34c947)


Просмотр задачи
Метод: GET /tasks/{id}
![Screenshot_34](https://github.com/user-attachments/assets/c1f12c69-25aa-4cef-be61-e51bac602eda)


Обновление задачи
Метод: PUT /tasks/{id}
![Screenshot_35](https://github.com/user-attachments/assets/841786ca-be80-4a59-8023-551aa2f2934e)

Удаление задачи
Метод: DELETE /tasks/{id}
![Screenshot_36](https://github.com/user-attachments/assets/274fc558-dc0e-43c4-adda-330f15e18c6c)

### Общая информация
- Проект сделан примерно за 4-5 часов
- Материал для выполнения: курс «Backend-разработчик на Go» от Skillfactory, ютуб.
- Потерял около 30-40 минут из за невнимательности с конфигурацией sql.init файла, из за этого не появлялась таблица "tasks"
- Было ненмого непонятно с докер файлом, нужно больше практики.


MIT License.
