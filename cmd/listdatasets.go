package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var dataset string

var datasetCmd = &cobra.Command{
	Use:   "listdatasets",
	Short: "List of Government Datasets.",
	Long: `Command will return the available dataset endpoints you can call and query.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dataset called 1")
		for _, set := range getDataSets() {
			fmt.Println(set)
		}
		fmt.Println("dataset called 2")
	},
}

func init() {
	rootCmd.AddCommand(datasetCmd)

	//datasetCmd.Flags().StringVarP(&dataset, "dataset", "d","", "dataset which you want to grab")
	datasetCmd.MarkFlagRequired("extension")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// datasetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// datasetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getDataSets() []string {
	var datasets = []string{"Student Enrollment", "Educator Salary"}
	return datasets
}
