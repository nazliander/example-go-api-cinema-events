FROM golang:alpine3.14

COPY . /go/src/app
WORKDIR /go/src/app

# Build
RUN go build -o cinema-events

EXPOSE 6661

CMD ["./cinema-events"]
