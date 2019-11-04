# arrival-time-service
Arrival time service is responsible to define a minimal arrival time of available cars around a given location.

### Development
Requirements
* go version 1.13.1 or higher
* MongoDB version 4.2.1

Make targets:
* `make generate` - generate go-clients by its swagger-specification
* `make run` - run service

### Development environment
1. Export config variables or create a `.env` file consists of configuration variables. Allowed config variables are:
    * `ARRIVAL_TIME_SERVICE_LISTEN_ADDR` - http-listener address in form of `host:port`
    * `ARRIVAL_TIME_SERVICE_CAR_SERVICE_HOST` - cars-service host
    * `ARRIVAL_TIME_SERVICE_CAR_SERVICE_BASE_PATH` - cars-service base path
    * `ARRIVAL_TIME_SERVICE_CAR_SERVICE_LIMIT` - a limit of cars responded by cars-service
    * `ARRIVAL_TIME_SERVICE_PREDICT_SERVICE_HOST` - prediction-service host
    * `ARRIVAL_TIME_SERVICE_PREDICT_SERVICE_BASE_PATH` - prediction-service base path 
    * `ARRIVAL_TIME_SERVICE_DATABASE_DSN` - MongoDB connection string
    * `ARRIVAL_TIME_SERVICE_DATABASE_NAME` - MongoDB database path 
    * `ARRIVAL_TIME_SERVICE_COLLECTION_NAME` - MongoDB collection name for storing routes
    * `ARRIVAL_TIME_SERVICE_SEARCH_RADIUS` - a radius within searching of stored routes allowed
    
    An example `.env` filed available in `.env.example`

2. Execute `docker-compose up` to launch a Mongo server in docker-container
