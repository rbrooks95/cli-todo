/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"encoding/csv"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add a task to your todo list.",
	Long: `To add a task to your todo list. give it a try!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called right now!")
		//  creating the csv file
		prompt := args[0]
		Write(prompt)
		


		

	},
}


func Write(prompt string)  {
	file, err := os.OpenFile("todo.csv", os.O_APPEND| os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil{
			panic(err)
			
		}
		defer file.Close()
		// create the writer 
		writer := csv.NewWriter(file)

		defer writer.Flush()
		if err := writer.Write([]string{prompt}); err != nil {
			fmt.Printf("error writing to csv: %v\n", err)
			return
		}


}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
