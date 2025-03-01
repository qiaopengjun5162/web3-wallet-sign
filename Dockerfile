FROM golang:1.22.8-alpine3.18 as builder

#RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash
RUN apk add --no-cache \
    make=4.3-r1 \
    ca-certificates=20230506-r0 \
    gcc=12.2.1_git20220924-r10 \
    musl-dev=1.2.4-r2 \
    linux-headers=6.3-r0 \
    git=2.40.1-r0 \
    jq=1.6-r3 \
    bash=5.2.15-r5

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN go mod download

ARG CONFIG=config.yml

# build web3-wallet-sign with the shared go.mod & go.sum files
COPY . /app/web3-wallet-sign

WORKDIR /app/web3-wallet-sign

RUN make

FROM alpine:3.18

COPY --from=builder /app/web3-wallet-sign/web3-wallet-sign /usr/local/bin
COPY --from=builder /app/web3-wallet-sign/${CONFIG} /etc/web3-wallet-sign/

RUN chmod +x /usr/local/bin/web3-wallet-sign

WORKDIR /app

ENTRYPOINT ["web3-wallet-sign"]
CMD ["-c", "/etc/web3-wallet-sign/config.yml"]
