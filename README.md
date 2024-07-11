# HTTP Proxy Server

HTTP Proxy Server — это HTTP-сервер, который проксирует запросы к сторонним сервисам и возвращает ответы в формате JSON. Запросы и ответы сохраняются локально в памяти.

Линк на коллекцию в Postman - https://www.postman.com/speeding-equinox-936930/workspace/htpp-proxy-server/collection/28284936-1369f3e5-8515-4a0b-99cb-4fe29f6718b0?action=share&creator=28284936
## Оглавление

- [Функциональность](#функциональность)
- [Установка](#установка)
- [Запуск](#запуск)
- [Использование](#использование)
- [Примеры запросов](#примеры-запросов)
- [Контейнеризация и деплой](#контейнеризация-и-деплой)
- [Автор](#автор)

## Функциональность

- Обработка входящих HTTP-запросов от клиентов.
- Извлечение метода, URL и заголовков из тела запроса.
- Отправка соответствующего HTTP-запроса к стороннему сервису.
- Получение ответа от стороннего сервиса и возвращение результата клиенту в формате JSON.
- Сохранение запросов и ответов локально в памяти.
- Валидация входных данных.

## Установка

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/yermanovberik/http-proxy-server
    cd http-proxy-server
    ```

2. Установите зависимости и соберите проект:

    ```sh
    go mod download
    go build -o bin/http-proxy-server
    ```
3.Собрание проекта:
 ```make build
    ```

4.Запуск проекта:
 ``` make run
    ```



Запрос 1: GET запрос
URL: https://http-proxy-server-u0dt.onrender.com/proxy

Метод: POST

Тело запроса:
{
    "method": "GET",
    "url": "https://jsonplaceholder.typicode.com/posts/1"
}



Ответ:

 {
    "id": "1720678827142878949",
    "status": 200,
    "headers": {
        "Access-Control-Allow-Credentials": "true",
        "Age": "4339",
        "Alt-Svc": "h3=\":443\"; ma=86400",
        "Cache-Control": "max-age=43200",
        "Cf-Cache-Status": "HIT",
        "Cf-Ray": "8a16b00e3a6540c0-SIN",
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Thu, 11 Jul 2024 06:20:27 GMT",
        "Etag": "W/\"124-yiKdLzqO5gfBrJFrcdJ8Yq0LGnU\"",
        "Expires": "-1",
        "Nel": "{\"report_to\":\"heroku-nel\",\"max_age\":3600,\"success_fraction\":0.005,\"failure_fraction\":0.05,\"response_headers\":[\"Via\"]}",
        "Pragma": "no-cache",
        "Report-To": "{\"group\":\"heroku-nel\",\"max_age\":3600,\"endpoints\":[{\"url\":\"https://nel.heroku.com/reports?ts=1712085044&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=5bmno5H5ge72f6ZZ4%2B52pSdbReNAenxXqSzCrWHAqk4%3D\"}]}",
        "Reporting-Endpoints": "heroku-nel=https://nel.heroku.com/reports?ts=1712085044&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=5bmno5H5ge72f6ZZ4%2B52pSdbReNAenxXqSzCrWHAqk4%3D",
        "Server": "cloudflare",
        "Vary": "Origin, Accept-Encoding",
        "Via": "1.1 vegur",
        "X-Content-Type-Options": "nosniff",
        "X-Powered-By": "Express",
        "X-Ratelimit-Limit": "1000",
        "X-Ratelimit-Remaining": "999",
        "X-Ratelimit-Reset": "1712085075"
    },
    "length": 292
}


Запрос 2: DELETE запрос
URL: https://http-proxy-server-u0dt.onrender.com/proxy

Метод: POST

Тело запроса:

{
    "method": "DELETE",
    "url": "https://jsonplaceholder.typicode.com/posts/1"
}

Ответ:

 {
    "id": "1720672257777060000",
    "status": 200,
    "headers": {
        "Access-Control-Allow-Credentials": "true",
        "Alt-Svc": "h3=\":443\"; ma=86400",
        "Cache-Control": "no-cache",
        "Cf-Cache-Status": "DYNAMIC",
        "Cf-Ray": "8a160fadafb9c270-VIE",
        "Content-Length": "2",
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Thu, 11 Jul 2024 04:30:58 GMT",
        "Etag": "W/\"2-vyGp6PvFo4RvsFtPoIWeCReyIC8\"",
        "Expires": "-1",
        "Nel": "{\"report_to\":\"heroku-nel\",\"max_age\":3600,\"success_fraction\":0.005,\"failure_fraction\":0.05,\"response_headers\":[\"Via\"]}",
        "Pragma": "no-cache",
        "Report-To": "{\"group\":\"heroku-nel\",\"max_age\":3600,\"endpoints\":[{\"url\":\"https://nel.heroku.com/reports?ts=1720672258&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=b%2Fy27gkxVosKNlMJ5JN%2BGaVIzB05UWxXxtatk2MKqOc%3D\"}]}",
        "Reporting-Endpoints": "heroku-nel=https://nel.heroku.com/reports?ts=1720672258&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=b%2Fy27gkxVosKNlMJ5JN%2BGaVIzB05UWxXxtatk2MKqOc%3D",
        "Server": "cloudflare",
        "Vary": "Origin, Accept-Encoding",
        "Via": "1.1 vegur",
        "X-Content-Type-Options": "nosniff",
        "X-Powered-By": "Express",
        "X-Ratelimit-Limit": "1000",
        "X-Ratelimit-Remaining": "998",
        "X-Ratelimit-Reset": "1720672262"
    },
    "length": 2
}
