FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN mkdir -p /tmp/in

RUN rm go.mod

RUN go mod init

RUN go mod tidy

EXPOSE 2020

CMD ["go", "run", "."]