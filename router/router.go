package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/witekblach/fridge/data"
	"github.com/witekblach/fridge/handler"
	ingredientHandler "github.com/witekblach/fridge/handler/ingredient"
	"time"
)

func NewChiRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handler.HomepageHandler)

	r.Route("/ingredients", func(r chi.Router) {
		r.Get("/", ingredientHandler.Get)
		r.Post("/", ingredientHandler.Create)
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
			r.Use(ingredientHandler.IngredientsCtx)
			r.Get("/", ingredientHandler.Get)
			r.Delete("/", ingredientHandler.Delete)
			//r.Get("/", getArticle)       // GET /articles/123
			//r.Put("/", updateArticle)    // PUT /articles/123
			//r.Delete("/", deleteArticle) // DELETE /articles/123
		})
	})

	return r
}

func sampleData() {
	r := []data.Recipe{{
		Name:        "burger",
		Instruction: "smash meat, slam it on a grill. Buns - oven baked, and you can slice a tomato just for kicks",
		Ingredients: []data.Ingredient{
			{Name: "meat", Amount: "200 g"},
			{Name: "bun", Amount: "one, perhaps"},
			{Name: "tomato", Amount: "one"}}}, {
		Name:        "penne pollo pesto",
		Instruction: "cooka de pasta until good. Fire up that pan, and cook the chicken. Then combine it all :)",
		Ingredients: []data.Ingredient{
			{Name: "pasta penne", Amount: "200 g"},
			{Name: "pesto", Amount: "one jar"},
			{Name: "chicken", Amount: "100 g"}}}}

	fmt.Printf("%+v", r)
}
