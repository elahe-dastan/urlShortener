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
