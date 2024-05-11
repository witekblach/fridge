package ingredient

import (
	"encoding/json"
	"github.com/witekblach/fridge/data"
	"log/slog"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	ingredientToRemove := r.Context().Value("ingredient").(*data.Ingredient)

	err := data.RemoveIngredient(*ingredientToRemove)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredientToRemove)
}
