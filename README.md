# 451-blackhole
Web server that doing only ONE thing: reply with ONE static html page with 451 status code for ANY request.

## Build

1. git clone this
2. go build (or just `make`)

## Best practicies?

With nginx you can easily get HTTPS support. Look at the config example below:

``` nginx
server {
    listen       10.50.140.73:80;                       # for http-redirection
    listen       10.50.140.73:443 ssl;                  # for dns-spoofed https requests

    ssl_certificate      ssl/cert.pem;
    ssl_certificate_key  ssl/cert.key;
    ssl_session_timeout  5m;
    ssl_protocols  SSLv2 SSLv3 TLSv1;
    ssl_ciphers  ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
    ssl_prefer_server_ciphers   on;

    location / {
            proxy_pass http://127.0.0.1:8080;
    }
}
```

## What else I need or may to do?

There are two env-variables you can pass to blackhole binary:

- `LISTEN` - golang listen string, e.g. ":8080" or "127.0.0.1:80"
- `FILENAME` - PATH to file with HTML for responses.

## How to check it works ok?

``` shell
echo lol > index.html
go build
FILENAME=index.html LISTEN="127.0.0.1:8080" ./451-blackhole
```

In another shell run:

``` shell
curl -vL http://127.0.0.1:8080
```

Output should be something like:

``` shell
* Rebuilt URL to: http://127.0.0.1:8080/
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.49.1
> Accept: */*
> 
< HTTP/1.1 451 Unavailable For Legal Reasons
< Date: Wed, 22 Mar 2017 08:37:33 GMT
< Content-Length: 4
< Content-Type: text/plain; charset=utf-8
< 
lol
* Connection #0 to host 127.0.0.1 left intact
```
