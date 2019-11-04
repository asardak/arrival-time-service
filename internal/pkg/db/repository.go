package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/asardak/arrival-time-service/internal/app"

	"go.mongodb.org/mongo-driver/mongo"
)

type Point struct {
	Location struct {
		Type        string     `json:"type"`
		Coordinates [2]float64 `json:"coordinates"`
	} `json:"location"`
	Time int64 `json:"time"`
}

type Repository struct {
	client         *mongo.Client
	db             string
	collectionName string
	searchRadius   int
}

func NewRepository(client *mongo.Client, db string, collection string, searchRadius int) *Repository {
	return &Repository{
		client:         client,
		db:             db,
		collectionName: collection,
		searchRadius:   searchRadius,
	}
}

func (r *Repository) FindRoute(ctx context.Context, point app.Point) (app.Route, error) {
	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": bson.A{point.Lng, point.Lat},
				},
				"$maxDistance": r.searchRadius,
			},
		},
	}

	var res Point
	err := r.collection().FindOne(ctx, filter).Decode(&res)
	if err == nil {
		return app.Route{
			Point: app.Point{
				Lat: res.Location.Coordinates[1],
				Lng: res.Location.Coordinates[0],
			},
			Time: res.Time,
		}, nil
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return app.Route{}, app.ErrRouteNotFound
	}

	return app.Route{}, err
}

func (r *Repository) SaveRoute(ctx context.Context, route app.Route) error {
	_, err := r.collection().InsertOne(ctx, bson.M{
		"time": route.Time,
		"location": bson.M{
			"type":        "Point",
			"coordinates": bson.A{route.Lng, route.Lat},
		},
		"expireAt": time.Now().Add(time.Minute),
	})

	return err
}

func (r *Repository) collection() *mongo.Collection {
	return r.client.Database(r.db).Collection(r.collectionName)
}
