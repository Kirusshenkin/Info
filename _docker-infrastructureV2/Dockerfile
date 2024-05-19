# Используйте официальный образ Golang
FROM golang:1.21-alpine as builder

# Установите рабочий каталог
WORKDIR /app

# Скопируйте модульные файлы в рабочий каталог
COPY go.mod go.sum ./

# Загрузите зависимости
RUN go mod download

# Скопируйте исходный код проекта в рабочий каталог
COPY . .

# Соберите приложение, указывая путь к main.go относительно WORKDIR
RUN CGO_ENABLED=0 GOOS=linux go build -o cryptoApi cmd/main.go

# Используйте официальный образ Alpine для финального образа
FROM alpine:latest
WORKDIR /root/

# Скопируйте собранный бинарник в финальный образ
COPY --from=builder /app/cryptoApi .

# Откройте порт, на котором работает приложение
EXPOSE 8080

# Запустите приложение
CMD ["./cryptoApi"]
