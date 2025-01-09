FROM golang:alpine3.21
COPY ./main.go /go/src
RUN go build -o /go/bin/app /go/src/main.go
CMD ["/go/bin/app" ]
