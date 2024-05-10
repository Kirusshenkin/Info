# Используйте официальный образ Golang
FROM golang:1.16-alpine as builder

# Установите рабочий каталог
WORKDIR /app

# Скопируйте модульные файлы
COPY go.mod go.sum ./

# Загрузите зависимости
RUN go mod download

# Скопируйте исходный код проекта
COPY . .

# Соберите приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o crypto-prices cmd/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /root/

# Скопируйте собранный бинарник в финальный образ
COPY --from=builder /app/crypto-prices .

# Откройте порт, на котором работает приложение
EXPOSE 8080

# Запустите приложение
CMD ["./crypto-prices"]
