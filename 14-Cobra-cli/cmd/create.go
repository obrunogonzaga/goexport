/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/obrunogonzaga/pos-go-expert/14-Cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category`,
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		_, err := categoryDb.Create(categoryName, description)
		if err != nil {
			return err
		}
		return nil
	}
}

var categoryName string
var description string

func init() {
	createCmd := newCreateCmd(GetCatergoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&categoryName, "name", "n", "", "Category name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
