package cmd

import (
	"fmt"

	"github.com/patosinsuerte/ducky-clip/internal/repository"
	"github.com/patosinsuerte/ducky-clip/models"
	"github.com/spf13/cobra"
)

func init() {
	RootCMD.AddCommand(listCMD)
}

var listCMD = &cobra.Command{
	Use:     "list [id]",
	Aliases: []string{"ls"},
	Short:   "List saved snippets (all or by ID)",
	Long:    "Displays all snippets saved in your pond, or a specific one if an ID is provided.",
	Args:    cobra.MaximumNArgs(1),
	Example: `  ducky list          # List all snippets
  ducky ls 5 # Show details for snippet with ID 5`,
	Run: func(cmd *cobra.Command, args []string) {

		////CON IDDDD
		if len(args) > 0 {
			id := args[0]
			snippet, err := repository.ListById(id)
			if err != nil {
				fmt.Printf("\033[31m❌ Error: Snippet with ID %s not found: %v\033[0m\n", id, err)
				return
			}

			renderSingleSnippet(snippet)
			return
		}

		//// SIN NADA EL COMANDO PELAO
		snippets, err := repository.ListAll()
		if err != nil {
			fmt.Printf("\033[31m❌ Error reading database: %v\033[0m\n", err)
			return
		}

		if len(snippets) == 0 {
			fmt.Println("🦆 Your pond is empty. Use 'add' to save your first snippet!")
			return
		}
		renderAllSnippets(snippets)
	},
}

// FUNCIONES PRIVADAS
func renderSingleSnippet(s *models.Snippet) {
	fmt.Printf("\n\033[33m--- 🦆 SNIPPET DETAILS --- \033[0m\n")
	fmt.Printf("🆔 ID:       %d\n", s.ID)
	fmt.Printf("📝 Name:   %s\n", s.Name)
	fmt.Printf("🌐 Language: %s\n", s.Language)
	fmt.Printf("📂 Category:      %s\n", s.Category)
	fmt.Printf("\n\033[36mCódigo:\033[0m\n%s\n", s.Code)
	fmt.Println("\033[33m----------------------------\033[0m")
	fmt.Println("\n💡 Use 'ducky cp [id]' to copy the snippet.")

}

func renderAllSnippets(snippets []models.Snippet) {
	fmt.Printf("\n--- 🦆 SAVED SNIPPETS (%d) ---\n\n", len(snippets))
	fmt.Printf("%-5s %-20s %-15s\n", "ID", "NAME", "LANGUAGE")
	fmt.Println("--------------------------------------------------")
	for _, s := range snippets {
		fmt.Printf("\033[33m[%-3d]\033[0m %-20s \033[36m%-15s\033[0m\n", s.ID, s.Name, s.Language)
	}
	fmt.Println("\n💡 Use 'ducky ls [id]' to see the complete code.")
}

//* TODO:::: HACER LAS BUSQUEDAS POR CATEGORIAAA NAME, ETCCCC
