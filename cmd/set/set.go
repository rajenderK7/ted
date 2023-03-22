/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package set

import (
	"fmt"
	"os"

	"github.com/rajenderK7/ted/cmd/set/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	keyValue []string
	key string
	v = viper.New()
)

func setKey(file string, key string, val string) error {
	if val == "" {
		val, err := utils.InputPrompt()
		if err != nil {
			return err
		}
		keyValue := []string{key, val}
		return setKeyValue(file, keyValue)
	}
	v.Set(key, val)
	if err := v.WriteConfigAs(file); err != nil {
		return err
	}
	fmt.Println("Key added to config file")
	return nil
}

func setOnlyKey(file string, key string) error {
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return setKey(file, key, "")
}

func setKeyValue(file string, keyValue_ []string) error {
	// Perform any validation if required
	if _, err := os.Stat(file); err != nil {
		fmt.Println("File doesn't exist")
		return err
	}
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	k_, val_ := keyValue_[0], keyValue_[1]
	return setKey(file, k_, val_)
}

// SetCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set or Add key-value pairs",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		key_, _ := cmd.Flags().GetString("key")
		keyValue_, _ := cmd.Flags().GetStringSlice("keyvalue")
		file := args[0]
		if file == "" {
			fmt.Println("Expected relative file path")
			return
		}
		if key_ != "" {
			if err := setOnlyKey(file, key_); err != nil {
				fmt.Println(err)
			}
		} else if l := len(keyValue_); l > 0 && l <= 2 {
			if err := setKeyValue(file, keyValue_); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	SetCmd.Flags().StringSliceVarP(&keyValue, "keyvalue", "K", []string{}, "Set key value pair")
	SetCmd.Flags().StringVarP(&key, "key", "k", "", "Generate value for the key and set")
}
