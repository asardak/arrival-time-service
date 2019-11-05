// +build db

package test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/asardak/arrival-time-service/internal/pkg/db"

	"github.com/asardak/arrival-time-service/internal/app"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Client

const (
	dbName         = "arrival-time-service"
	collectionName = "routes"
)

func init() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	dsn := os.Getenv("ARRIVAL_TIME_SERVICE_DATABASE_DSN")

	ctx := context.Background()
	conn, err = mongo.Connect(
		ctx, options.Client().ApplyURI(dsn),
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRepository_FindRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		mins := int64(10)
		lat := 55.112233
		lng := 37.112233
		expect := app.Route{
			Point: app.Point{Lat: lat, Lng: lng},
			Time:  mins,
		}

		_, err := conn.Database(dbName).Collection(collectionName).
			InsertOne(ctx, bson.M{
				"time": mins,
				"location": bson.M{
					"type":        "Point",
					"coordinates": bson.A{lng, lat},
				},
				"expireAt": time.Now().Add(time.Hour),
			})

		assert.Nil(t, err)

		repo := db.NewRepository(conn, dbName, collectionName, 500, time.Minute)
		r, err := repo.FindRoute(ctx, app.Point{Lat: 55.111111, Lng: 37.111111})

		assert.Nil(t, err)
		assert.Equal(t, expect, r)
	})

	t.Run("Expired", func(t *testing.T) {
		ctx := context.Background()
		mins := int64(10)
		lat := 65.112233
		lng := 47.112233

		_, err := conn.Database(dbName).Collection(collectionName).
			InsertOne(ctx, bson.M{
				"time": mins,
				"location": bson.M{
					"type":        "Point",
					"coordinates": bson.A{lng, lat},
				},
				"expireAt": time.Now().Add(-time.Hour),
			})

		assert.Nil(t, err)

		time.Sleep(time.Second)

		repo := db.NewRepository(conn, dbName, collectionName, 500, time.Minute)
		_, err = repo.FindRoute(ctx, app.Point{Lat: 65.111111, Lng: 47.111111})

		assert.Equal(t, app.ErrRouteNotFound, err)
	})

	t.Run("Not found", func(t *testing.T) {
		ctx := context.Background()
		mins := int64(10)
		lat := 55.112233
		lng := 37.112233

		_, err := conn.Database(dbName).Collection(collectionName).
			InsertOne(ctx, bson.M{
				"time": mins,
				"location": bson.M{
					"type":        "Point",
					"coordinates": bson.A{lng, lat},
				},
				"expireAt": time.Now().Add(time.Hour),
			})

		assert.Nil(t, err)

		repo := db.NewRepository(conn, dbName, collectionName, 500, time.Minute)
		_, err = repo.FindRoute(ctx, app.Point{Lat: 45.111111, Lng: 27.111111})

		assert.Equal(t, app.ErrRouteNotFound, err)
	})

}

func TestRepository_SaveRoute(t *testing.T) {
	ctx := context.Background()
	repo := db.NewRepository(conn, dbName, collectionName, 500, time.Minute)
	route := app.Route{Point: app.Point{Lat: 75.112233, Lng: 57.112233}, Time: 112233}
	err := repo.SaveRoute(ctx, route)

	assert.Nil(t, err)

	var res db.Route
	err = conn.Database(dbName).
		Collection(collectionName).
		FindOne(ctx, bson.M{"time": 112233}).
		Decode(&res)

	assert.Nil(t, err)
	assert.Equal(t, route.Lat, res.Location.Coordinates[1])
	assert.Equal(t, route.Lng, res.Location.Coordinates[0])
	assert.True(t, time.Now().Before(res.ExpireAt))
	assert.Equal(t, route.Time, res.Time)
}
