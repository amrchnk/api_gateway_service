#импортируем готовый образ, на котором будет работать приложение
FROM golang:1.16-buster

WORKDIR /app

COPY ./ ./

# импортируем зависимости и создаем бинарный файл
RUN go mod download
RUN go build -o design_app ./cmd/main.go

EXPOSE 8000
# точка входа в приложение
CMD ["./design_app"]