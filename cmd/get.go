/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var dataType string
// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Provide a dataset type to get details for",
	Long: `Upon providing a dataset type, you will be prompted to define filters and then this command will
			gather the response from the delaware data api.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if err := validateFlags(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
			os.Exit(1)
		}
		fmt.Print("You selected data type: " +  dataType)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&dataType, "datatype", "d", "", "Data type to query")
	getCmd.MarkFlagRequired("datatype")
}

func validateFlags() error {
	validDataTypes := []string{"red", "blue", "yellow"}
	for _, validDataType := range validDataTypes {
		if dataType == validDataType {
			// color is valid
			return nil
		}
	}
	// if we're here, color is invalid
	return fmt.Errorf("Value '%s' is invalid for flag 'dataType'. Valid values "+
		"come from the set %v", dataType, validDataTypes)
}
