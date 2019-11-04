package main

import (
	"context"
	"log"
	"os"

	"github.com/asardak/arrival-time-service/internal/app"
	"github.com/asardak/arrival-time-service/internal/pkg/api"
	"github.com/asardak/arrival-time-service/internal/pkg/client"
	"github.com/asardak/arrival-time-service/internal/pkg/db"
	cars "github.com/asardak/arrival-time-service/pkg/car-service/client"
	predict "github.com/asardak/arrival-time-service/pkg/predict-service/client"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := getConfig()

	conn := connectDB(c)
	cars := makeCarServiceClient(c)
	prediction := makePredictClient(c)
	arrivalTimeService := app.NewArrivalTimeService(cars, prediction)
	repo := db.NewRepository(conn, c.DatabaseName, c.CollectionName, c.SearchRadius)
	cachedService := app.NewCachedService(arrivalTimeService, repo)
	router := api.NewRouter(cachedService)
	server := makeServer(router)

	err := server.Start(c.ListenAddr)
	if err != nil {
		server.Logger.Fatal(err)
	}
}

type config struct {
	ListenAddr             string `split_words:"true"`
	CarServiceHost         string `split_words:"true"`
	CarServiceBasePath     string `split_words:"true"`
	CarServiceLimit        int64  `split_words:"true"`
	PredictServiceHost     string `split_words:"true"`
	PredictServiceBasePath string `split_words:"true"`
	DatabaseDSN            string `split_words:"true"`
	DatabaseName           string `split_words:"true"`
	CollectionName         string `split_words:"true"`
	SearchRadius           int    `split_words:"true"`
}

const envPrefix = "ARRIVAL_TIME_SERVICE"

func getConfig() config {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	var c config
	err = envconfig.Process(envPrefix, &c)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}

func connectDB(c config) *mongo.Client {
	ctx := context.Background()

	conn, err := mongo.Connect(
		ctx, options.Client().ApplyURI(c.DatabaseDSN),
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func makeCarServiceClient(c config) *client.CarService {
	carService := cars.NewHTTPClientWithConfig(
		nil,
		&cars.TransportConfig{
			Host:     c.CarServiceHost,
			Schemes:  []string{"https"},
			BasePath: c.CarServiceBasePath,
		},
	)

	return client.NewCarService(carService.Operations, c.CarServiceLimit)
}

func makePredictClient(c config) *client.PredictService {
	predictService := predict.NewHTTPClientWithConfig(
		nil,
		&predict.TransportConfig{
			Host:     c.PredictServiceHost,
			Schemes:  []string{"https"},
			BasePath: c.PredictServiceBasePath,
		},
	)

	return client.NewPredictService(predictService.Operations)
}

type Mounter interface {
	Mount(e *echo.Echo)
}

func makeServer(router Mounter) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Mount(e)

	return e
}
