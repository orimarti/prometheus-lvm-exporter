FROM golang:1.8

RUN apt-get update || exit 0
RUN apt-get install -y lvm2
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
