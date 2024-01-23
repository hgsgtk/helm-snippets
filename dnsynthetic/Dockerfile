FROM golang:1.21-alpine as builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download
RUN go build -o dnsynthetic .

FROM alpine as executor

COPY --from=builder /app/dnsynthetic .

CMD ["./dnsynthetic"]
