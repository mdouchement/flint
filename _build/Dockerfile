FROM golang:1.10.3

RUN apt-get update && apt-get install -y zip

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/astrocorp42/flint

CMD ["make"]
