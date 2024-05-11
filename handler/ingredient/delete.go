package ingredient

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("%+v", r.Context().Value("ingredient")))

	ingredientToRemove := r.Context().Value("ingredient")

	//data.RemoveIngredient(ingredientToRemove)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredientToRemove)
}
