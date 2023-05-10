<div align="center">
  <h1>
    IP Information REST API
  </h1>
  <p>
    Simple REST API server for get IP information
  </p>
  <p>
    <a href="https://github.com/aasaam/ip-info-rest-api/actions/workflows/build.yml" target="_blank"><img src="https://github.com/aasaam/ip-info-rest-api/actions/workflows/build.yml/badge.svg" alt="build" /></a>
    <a href="https://goreportcard.com/report/github.com/aasaam/ip-info-rest-api"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/aasaam/ip-info-rest-api"></a>
    <a href="https://hub.docker.com/r/aasaam/ip-info-rest-api" target="_blank"><img src="https://img.shields.io/docker/image-size/aasaam/ip-info-rest-api?label=docker%20image" alt="docker" /></a>
    <a href="https://github.com/aasaam/ip-info-rest-api/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/aasaam/ip-info-rest-api"></a>
    <a href="https://www.maxmind.com" target="_blank"><img src="https://img.shields.io/badge/IP%20Geolocation-maxmind-00AEEF" alt="maxmind" /></a>
  </p>
</div>

## Guide

For see available options

```bash
docker run --rm ghcr.io/aasaam/ip-info-rest-api:latest -h
```

## Usage

```bash
$ curl -s http://127.0.0.1:4000/info/8.8.8.8 | jq
{
  "IP": "8.8.8.8",
  "IsValid": true,
  "AutonomousSystemNumber": 15169,
  "AutonomousSystemOrganization": "GOOGLE",
  "City": "",
  "CityGeoNameID": 0,
  "Country": "US",
  "LocationLatitude": 37.751,
  "LocationLongitude": -97.822
}
```

## Need more accurate data

- You can use [commercial version](https://www.maxmind.com) for more accurate data.

<div>
  <p align="center">
    <a href="https://aasaam.com" title="aasaam software development group">
      <img alt="aasaam software development group" width="64" src="https://raw.githubusercontent.com/aasaam/information/master/logo/aasaam.svg">
    </a>
    <br />
    aasaam software development group
  </p>
</div>
