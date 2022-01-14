ARG GO_VERSION=1.17
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY . ./

RUN go build -o goBinary .

## Distribution 
FROM alpine:latest

RUN apk add --no-cache --upgrade ca-certificates tzdata 

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/goBinary ./

CMD [ "/app/goBinary" ]