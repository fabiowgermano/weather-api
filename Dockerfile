FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/weathersystem

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cmd/weathersystem/.env .
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
