FROM golang:1.23-alpine AS builder
ENV GOTOOLCHAIN=auto
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/

FROM alpine:3.19

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/main .
COPY --from=builder /app/web ./web

EXPOSE 8080

ENTRYPOINT ["./main"]