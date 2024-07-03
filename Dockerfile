FROM golang:alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /src/http_server ./app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /src/http_server .
ENV DIRECTORY_PATH=/data
EXPOSE 3000
ENTRYPOINT ["/app/http_server"]