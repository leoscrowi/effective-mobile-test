# Subscription Service
[Swagger](docs/swagger.yaml)

# Примеры использования:
## Создание подписки:

### Request:
```http request
POST http://localhost:8080/subscriptions
Content-Type: application/json

{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025",
  "end_date": "08-2025"
}
```
### Response:
```http request
HTTP/1.1 201 Created
Content-Type: application/json
Date: Fri, 27 Feb 2026 13:47:58 GMT
Content-Length: 46

{
  "id": "e0887e7e-5e79-4bdb-afc4-6e41f2a07c0e"
}
```
### Request:
```http request
POST http://localhost:8080/subscriptions
Content-Type: application/json

{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025",
  "end_date": "08-2025"
}
```

### Response:
```http request
HTTP/1.1 201 Created
Content-Type: application/json
Date: Fri, 27 Feb 2026 13:48:53 GMT
Content-Length: 46

{
  "id": "33db9388-2a55-46d4-8978-48dad282b75e"
}
```

## Удаление подписки:
### Request:
```http request
DELETE http://localhost:8080/subscriptions/e0887e7e-5e79-4bdb-afc4-6e41f2a07c0e
```
### Response:
```http request
HTTP/1.1 204 No Content
Date: Fri, 27 Feb 2026 14:00:29 GMT

<Response body is empty>
```

## Получение подписок:
### Request:
```http request
GET http://localhost:8080/subscriptions
```

### Response:
```http request
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 27 Feb 2026 14:09:08 GMT
Content-Length: 366

{
  "subscriptions": [
    {
      "id": "2d662e56-8b03-4a11-899c-29b099996839",
      "service_name": "Yandex Plus",
      "price": 400,
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "start_date": "07-2025",
      "end_time": ""
    },
    {
      "id": "8bc5426b-021d-4d3e-8a35-f21c1a717bda",
      "service_name": "Yandex Plus",
      "price": 400,
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "start_date": "07-2025",
      "end_time": ""
    }
  ]
}
```

### Request:
```http request
GET http://localhost:8080/subscriptions/2d662e56-8b03-4a11-899c-29b099996839
```
### Response:
```http request
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 27 Feb 2026 14:09:49 GMT
Content-Length: 173

{
  "id": "2d662e56-8b03-4a11-899c-29b099996839",
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025",
  "end_time": ""
}
```

### Обновление подписки:
### Request:
```http request
PATCH http://localhost:8080/subscriptions/2d662e56-8b03-4a11-899c-29b099996839
Content-Type: application/json

{
  "service_name": "Yandex Plus Updated",
  "price": 500,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025",
  "end_date": "09-2025"
}
```
### Response:
```http request
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 27 Feb 2026 14:20:00 GMT
Content-Length: 47

{
  "message": "Subscription edited successfully"
}
```

## Получение по суммы по фильтру:
### Request (без фильтра):
```http request
### Get subscriptions amount (all subscriptions)
GET http://localhost:8080/subscriptions/amount
Content-Type: application/json

{
  "user_id": "00000000-0000-0000-0000-000000000000",
  "service_name": ""
}
```
### Response:
```http request
Content-Type: application/json
Date: Fri, 27 Feb 2026 14:27:00 GMT
Content-Length: 21

{
  "total_amount": 900
}
```

### Request (с фильтром по имени сервиса):
```http request
GET http://localhost:8080/subscriptions/amount
Content-Type: application/json

{
  "user_id": "00000000-0000-0000-0000-000000000000",
  "service_name": "Yandex Plus"
}
```
### Response:
```http request
Content-Type: application/json
Date: Fri, 27 Feb 2026 14:27:44 GMT
Content-Length: 21

{
  "total_amount": 400
}
```

### Request (с фильтром по user_id):
```http request
GET http://localhost:8080/subscriptions/amount
Content-Type: application/json

{
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "service_name": ""
}
```

### Response:
```http request
Content-Type: application/json
Date: Fri, 27 Feb 2026 14:28:24 GMT
Content-Length: 21

{
  "total_amount": 900
}

```