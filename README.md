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

```
$ go get github.com/spf13/viper
```

Gorm library to interact with database

```
$ go get -u github.com/jinzhu/gorm
```

Cobra library to add CLI

```
$ go get -u github.com/spf13/cobra/cobra
```

## Run

```
$ go build
```
We need to create tables and run KGS only once so there is no need to run the command below evrytime 
you want to run the project

```
$ go run main.go setupdb
```
This command runs the APIs

```
$ go run main.go server
```
