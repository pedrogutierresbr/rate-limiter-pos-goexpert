FROM golang:1.22.3 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/ -o golang_rate_limit

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/.env /app/cmd/golang_rate_limit ./
ENTRYPOINT ["./golang_rate_limit"]