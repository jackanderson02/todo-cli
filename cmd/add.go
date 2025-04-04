/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"regexp"
	"strconv"
	"fmt"
	"time"
	"todo/todo"
	"todo/format"
	"github.com/spf13/cobra"
)

const DEFAULT_TO_DO_TIME = "2"
const DEFAULT_TO_DO_UNIT = "h"
const DEFAULT_TO_DO_TIMEUNIT= DEFAULT_TO_DO_TIME + DEFAULT_TO_DO_UNIT

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			format.ShowErrorMessage("this command requires at least one argument. \n proper usage: td add x ?(d/h). For example td add 'have a pint' 1d to give yourself a day to have a pint")
		}

		// Perform regex matching on days and hours arguments
		// td add x 1h
		timeArg := GetArgString(args, 1, DEFAULT_TO_DO_TIMEUNIT)
		regex := regexp.MustCompile(`^(\d+)([hd])$`)

		var matchedTime = DEFAULT_TO_DO_TIME
		var matchedUnit = DEFAULT_TO_DO_UNIT
		if regex.MatchString(timeArg){
			groups := regex.FindStringSubmatch(timeArg)
			matchedTime = groups[1]
			matchedUnit = groups[2]
		}else{
			format.ShowWarningMessage(fmt.Sprintf("Failed to determine intended duration, using default %s", matchedTime))
		}
		createdAt := time.Now()

		var doBy time.Time;
		intTime, _ := strconv.Atoi(matchedTime)
		if matchedUnit == "h"{
			doBy = createdAt.Add(time.Hour * time.Duration(intTime))
		}else{
			doBy = createdAt.AddDate(0,0, intTime)
		}

		td := todo.ToDoListItem{
			Do:  GetArgString(args, 0, "Nothing"),
			DoBy: doBy,
			CreatedAt: createdAt,
		}

		err := todo.DefaultToDoListSqlite().Add(&td)

		if err != nil{
			format.ShowErrorMessage(err.Error())
		}else{
			format.ShowSuccessMessage(fmt.Sprintf("Added %s", td.String()))
		}

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
