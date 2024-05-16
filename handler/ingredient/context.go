package ingredient

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/witekblach/fridge/data"
	"net/http"
)

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
func IngredientsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ingredientName := chi.URLParam(r, "ingredientName")

		getIngredient, err := data.GetIngredient(ingredientName)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
			//http.Error(w, http.StatusText(404), 404)
			//return
		}

		ctx := context.WithValue(r.Context(), "ingredient", getIngredient)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
