<p align="center">
  <img src="https://github.com/marceloboeira/crowd/blob/master/docs/logo.png?raw=true" width="150">
  <h3 align="center">crowd</h3>
  <p align="center">High Available Reverse Proxy for Asynchronous Message Consumption<p>
  <p align="center">
    <a href="https://travis-ci.org/marceloboeira/crowd"><img src="https://img.shields.io/travis/marceloboeira/crowd.svg?maxAge=360"></a>
    <a href="http://waffle.io/marceloboeira/bojack"><img src="https://img.shields.io/waffle/label/marceloboeira/bojack/ready.svg?maxAge=360"></a>
    <a href="http://github.com/marceloboeira/crowd/releases"><img src="https://img.shields.io/github/release/marceloboeira/crowd.svg?maxAge=360"></a>
  </p>
</p>

## Motivation

The project started because of a common usecase for an endpoint(HTTP) to ingest data, where the user on the client side expects to receive a delivery confirmation that the message has been received, yet the response itself doens't matter.

e.g.:

Request:
```
POST /leads HTTP/1.1
Content-Type: application/json
Accept: application/json
Content-Length: 327
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15

{"message":"this is a test message"}
```

Response:
```
HTTP/1.1 200
Content-Type: application/json;charset=UTF-8
Vary: Accept-Encoding
Connection: keep-alive
Date: Sun, 09 Dec 2018 11:07:54 GMT
Content-Encoding: gzip
Transfer-Encoding: Identity

{"status": "OK"}
```

The whole point is to have a **high available** endpoint to avoid package loss, disregarding the processing.

which is a quite common pattern for data-ingestion in general: statistics, events, tracking, order-placement, queue items, ...

### Use cases
> Some usecases

You can use crowd to handle back-pressure on existing endpoints and scale your current REST APIs.

<img src="https://github.com/marceloboeira/crowd/blob/master/docs/usecase-1.png?raw=true" width="500">

Or even as an entrypoint to your stream/queues with multiple consumers of the data:

<img src="https://github.com/marceloboeira/crowd/blob/master/docs/usecase-2.png?raw=true" width="500">
