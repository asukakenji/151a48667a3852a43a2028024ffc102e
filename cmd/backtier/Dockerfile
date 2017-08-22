FROM golang:1.8

WORKDIR /go/src/app

RUN go-wrapper download github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/backtier
RUN go-wrapper install github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/backtier
CMD ["backtier"]
