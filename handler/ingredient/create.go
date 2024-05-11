package ingredient

import (
	"encoding/json"
	"github.com/witekblach/fridge/data"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var req data.CreateIngredientRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ingredientToAdd := data.Ingredient{Name: req.Name}

	data.AddIngredient(ingredientToAdd)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredientToAdd)
}
