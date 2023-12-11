FROM golang:1.21 as builder

WORKDIR /go/src/github.com/clibing/go-common

COPY . .

ENV CGO_CFLAGS "-g -O2 -Wno-return-local-addr"

ENV GOPROXY "https://goproxy.cn,direct"

RUN make single

FROM clibing/ubuntu

WORKDIR /

COPY --from=builder /go/src/clibing/go-common/dist/single /bin/go-common

COPY --from=builder /go/src/clibing/go-common/configs/config.yaml /etc/config.yaml

ENV GIN_MODE=release

CMD ["/bin/go-common", "server", "--config", "/etc/config.yaml"]