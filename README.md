# Сервис проведения тендеров

В проекте реализованы следующие пункты:

### 1. Проверка доступности сервера

- **Эндпоинт:** GET /ping
- **Цель:** Убедиться, что сервер готов обрабатывать запросы.
- **Ожидаемый результат:** Статус код 200 и текст "ok".

```yaml
GET /api/ping

Response:

  200 OK

  Body: ok
```

### 2. Тестирование функциональности тендеров

#### Получение списка тендеров

- **Эндпоинт:** GET /tenders
- **Описание:** Возвращает список тендеров с возможностью фильтрации по типу услуг.
- **Ожидаемый результат:** Статус код 200 и корректный список тендеров.

```yaml
GET /api/tenders

Response:

  200 OK

  Body: [ {...}, {...}, ... ]
```

#### Создание нового тендера

- **Эндпоинт:** POST /tenders/new
- **Описание:** Создает новый тендер с заданными параметрами.
- **Ожидаемый результат:** Статус код 200 и данные созданного тендера.

```yaml
POST /api/tenders/new

Request Body:

  {

    "name": "Тендер 1",

    "description": "Описание тендера",

    "serviceType": "Construction",

    "status": "Open",

    "organizationId": 1,

    "creatorUsername": "user1"

  }

Response:

  200 OK

  Body:
 
  {
    "id": 1,
    "name": "Тендер 1",
    "description": "Описание тендера",
    ...
  }
```

#### Получение тендеров пользователя

- **Эндпоинт:** GET /tenders/my
- **Описание:** Возвращает список тендеров текущего пользователя.
- **Ожидаемый результат:** Статус код 200 и список тендеров пользователя.

```yaml
GET /api/tenders/my?username=user1

Response:

  200 OK

  Body: [ {...}, {...}, ... ] 
```

#### Редактирование тендера

- **Эндпоинт:** PATCH /tenders/{tenderId}/edit
- **Описание:** Изменение параметров существующего тендера.
- **Ожидаемый результат:** Статус код 200 и обновленные данные тендера.

```yaml
PATCH /api/tenders/1/edit

Request Body:

  {

    "name": "Обновленный Тендер 1",

    "description": "Обновленное описание"

  }

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Обновленный Тендер 1",
    "description": "Обновленное описание",
    ...
  } 
```

# Запуск проекта

### Настройка приложения производится через переменные окружения

Изменить необходимые параметры в .evn

- `SERVER_ADDRESS` — адрес и порт, который будет слушать HTTP сервер при запуске. Пример: 0.0.0.0:8080.
- `POSTGRES_CONN` — URL-строка для подключения к PostgreSQL в формате postgres://{username}:{password}@{host}:{5432}/{dbname}.
- `POSTGRES_JDBC_URL` — JDBC-строка для подключения к PostgreSQL в формате jdbc:postgresql://{host}:{port}/{dbname}.
- `POSTGRES_USERNAME` — имя пользователя для подключения к PostgreSQL.
- `POSTGRES_PASSWORD` — пароль для подключения к PostgreSQL.
- `POSTGRES_HOST` — хост для подключения к PostgreSQL (например, localhost).
- `POSTGRES_PORT` — порт для подключения к PostgreSQL (например, 5432).
- `POSTGRES_DATABASE` — имя базы данных PostgreSQL, которую будет использовать приложение.


Далее добавить в базу данных таблицу (`database.sql`)

Запуск

```
sudo docker-compose up --build
```

После чего должно появиться сообщение: `Connected in PostgreSQL...`

Далее можно тестировать пункты, которые описаны выше.

Пример:

```
curl -X POST http://localhost:8080/api/tenders/new -H "Content-Type: application/json" -d '{
    "name": "Тендер 1",
    "description": "Описание тендера",
    "serviceType": "Construction",
    "organizationId": "e2746c8d-5ed3-4349-a17d-aa5e4ac55691",
    "creatorUsername": "user1"
}'
```
