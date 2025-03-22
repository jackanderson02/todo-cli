/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"todo/todo"
	"github.com/spf13/cobra"
)


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	RunE: func(cmd *cobra.Command, args []string) error{
		if len(args) == 0{
			return errors.New("this command requires at least one argument. \n proper usage: td add x ?d ?h")
		}

		todo.DefaultToDoListSqlite().Add(todo.ToDoListItem{
			Do:  todo.GetArgString(args, 0, "Nothing"),
			ByDays:  todo.GetArg(args, 1, 2, todo.StringToInt),
			ByHours:  todo.GetArg(args, 2, 1, todo.StringToInt),
		})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
