package cli

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/witekblach/fridge/typesense"
	"log/slog"
)

var createCollectionCmd = &cobra.Command{
	Use:   "createCollection",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := typesense.NewTypeSenseClient()
		if err != nil {
			return
		}

		schema := &api.CollectionSchema{
			Name: typesense.CollectionIngredients,
			Fields: []api.Field{
				{
					Name: typesense.CollectionIngredientName,
					Type: "string",
				},
				{
					Name: typesense.CollectionIngredientAmount,
					Type: "string",
				},
			},
		}

		create, err := typesense.TypeSense.Collections().Create(context.Background(), schema)
		if err != nil {
			slog.Info("%+v", err.Error())
			return
		}

		slog.Info("%+v", create)
	},
}

func init() {
	rootCmd.AddCommand(createCollectionCmd)
}
