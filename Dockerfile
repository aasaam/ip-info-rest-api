FROM golang:1.19-buster AS builder

ADD . /src

RUN cd /src \
  && go mod tidy \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ip-info-rest-api \
  && ls -lah /src/ip-info-rest-api

FROM alpine
COPY --from=builder /src/ip-info-rest-api /usr/bin/ip-info-rest-api

ADD /tmp/GeoLite2-ASN.mmdb /tmp/GeoLite2-ASN.mmdb
ADD /tmp/GeoLite2-City.mmdb /tmp/GeoLite2-City.mmdb

ENTRYPOINT ["/usr/bin/ip-info-rest-api"]
