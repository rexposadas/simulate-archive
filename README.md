simulate
========

### Why simulate

Unit and integration test run for a very limited amount of time.  They are not very good at detecting issues which may arise from long usage of your API.  These issues are performance degradation and data corruption.

Simulate can be as intrusive as you want (that's a good thing).  It can act as an outside application which treats your API as a blackbox or can be imported in your Go application and give it access to features you want to test. Note that having simulate run separately from your application allows your application to be written in any language since the interaction will strictly be via REST.


### Things you can do with this simulator

1. Hit endpoints much like any API test applications
1. Validate that actions taken on your API is represented correctly in your database.
1. Detect issues in your application which can surface only after long periods of use, such as data corruption and performance degradation.
2. You can simulate release candidate before moving it further along the deployment chain.

### Getting started

The quickest way to get started is to run the sample application.

Get the package.

	go get github.com/rexposadas/simulate

CD into the newly created directory and run the sample application:

	go run app/sample/sample.go

You will see the sample application make a GET request to `http://google.com` and print out the response time.

### Getting started

Create a simple application. In your main.go

	package main

	import "github.com/rexposadas/simulate/server"

	// GetGoogle make a GET request to http://google.com
	// error handling has been omited for simplicity. Please add error
	// handling in your code.
	func GetGoogle() error {
		_, err := simhttp.Get("http://google.com")
		return err
	}

	func main() {
		server.Run(7000)  // starts the simulator on port 7000
		// Create job and send to scheduler
		g := server.NewJob()
		g.Run = GetGoogle
		server.Jobs <- g

		// Simulate will run the job once a second which is the default behavior.

		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
		<-sigc
	}

That's it. This will make a GET request to your API and print out response times.




### Roadmap

* REST API and etcd to accept job on-the-fly
* LogStash integration (or the entire ELK stack)
* More examples