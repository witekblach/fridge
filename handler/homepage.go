package handler

import (
	"encoding/json"
	"github.com/witekblach/fridge/data"
	"log/slog"
	"net/http"
	"os"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	ingredients, err := data.ShowAllIngredients()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}
