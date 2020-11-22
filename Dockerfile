# hadolint ignore=DL3006,DL3007
FROM resin/rpi-raspbian AS builder

# hadolint ignore=DL3008, DL3015
RUN apt-get update && \
    apt-get install -qy build-essential wiringpi git curl ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*


RUN curl -sSLO https://golang.org/dl/go1.15.5.linux-armv6l.tar.gz && \
    mkdir -p /usr/local/go && \
    tar -xvf go1.15.5.linux-armv6l.tar.gz -C /usr/local/go/ --strip-components=1

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN go get -v -t -d ./...
RUN go build -o app .

CMD ["./app"]
