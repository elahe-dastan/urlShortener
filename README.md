# urlShortener

Simple project created with golang and postgres

## Introduction

The aim of the project is to map each long URL given to it to a short URL. After the mapping each short URL
will be redirected to the original URL.
The short URL can be custom or not, if not a key generation service (KGS) is used to assign a random 
short URL to the given long URL.

## Installation

The following dependencies have been used in the project

Viper library to manage configuration

```sh
$ go get github.com/spf13/viper
```

Gorm library to interact with database

```sh
$ go get -u github.com/jinzhu/gorm
```

Cobra library to add CLI

```sh
$ go get -u github.com/spf13/cobra/cobra
```

## Run

```sh
$ go build
```
We need to create tables and run KGS only once so there is no need to run the command below evrytime 
you want to run the project

```sh
$ go run main.go setupdb
```
This command runs the APIs

```sh
$ go run main.go server
```

To send HTTP request 

```sh
$ curl -X POST -d '{"LongURL": "http://www.google.com"}' -H 'Content-Type: application/json' 127.0.0.1:8080/urls
$ curl -L 127.0.0.1:8080/redirect/shortURL
```
If you want to use custom short URL or custom expiration date add the followings to the body of 
your post method

```
"ShortURL":"Custom short URL"
"ExpirationTime":"A date"
```
The performance of the project has been tested, here is the result

```sh
Bombarding http://localhost:8080/urls with 10 request(s) using 1 connection(s)
 10 / 10 [==================================================================================] 100.00% 47/s 0s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec       256.97     147.30     371.20
  Latency        4.55ms     4.23ms    17.22ms
  HTTP codes:
    1xx - 0, 2xx - 10, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    83.45KB/s

```
```sh
Bombarding http://localhost:8080/urls with 100 request(s) using 2 connection(s)
 100 / 100 [==================================================] 100.00% 166/s 0s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec       199.42      59.57     311.00
  Latency       10.15ms     7.28ms    33.94ms
  HTTP codes:
    1xx - 0, 2xx - 100, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    74.76KB/s

```
