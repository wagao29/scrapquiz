FROM golang:1.23.0-alpine3.20

WORKDIR /data
COPY ./ ./

RUN go mod download

RUN go build -o app cmd/main.go

EXPOSE 8080

CMD [ "./app" ]
