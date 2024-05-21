package cli

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/witekblach/fridge/data"
	"github.com/witekblach/fridge/typesense"
	"log/slog"
	"math/rand"
	"strconv"
)

var listCmd = &cobra.Command{
	Use:   "index",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := typesense.NewTypeSenseClient()
		if err != nil {
			return
		}

		recipes := sampleData()

		for _, r := range recipes {
			for _, i := range r.Ingredients {
				document := struct {
					Id               string `json:"id"`
					IngredientName   string `json:"ingredient_name"`
					IngredientAmount string `json:"ingredient_amount"`
				}{
					Id:               strconv.Itoa(rand.Int()),
					IngredientName:   i.Name,
					IngredientAmount: i.Amount,
				}
				slog.Info("info", document)

				typesense.TypeSense.Collection(typesense.CollectionIngredients).Documents().Create(context.Background(), document)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func sampleData() []data.Recipe {
	return []data.Recipe{{
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
}
