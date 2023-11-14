# build
FROM --platform=linux/amd64 golang:1.20.2-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/services ./src/services
RUN CGO_ENABLED=0 go build -o bin/migration ./src/migration

FROM --platform=linux/amd64 alpine:3.16.7
COPY --from=builder /app/bin/services /
COPY --from=builder /app/bin/migration /
COPY --from=builder /app/.env.example /.env
COPY --from=builder /app/src /app
EXPOSE 10000
CMD ["/services"]