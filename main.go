package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/witekblach/fridge/data"
	"github.com/witekblach/fridge/db"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	err := actualMain()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func actualMain() error {
	godotenv.Load()

	err := db.NewMongoClient()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", homepageHandler)

	r.Route("/ingredient", func(r chi.Router) {
		r.Post("/", postIngredientHandler)
		//	r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017
		//
		//	r.Post("/", createArticle)                                        // POST /articles
		//	r.Get("/search", searchArticles)                                  // GET /articles/search
		//
		//	// Regexp url parameters:
		//	r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto
		//
		//	// Subrouters:
		//	r.Route("/{articleID}", func(r chi.Router) {
		//		r.Use(ArticleCtx)
		//		r.Get("/", getArticle)                                          // GET /articles/123
		//		r.Put("/", updateArticle)                                       // PUT /articles/123
		//		r.Delete("/", deleteArticle)                                    // DELETE /articles/123
		//	})
	})

	slog.Info("ListenAndServe on " + ":" + os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)

	return nil
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	ingredients, err := data.ShowAllIngredients()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}

func postIngredientHandler(w http.ResponseWriter, r *http.Request) {
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
