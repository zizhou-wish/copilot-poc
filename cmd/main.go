package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"copilot-poc/internal/graph"
)

const defaultPort = "8080"

// Config is the top level service configd
type Config struct {
	MongoDB MongoDBConfig `mapstructure:"mongodb"`
}

// Create a struc for my mongoDB config
type MongoDBConfig struct {
	URI string `mapstructure:"uri"`
	DB  string `mapstructure:"db"`
}

func main() {
	// Load configuration from JSON file using Viper
	config, err := initConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)

	}

	// use config to connect to MongoDB
	client, ctx, cancel, err := initMongoDB(config)
	defer cancel()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect to MongoDB: %v", err)
		}
	}()

	// Get MongoDB todo collection
	todoCollection := client.Database(config.MongoDB.DB).Collection("todos")
	// Get MongoDB user collection
	userCollection := client.Database(config.MongoDB.DB).Collection("users")

	// Set up GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		TodoCollection: todoCollection,
		UserCollection: userCollection,
	}}))

	// Set up HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Start HTTP server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initMongoDB(config Config) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDB.URI))
	if err != nil {
		return nil, ctx, cancel, err
	}

	return client, ctx, cancel, err
}

// initialize the config and return it
func initConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return Config{}, fmt.Errorf("fatal error config file: %w", err)
	}

	// use unmarshal to get the config into a struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	// Check if the required config is set
	if config.MongoDB.URI == "" {
		return config, errors.New("MongoDB URI not set")
	}
	if config.MongoDB.DB == "" {
		return config, errors.New("MongoDB DB not set")
	}

	return config, err
}
