package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/witekblach/fridge/typesense"
)

// queryCollectionCmd represents the queryCollection command
var queryCollectionCmd = &cobra.Command{
	Use:   "queryCollection",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := typesense.NewTypeSenseClient()
		if err != nil {
			return
		}

		searchParameters := &api.SearchCollectionParams{
			Q:       "k",
			QueryBy: typesense.CollectionIngredientName,
			//FilterBy: pointer.String("num_employees:>100"),
			//SortBy:   &([]string{"num_employees:desc"}),
		}

		retrieve, err := typesense.TypeSense.Collection(typesense.CollectionIngredients).Documents().Search(context.Background(), searchParameters)
		if err != nil {
			return
		}

		fmt.Printf("%+v", retrieve.SearchTimeMs)
	},
}

func init() {
	rootCmd.AddCommand(queryCollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCollectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCollectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
