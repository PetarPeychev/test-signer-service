FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY api/ api/
RUN CGO_ENABLED=0 GOOS=linux go build -o run

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/run .
CMD ["./run"]
