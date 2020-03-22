[![Drone (cloud)](https://img.shields.io/drone/build/elahe-dastan/urlShortener.svg?style=flat-square)](https://cloud.drone.io/elahe-dastan/urlShortener)

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
For a small API documentation check this URL:
https://app.swaggerhub.com/apis/elahe.dstn/urlshortener/1.0.0

Different ideas to build a URL shortener
I have thought about four possible ideas to build a URL shortener that will be discussed below

1. Generating a random short URL while inserting :
This approach is the easiest one in which every time the user posts a long URL and expects a short URL back
we can generate a random short URL at that moment and insert to the database

Advantages:
Super easy to write the code 

Disadvantages: 
1- Time is taken to produce a random short URL 
2- The short URL generated may be duplicate that returns error while inserting to the database so should be 
regenerated and infects the performance

2. Encoding actual URL :
There are hash functions like MD% or SHA256 that we can use to generate the short URL of a long one ,this 
functions produce a long output so we can consider just a part of it as the short URL

Disadvantages:
I think the most noticeable disadvantage of this approach is collision, we know that two long URL may have the
same short URL using this way

3. Base conversion:
This approach has the most light weight database because we don't save the short URL at all when we insert a 
long URL we get it's ID back and based on the number of characters we want to use in our short URL we can
convert this ID to a short URL and when searching for a short URL we should first convert it to ID 

Disadvantage:
The base Conversion takes time for both inserting and redirecting operation

PM: I have a small implementation for this idea in c#, here is the link:
https://github.com/elahe-dastan/urlshortener_alibaba

4. Generating keys offline:
In this approach we generate all the possible short URLs with a specified length and keep them in a table
with a boolean column that shows if we have used a short URL or not every time we want to insert a long URL 
we can pick up on the unused short URLs from this table 

Advantages: No time is taken to generate a short URL

Disadvantage: Operation with database to pick up a short URL takes time