package dal

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"apipost-dcm/internal/pkg/conf"
)

var (
	m *mongo.Client
)

func MustInitMongo() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m, err = mongo.Connect(ctx, options.Client().ApplyURI(conf.Conf.MongoDB.DSN).
		SetReadPreference(readpref.Secondary()).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority(), writeconcern.J(true), writeconcern.WTimeout(1000))).
		SetMaxPoolSize(conf.Conf.MongoDB.PoolSize))

	if err != nil {
		panic(fmt.Errorf("mongo err:%w", err))
	}

	err = m.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongodb initialized")
}

func GetMongo() *mongo.Client {
	return m
}

func MongoDB() string {
	return conf.Conf.MongoDB.Database
}
