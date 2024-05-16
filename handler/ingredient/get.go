package ingredient

import (
	"encoding/json"
	"github.com/witekblach/fridge/data"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	ingredient := r.Context().Value("ingredient").(*data.Ingredient)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredient)
}
