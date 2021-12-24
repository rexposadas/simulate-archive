simulate
========

This is a library in its *alpha* stage. 


### The motivation

It was born out of the need to limit expenses by not hiring a QA team, but still wanting to build a robust and well tested API. I also created this library to answer questions we, my current firm, could not answer by simply hiring (if we wanted to hire). 

*How do you test a thousand users using your API? How about a thousand mobile devices?*

Obviously you will need to programatically test this since it would be nonsense to hire a thousand testers. 

The next question was, how do we know how well our application behaves over a long period of time *before* releasing it to the public?  Say I want to release a big patch, which may have breaking changes.  I want to have a certain level of confidence before releasing it to production.  Yes, integretion and unit tests helps a lot.  But to gain more confidence, I woudl like to simulate actions against the code over a 24 hour period and see if something goes wrong.  Again, this has to be done programatically.  Unit and integration test run for a very limited amount of time.  They are not very good at detecting issues which may arise from long usage of your API.  Some of these issues are performance degradation and data corruption.

Simulate can be as intrusive as you want (don't worry, that's a good thing.  I can explain).  It can act as an outside application which treats your API as a blackbox or can be imported in your Go application and give it access to features you want to test. Note that having simulate run separately from your application allows your application to be written in any language since the interaction will strictly be via REST.


### Things you can do with this simulator

1. Hit endpoints much like any API test applications
1. Create and run "actors" which validate that actions taken on your API is represented correctly in your database.
1. Detect issues in your application which can surface only after long periods of use, such as data corruption and performance degradation.  Simulate can send data to [influxDB](http://influxdb.com/).
1. Test release candidates before moving it further along the deployment chain.

### Getting started

The quickest way to get started is to run the sample application.

Installation:

	go get github.com/rexposadas/simulate

CD into the newly created directory and run the sample application:

	go run apps/example1/main.go

You will see the sample application make a GET request to `http://google.com` and print out the response time.

To get started with your own simulation, refer to the code in `apps/sample/example1/main.go`. That example is meant as a walk-through of a simple simulation.

### Output

There are two ways simulate outputs results. The default is via stdout. The other is sending it's results to InfluxDB. 

### Use InfluxDB and Graphana

### Roadmap

* Metrics integration 
* More examples
* Determine list of features
