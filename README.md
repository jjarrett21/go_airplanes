# go_airplanes

This project is a simple API that allows the user to send JSON HTTP request to CRUD information about airplanes. The application is written in golang and uses a sqlite databse. It is also dockerized. 

## Getting Started

You will need to clone this repository from either the github provided or docker.

### Prerequisites

This application is written in go, you will need to install go and set your path correctly.

The following links show how to setup go on multiple OS'

* [Mac](https://medium.com/@AkyunaAkish/setting-up-a-golang-development-environment-mac-os-x-d58e5a7ea24f)

* [Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)

* [Linux](https://www.tecmint.com/install-go-in-linux/)

### Installing

Once you have go installed, you will need to migrate to the directory that you have the project saved in. 
Once their to run the main application you will need to do issue the following command.

```
go run main.go
```

Or you can run the docker with

```
docker run --rm -it -p 8080:8080 jjarrett21/go_airplanes
```


Which should start the api on your localhost at port :8080

* [airplanes](http://localhost:8080)

I reccomend using a rest client such was Advanced Rest Client to POST/GET/PUT/DELETE data but,
you can also issue commands via curl

Example PUT request with curl:

```
curl -i -X PUT http://localhost:8080/airplane/1 -d '{"Year": 2001}'
```
This would update the information for plane in index 1 and change it's year to 2001.

* [Advanced Rest Client](https://install.advancedrestclient.com/#/install)

## Running the tests

In order to run unit tests for this application, in the same directory you will need to run the command

``` 
go test airplanes_test.go
```

This will run all the unit tests currently written for the applicaiton.

### Break down into en7 to end tests

Testing for this is simple and for now only tests wether or not each aspect of CRUD is functional.
It simply checks to see if the status is OK or 200.

## Deployment

This is dockerized so it can be deployed easily on any linux based system.

## Built With

* [Go](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://github.com/jinzhu/gorm)
* [Docker](https://www.docker.com/)

## Authors

* **James Jarrett** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments


