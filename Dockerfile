FROM balenalib/raspberrypi3-golang:1.15

# hadolint ignore=DL3008, DL3015
RUN apt-get update && \
    apt-get install -qy wiringpi && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN go get -v -t -d ./...
RUN go build -o app .

CMD ["./app"]
