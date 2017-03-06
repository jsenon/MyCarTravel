# MyCarTravel


Weberserver for Distance and time calculation based on Google MAP API

### Prerequisite

You need to have:

* Go 1.8
* Go Environment properly set
* Google API Key 

### Installation 

Build WebApp
```sh
$ go build src/MycarTravel/mycartravel.go
```

Run WebApp
```
$ export GOOGLE_APIKEY=YOURAPIKEY
$ ./mycartravel
```

Command line function is also available in program

### Access

Webserver socket created on 127.0.0.1:9000
URL http://127.0.0.1:9000/hello

### Todo

Result in HTML page
