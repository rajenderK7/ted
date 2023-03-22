/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package read

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	file string
)

// ReadCmd represents the read command
var ReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read key value pairs from config files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		v.SetConfigFile(file)
		if err := v.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		config := v.AllSettings()
		fmt.Println(config)
	},
}

func init() {
	// Here Viper is used to read different config files
	// The file here is expected to be an absolute path
	ReadCmd.Flags().StringVarP(&file, "file", "f", "", "Read config from the provided file")
}
