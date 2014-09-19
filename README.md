simulate
========

### Why simulate

Unit and integration test run for a very limited amount of time.  They are not very good at detecting issues which may arise from long usage of your API.  These issues are performance degradation and data corruption.


### Things you can do with this simulator

1. Hit endpoints much like any API test applications
1. Chain endpoints. A good example would be the registration process.

### Getting started

The quickest way to get started is to run the sample application.

Get the package.

	go get github.com/rexposadas/simulate 

CD into the newly created directory and run the sample application:

	go run app/sample/sample.go

You will see the sample application make a GET request to `http://google.com` and print out the response time. 

### To get started with your own application

Create a simple application. In your main.go 

	package main

	import "github.com/rexposadas/simulate"

	func main(){

		simulate.Run()  // starts the simulator
		simulate.Add("http://google.com")  // add you own endpoint here
	}

That's it. This will make a GET request on your API endpoint and print out the response. 


### Roadmap

* REST API and etcd to accept job on-the-fly
* LogStash integration (or the entire ELK stack)
* More examples