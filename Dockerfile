FROM golang:1.19-alpine3.17 AS builder

WORKDIR /app

RUN apk add make
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:3.17 AS runner

WORKDIR /app

COPY --from=builder /app/dist ./

CMD ["./main.exe"]