FROM golang:1.17

WORKDIR /go/src

COPY . .
RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]