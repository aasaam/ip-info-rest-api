#!/bin/bash

set -e

# mmdb
docker rm -f maxmind-lite-docker-test
docker run --name maxmind-lite-docker-test -d ghcr.io/aasaam/maxmind-lite-docker tail -f /dev/null
docker cp maxmind-lite-docker-test:/GeoLite2-City.mmdb /tmp/GeoLite2-City.mmdb
docker cp maxmind-lite-docker-test:/GeoLite2-ASN.mmdb /tmp/GeoLite2-ASN.mmdb
docker rm -f maxmind-lite-docker-test
