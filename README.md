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

	go run apps/example1/main.go

You will see the sample application make a GET request to `http://google.com` and print out the response time.

To get started with your own simulation, refer to the code in `apps/sample/example1`. `main.go` from that example is meant as a walk-through of a simple simulation.


### Roadmap

* REST API and etcd to accept job on-the-fly
* LogStash integration (or the entire ELK stack)
* More examples