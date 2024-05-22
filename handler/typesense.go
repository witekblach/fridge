package handler

import (
	"context"
	"encoding/json"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/witekblach/fridge/typesense"
	"log/slog"
	"net/http"
	"os"
)

func TypesenseHandler(w http.ResponseWriter, r *http.Request) {
	err := typesense.NewTypeSenseClient()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	searchParameters := &api.SearchCollectionParams{
		Q:       r.URL.Query().Get("q"),
		QueryBy: typesense.CollectionIngredientName,
	}

	search, err := typesense.TypeSense.Collection(typesense.CollectionIngredients).Documents().Search(context.Background(), searchParameters)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(search.Hits)
}
