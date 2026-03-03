package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/patosinsuerte/ducky-clip/internal/repository"
	"github.com/spf13/cobra"
)

var rmAllFlag bool

func init() {
	rmCMD.Flags().BoolVarP(&rmAllFlag, "all", "a", false, "Remove all snippets from the pond")
	RootCMD.AddCommand(rmCMD)
}

var rmCMD = &cobra.Command{
	Use:   "rm [id]",
	Short: "Remove one snippet by ID or use --all to empty the tank",
	Args:  cobra.MaximumNArgs(1), // Acepta 0 o 1 argumento (el ID)
	Run: func(cmd *cobra.Command, args []string) {

		// ELIMINA TODO
		if rmAllFlag {
			if confirmarAccion("⚠️  Are you sure you want to remove the ENTIRE pond?") {
				fmt.Println("🚀 Emptying the pond...")
				if err := repository.DeleteAll(); err != nil {
					fmt.Printf("❌ Error deleting everything: %v\n", err)
					return
				}
				fmt.Println("✅ All snippets were deleted.")
			} else {
				fmt.Println("🦆 Operation cancelled. Your ducks are safe.")
			}
			return
		}

		// Eliminar por ID
		if len(args) > 0 {
			id := args[0]
			// TODO: HACER LA CONFIRMACIONNN
			fmt.Printf("🗑️  Removing snippet with ID: %s...\n", id)

			err := repository.DeleteById(id)
			if err != nil {
				fmt.Printf("❌ 🦆 Quack!: The ID could not be deleted %s: %v\n", id, err)
				return
			}
			fmt.Println("✅ Snippet successfully removed.")
			return
		}

		// ERROR MUESTRA LA AYUDA
		fmt.Println("❌ 🦆 Quack!: You must specify an ID (e.g., ducky rm 5) or use --all")
		cmd.Help()
	},
}

// CONFIRMACION DEL DELETE ALL
func confirmarAccion(mensaje string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n): ", mensaje)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	return input == "y" || input == "yes"
}
