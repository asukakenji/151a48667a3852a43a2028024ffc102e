FROM golang:1.8

WORKDIR /go/src/app
COPY frontier.json .

RUN go-wrapper download github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/frontier
RUN go-wrapper install github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/frontier
CMD ["frontier"]
