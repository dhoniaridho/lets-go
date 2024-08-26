FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV PORT=3000

RUN apk add make
COPY . .
RUN go mod download
RUN make build-web

FROM golang:1.22-alpine3.20 AS runner


WORKDIR /app
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY --from=builder /app/dist ./
COPY db db
COPY docker/entrypoint/run.sh /app/run.sh
RUN chmod +x ./run.sh

ENTRYPOINT [ "./run.sh" ]
EXPOSE 3000
