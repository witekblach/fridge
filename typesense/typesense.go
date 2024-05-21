package typesense

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"os"
)

var TypeSense *typesense.Client

type MongoClient struct {
	client *mongo.Client
}

func NewTypeSenseClient() error {
	var (
		tsHost = os.Getenv("TYPESENSE_HOST")
		apiKey = os.Getenv("TYPESENSE_API_KEY")
		uri    = fmt.Sprintf("http://%s:8108", tsHost)
	)
	
	slog.Info("error", tsHost)
	slog.Info("error", apiKey)
	slog.Info("error", uri)

	client := typesense.NewClient(
		typesense.WithServer(uri),
		typesense.WithAPIKey(apiKey),
	)

	TypeSense = client

	return nil
}
