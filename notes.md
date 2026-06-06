go mod init myapp          # создаём модуль
go get github.com/gin-gonic/gin@latest
go mod tidy                # чистим неиспользуемые зависимости
go list -m all             # смотрим все зависимости
go run .                   # запуск всего пакета