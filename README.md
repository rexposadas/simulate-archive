simulate
========

### Why simulate

Unit and integration test, often run for a very limited amount of time.  They are not very good detecting issues that may arise from long usage of your API.  These issues are
performance degradation and data corruption.


### Things you can do with this simulator

1. Hit endpoints much like any API test applications
1. Chain endpoints. A good example would be the registration process.
1.

### Getting started

Get the package.

	go get github.com/rexposadas/simulate 

Create a simple application. In your main.go 

	package main

	import "github.com/rexposadas/simulate"

	func main(){

		simulate.Run()  // starts the simulator
		simulate.Add("http://google.com")  // add you own endpoint here
	}

That's it. What this will do is run make a GET request to `google.com` and prints out the response.  You can checkout the sample apps under the `apps` folder.








