package ingredient

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/witekblach/fridge/data"
	"net/http"
)

func IngredientsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ingredientName := chi.URLParam(r, "ingredientName")

		getIngredient, err := data.GetIngredient(ingredientName)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "ingredient", getIngredient)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
