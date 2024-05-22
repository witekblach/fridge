package typesense

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"os"
)

var TypeSense *typesense.Client

func NewTypeSenseClient() error {
	var (
		tsHost = os.Getenv("TYPESENSE_HOST")
		apiKey = os.Getenv("TYPESENSE_API_KEY")
		uri    = fmt.Sprintf("http://%s:8108", tsHost)
	)

	client := typesense.NewClient(
		typesense.WithServer(uri),
		typesense.WithAPIKey(apiKey),
	)

	TypeSense = client

	return nil
}
