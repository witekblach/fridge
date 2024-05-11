package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/witekblach/fridge/data"
	"github.com/witekblach/fridge/handler"
	"github.com/witekblach/fridge/handler/ingredient"
	"net/http"
	"time"
)

func NewChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handler.HomepageHandler)

	r.Route("/ingredients", func(r chi.Router) {
		r.Get("/", handler.HomepageHandler)
		r.Post("/", ingredient.Create)
		//	r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017
		//
		//	r.Post("/", createArticle)                                        // POST /articles
		//	r.Get("/search", searchArticles)                                  // GET /articles/search
		//
		//	// Regexp url parameters:
		//	r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto
		//
		//	// Subrouters:
		r.Route("/{ingredientName}", func(r chi.Router) {
			r.Use(IngredientsCtx)
			r.Delete("/", ingredient.Delete)
			//r.Get("/", getArticle)       // GET /articles/123
			//r.Put("/", updateArticle)    // PUT /articles/123
			//r.Delete("/", deleteArticle) // DELETE /articles/123
		})
	})

	return r
}

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
