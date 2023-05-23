# Используем образ Golang для сборки
FROM golang:1.20 AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o chat-service

# Используем образ Alpine для запуска
FROM alpine:latest

# Устанавливаем необходимые пакеты
RUN apk --no-cache add ca-certificates

# Копируем бинарный файл из предыдущего образа
COPY --from=build /app/chat-service /usr/local/bin/chat-service

# Устанавливаем рабочую директорию
WORKDIR /usr/local/bin

# Запускаем приложение при запуске контейнера
CMD ["chat-service"]