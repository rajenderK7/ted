package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type GenerateContent struct {
	Label      string
	ErrMessage string
}

var items = []string{"Base 64", "SHA 256", "Cancel"}

func InputPrompt() (string, error) {
	index := -1
	var standard string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label:    "Choose an encoding standard for the value",
			Items:    items,
		}

		index, standard, err = prompt.Run()
	}

	if standard == items[2] {
		return "", errors.New("cancelled")
	}

	validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New("please enter a string")
        }
        return nil
    }

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
	}

	templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }
	
	prompt := promptui.Prompt{
		Label: "Enter a base value to encode",
		Templates: templates,
		Validate: validate,
	}

	baseValue, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
	switch standard {
	case "Base 64": return GenerateBase64(baseValue), nil
	case "SHA 256": return GenerateSHA256(baseValue), nil
	default: return "", errors.New("something went wrong")
	}
}