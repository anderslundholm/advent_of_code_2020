package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/anderslundholm/advent_of_code_2020/day01"
	"github.com/anderslundholm/advent_of_code_2020/day02"
	"github.com/anderslundholm/advent_of_code_2020/day03"
	"github.com/anderslundholm/advent_of_code_2020/day04"
	"github.com/anderslundholm/advent_of_code_2020/day05"
	"github.com/anderslundholm/advent_of_code_2020/day06"
	"github.com/anderslundholm/advent_of_code_2020/day07"
	"github.com/anderslundholm/advent_of_code_2020/day08"
	"github.com/anderslundholm/advent_of_code_2020/day09"
	"github.com/anderslundholm/advent_of_code_2020/day10"
	"github.com/anderslundholm/advent_of_code_2020/day11"
	"github.com/anderslundholm/advent_of_code_2020/day12"
	"github.com/anderslundholm/advent_of_code_2020/day13"
	"github.com/anderslundholm/advent_of_code_2020/day14"
	"github.com/anderslundholm/advent_of_code_2020/day15"
	"github.com/anderslundholm/advent_of_code_2020/day16"
	"github.com/anderslundholm/advent_of_code_2020/day17"
	"github.com/anderslundholm/advent_of_code_2020/day18"
	"github.com/anderslundholm/advent_of_code_2020/day19"
	"github.com/anderslundholm/advent_of_code_2020/day20"
	"github.com/anderslundholm/advent_of_code_2020/day21"
	"github.com/anderslundholm/advent_of_code_2020/day22"
	"github.com/anderslundholm/advent_of_code_2020/day23"
	"github.com/anderslundholm/advent_of_code_2020/day24"
	"github.com/anderslundholm/advent_of_code_2020/day25"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	dayFlag  string
	partFlag string
	rootCmd  = &cobra.Command{
		Use:   "advent_of_code_2020",
		Short: "Advent of Code 2020",
		Long:  `My solutions to the Avent of Code 2020 code challange.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	day01.AppendCommand(rootCmd)
	day02.AppendCommand(rootCmd)
	day03.AppendCommand(rootCmd)
	day04.AppendCommand(rootCmd)
	day05.AppendCommand(rootCmd)
	day06.AppendCommand(rootCmd)
	day07.AppendCommand(rootCmd)
	day08.AppendCommand(rootCmd)
	day09.AppendCommand(rootCmd)
	day10.AppendCommand(rootCmd)
	day11.AppendCommand(rootCmd)
	day12.AppendCommand(rootCmd)
	day13.AppendCommand(rootCmd)
	day14.AppendCommand(rootCmd)
	day15.AppendCommand(rootCmd)
	day16.AppendCommand(rootCmd)
	day17.AppendCommand(rootCmd)
	day18.AppendCommand(rootCmd)
	day19.AppendCommand(rootCmd)
	day20.AppendCommand(rootCmd)
	day21.AppendCommand(rootCmd)
	day22.AppendCommand(rootCmd)
	day23.AppendCommand(rootCmd)
	day24.AppendCommand(rootCmd)
	day25.AppendCommand(rootCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent_of_code_2020.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".advent_of_code_2020" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".advent_of_code_2020")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
