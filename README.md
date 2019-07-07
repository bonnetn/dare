Dare
====
> Just a CRUD API + its React client.

Description
-----------

#### API
The API is built on top of [GRPC](https://grpc.io/) and stores objects in a [MongoDB](https://www.mongodb.com/) database.
(It is basically a CRUD API)

This API was built using the [Clean architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

#### Client
The client is built with [ReactJS](https://reactjs.org/) and [Redux](https://redux.js.org/) libraries.

It communicates with the API using [GRPC-web](https://github.com/grpc/grpc-web).

We are using [Envoy](https://www.envoyproxy.io/) proxy to convert the GRPC-web protocol to GRPC.




How to run
----------

### Dependencies
In order to run this project you need `make`, `go`, `yarn` and `docker-compose`.
 
### Launching

In order to launch the dependencies (SQL server, Redis...) locally, run:
> make docker-deps

Launch the API server in your IDE, or with Make:
> make run-api

In another terminal, launch the React client.
> make run-react
