# httpmirror: A single binary HTTP server that mirrors all request data (headers and body) in the response


## Install

```
$ go get github.com/multiprocessio/httpmirror@latest
```

## Usage

```
$ httpmirror 8081
2022/03/14 15:07:29 Listening on :8081
```

Then in another terminal:

```
$ curl -X POST -d '{"foo": "bar"}' -H "Content-Type: application/json" localhost:8081
POST / HTTP/1.1
Host: localhost:8081
User-Agent: curl/7.77.0
Content-Length: 14
Accept: */*
Content-Type: application/json

{"foo": "bar"}
```

## Why?

To make integration testing easier.
