package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sync"
)

var (
	inputJSONFile string
)

var AnalyzeCmd = &cobra.Command{
	Use: "analyze",
	Short: "Analyze a log file",
	Long: `Analyze a log file using a JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {

		if inputJSONFile == "" {
			fmt.Println("Aucun fichier JSON d'entrée spécifié. Utilisez --config (ou -c) pour fournir les JSONs.")
			return
		}

		target, err := os.ReadFile(inputJSONFile)
		if err != nil {
			fmt.Printf("Erreur lors du chargement des cibles depuis le fichier %s : %v\n", inputJSONFile, err)
			return
		}

		//do smt this target

		var wg sync.WaitGroup
	}
}
