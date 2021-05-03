FROM golang:1.16
COPY . /go/src
WORKDIR /go/src/cmd
RUN go mod download && go build main.go
CMD ["./main"]