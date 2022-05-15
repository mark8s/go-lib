/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var (
	year, month int
)

func showCalendar() {
	now := time.Now()

	showYear := year
	if showYear == 0 {
		// 默认使用今年
		showYear = int(now.Year())
	}

	showMonth := time.Month(month)
	if showMonth == 0 {
		showMonth = now.Month()
	}

	showTime := time.Date(showYear, showMonth, 1, 0, 0, 0, 0, now.Location())

	weekdays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	for _, weekday := range weekdays {
		fmt.Printf("%5s", weekday)
	}
	fmt.Println()

	for {
		startWd := showTime.Weekday()
		fmt.Printf("%s", strings.Repeat(" ", int(startWd)*5))

		for ; startWd <= time.Saturday; startWd++ {
			fmt.Printf("%5d", showTime.Day())
			showTime = showTime.Add(time.Hour * 24)

			if showTime.Month() != showMonth {
				return
			}
		}
		fmt.Println()
	}
}

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if year < 1000 && year > 9999 {
			fmt.Fprintln(os.Stderr, "invalid year should in [1000, 9999], actual:%d", year)
			os.Exit(1)
		}

		if month < 1 && year > 12 {
			fmt.Fprintln(os.Stderr, "invalid month should in [1, 12], actual:%d", month)
			os.Exit(1)
		}

		showCalendar()
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	dateCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "year to show (should in [1000, 9999]")
	dateCmd.PersistentFlags().IntVarP(&month, "month", "m", 0, "month to show (should in [1, 12]")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
