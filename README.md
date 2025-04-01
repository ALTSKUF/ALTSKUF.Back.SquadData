# SquadData
## Запуск
Для работы требуется запущенная база данных и сервер RabbitMQ по адресам, указанным в конфиге.
### Вручную
```
git clone https://github.com/ALTSKUF/ALTSKUF.Back.SquadData
cd ALTSKUF.Back.SquadData/src
go build -o ../app
cd ..
./app
```
### docker-compose
```
services:
  db:
    image: postgres:17.4-alpine3.21
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: mypassword
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data 

  app:
    build: 
      dockerfile: app.Dockerfile
    ports:
      - "8000:8000"
    depends_on:
      rabbitmq:
        condition: service_healthy
      db:
        condition: service_started
    environment:
      SQUAD_APP_PROFILE: release
      SQUAD_DB_HOST: db 
      SQUAD_DB_USER: postgres
      SQUAD_DB_PASSWORD: mypassword
      SQUAD_DB_NAME: postgres
      SQUAD_DB_PORT: 5432
      SQUAD_RMQ_USER: user
      SQUAD_RMQ_PASSWORD: password

  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      RABBITMQ_DEFAULT_USER: user
      RABBITMQ_DEFAULT_PASS: password
    healthcheck:
      test: ["CMD", "rabbitmq-diagnostics", "-q", "check_port_connectivity"]
      interval: 5s
      timeout: 5s
      retries: 3

volumes:
  postgres_data:
```

## Конфигурация
Все приложение настраивается через файл `config.yaml` или переменные окружения. Если используется `config.yaml`, он должен быть в одной директории с бинарником приложения.
Для переменных окружения ставится префикс `SQUAD_`, затем идет `APP_`, `RMQ_` или `DB_` и нужный параметр. Переменные окружения **перезаписывают** значения в файле `config.yaml`.
