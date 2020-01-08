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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("No Data Yet")
			}

			f, err := os.Open(home + "/.hcard")
			if err != nil {
				fmt.Println(err)
				return
			}
			files, err := f.Readdir(-1)
			f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, file := range files {
				fmt.Println(file.Name()[:len(file.Name())-5])
			}
		} else {

			home, _ := os.UserHomeDir()
			dirPath := filepath.Join(home, ".hcard")
			filepath := dirPath + "/" + name + ".json"
			viper.SetConfigFile(filepath)
			viper.AutomaticEnv()

			if err := viper.ReadInConfig(); err == nil {
				fmt.Println("Using config file:", viper.ConfigFileUsed())
			}
			names, description := viper.GetString("name"), viper.GetString("description")
			fmt.Println(names, description)

		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("name", "n", "", "Enter a file name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
