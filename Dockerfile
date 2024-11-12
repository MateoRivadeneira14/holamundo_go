# Fase de construcción
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o /go/bin/app -v ./...

# Fase final
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates

# Copiar el binario desde la fase de construcción al directorio /app en el contenedor
COPY --from=builder /go/bin/app /app/

# Asegúrate de que el binario tiene permisos de ejecución
RUN chmod +x /app/app

# Verifica si el binario tiene permisos correctos
RUN ls -l /app

# Establecer el comando por defecto
ENTRYPOINT ["/app/app"]
LABEL Name=holamundogo Version=0.0.1
EXPOSE 3002