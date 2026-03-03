package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/patosinsuerte/ducky-clip/internal/repository"
	"github.com/patosinsuerte/ducky-clip/internal/utils"
	"github.com/patosinsuerte/ducky-clip/models"
	"github.com/spf13/cobra"
)

// VERSION 2 interactiva para todo usuario

func init() {

	RootCMD.AddCommand(addCMD)

}

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "Save one snippet interactively",
	Run: func(cmd *cobra.Command, args []string) {
		var code string

		//SELE PIDE ELCODIGO AL USER
		if len(args) > 0 {
			code = strings.Join(args, " ")
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("\033[33m[1/5] 📝 Paste your code (press Ctrl+D when finished):\033[0m")
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			code = strings.Join(lines, "\n")
		}

		if strings.TrimSpace(code) == "" {
			fmt.Println("\n\033[31m❌ Quack! The code cannot be empty.\033[0m")
			return
		}

		//FUNC AUXIALR PARA CREAR UN SCANNER LIMPIO
		ask := func(label string) string {
			fmt.Printf("\033[36m%s:\033[0m ", label)
			s := bufio.NewScanner(os.Stdin)
			s.Scan()
			return strings.TrimSpace(s.Text())
		}

		fmt.Println("\n--- 🦆 Snippet Configuration ---")

		// DATOS
		//NOMBRE
		name := utils.SanitizeInput(ask("[2/5] Name"))
		if err := utils.NameValidator(name); err != nil {
			fmt.Printf("\n\033[31m❌ 🦆 Quack!: %v\033[0m\n", err)
			return
		}

		//LENGUAJE
		lang := utils.SanitizeInput(ask("[4/5] Lenguaje (ej: go, js)"))
		if err := utils.LanguageValidator(lang); err != nil {
			fmt.Printf("\n\033[31m❌ 🦆 Quack!: %v\033[0m\n", err)
			return
		}

		//DESCRIPCION
		desc := utils.SanitizeInput(ask("[3/5] Descripción"))
		if err := utils.DescriptionValidator(desc); err != nil {
			fmt.Printf("\n\033[31m❌ 🦆 Quack!: %v\033[0m\n", err)
			return
		}

		//CATEGORIA
		cat := utils.SanitizeInput(ask("[5/5] Categoría"))
		if err := utils.CategoryValidator(cat); err != nil {
			fmt.Printf("\n\033[31m❌ 🦆 Quack!: %v\033[0m\n", err)
			return
		}

		newSnippet := models.Snippet{
			Name:        utils.SanitizeInput(name),
			Language:    utils.SanitizeInput(lang),
			Category:    utils.SanitizeInput(cat),
			Description: utils.SanitizeInput(desc),
			Code:        code,
		}

		err := repository.AddOne(&newSnippet)

		if err != nil {
			fmt.Printf("\n\033[31m❌ 🦆 Quack!: Somenthing went wrong: %v\033[0m\n", err)
			return
		}
		fmt.Printf("\n\033[32m✅ 🦆 Quack!: Snippet '%s' saved successfully.\033[0m\n", name)
	},
}
