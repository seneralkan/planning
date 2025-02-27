FROM golang:1.23.6-alpine AS builder

ARG DOCKER_GIT_CREDENTIALS

RUN apk update && apk add --no-cache git

WORKDIR /microservice

COPY . .

RUN go mod download

RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/*

FROM scratch as runner

COPY --from=builder /microservice/app .

EXPOSE 8080

CMD ["/app"]