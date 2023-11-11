# cf-connecting-ip

This repository contains a middleware for Go HTTP servers that sets the real IP of the client from
the `CF-Connecting-IP` header, which is added by Cloudflare to HTTP requests coming through its services.

## Overview

The middleware intercepts incoming HTTP requests and checks for the `CF-Connecting-IP` header. It then sets
the `RemoteAddr` field of the request to the value found in this header, ensuring that subsequent handlers in the
middleware chain will see the original client IP address, not the IP address of Cloudflare's proxy.

### Installation

To install `cf-connecting-ip` middleware, use `go get`:

```sh
go get -u github.com/shigaichi/cf-connecting-ip
```