FROM golang

WORKDIR /go/src/github.com/inloop/nsq-transport
COPY . /go/src/github.com/inloop/nsq-transport

RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /go/src/github.com/inloop/nsq-transport/app /usr/local/bin/nsq-transport

CMD [ "nsq-transport" ]
ENTRYPOINT [ ]