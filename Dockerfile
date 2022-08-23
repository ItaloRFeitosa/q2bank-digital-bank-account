FROM golang:1.18-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/main .

FROM alpine:3.16 
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/out/main /app/main

CMD ["/app/main"]