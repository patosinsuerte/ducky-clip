package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/patosinsuerte/ducky-clip/internal/repository"
	"github.com/patosinsuerte/ducky-clip/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	RootCMD.AddCommand(copyCMD)
}

var copyCMD = &cobra.Command{
	Use:     "copy [id]",
	Aliases: []string{"cp"},
	Short:   "Copy snippets to the clipboard",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// verificar el argumento
		if len(args) > 0 {
			id := args[0]

			snipet, err := repository.ListById(id)

			if err != nil {
				fmt.Printf("❌ 🦆 Quack!: The ID could not be found %s: %v\n", id, err)
				return
			}

			copy := clipboard.WriteAll(utils.SanitizeInput(snipet.Code))

			if copy != nil {
				fmt.Printf("❌ 🦆 Quack!: The clipboard cannot be accessed: %v\n", err)
				return
			}
			fmt.Printf("📋 Copy!! The snippet \033[33m'%s'\033[0m It's already on your clipboard.\n", snipet.Name)
		}
	},
}
