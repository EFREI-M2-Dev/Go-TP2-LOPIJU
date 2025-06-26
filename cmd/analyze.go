package cmd

import (
	"errors"
	"fmt"
	"loganizer/internal/analyzer"
	"loganizer/internal/config"
	"loganizer/internal/reporter"
	"sync"

	"github.com/spf13/cobra"
)

var (
	inputJSONFile  string
	outputJSONFile string
)

var AnalyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a log file",
	Long:  `Analyze a log file using a JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputJSONFile == "" {
			fmt.Println("Aucun fichier JSON d'entrée spécifié. Utilisez --config (ou -c) pour fournir les JSONs.")
			return
		}

		configs, err := config.LoadConfig(inputJSONFile)
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la configuration : %v\n", err)
			return
		}

		var wg sync.WaitGroup
		results := make(chan analyzer.LogReport, len(configs))

		for _, logConf := range configs {
			wg.Add(1)
			go func(lc config.LogConfig) {
				defer wg.Done()
				report := analyzer.AnalyzeLog(lc.ID, lc.Path)
				results <- report
			}(logConf)
		}

		wg.Wait()
		close(results)

		reports := make([]analyzer.LogReport, 0, len(configs))
		for r := range results {
			reports = append(reports, r)

			// Gestion avancée des erreurs personnalisées
			var fileErr *analyzer.FileNotFoundError
			var parseErr *analyzer.ParsingError
			msgErreur := ""
			if r.Status == "FAILED" {
				if r.ErrorDetail != "" {
					if errors.As(fmt.Errorf(r.ErrorDetail), &fileErr) {
						msgErreur = "[ERREUR FICHIER] " + fileErr.Error()
					} else if errors.As(fmt.Errorf(r.ErrorDetail), &parseErr) {
						msgErreur = "[ERREUR PARSING] " + parseErr.Error()
					} else {
						msgErreur = r.ErrorDetail
					}
				}
			}

			fmt.Printf("[ID: %s] %s\n  Statut: %s\n  Message: %s\n  Erreur: %s\n",
				r.LogID, r.FilePath, r.Status, r.Message, msgErreur)
		}

		if outputJSONFile != "" {
			err := reporter.ExportReports(reports, outputJSONFile)
			if err != nil {
				fmt.Printf("Erreur lors de l'export du rapport : %v\n", err)
				return
			}
			fmt.Printf("Rapport exporté dans %s\n", outputJSONFile)
		}
	},
}

func init() {
	AnalyzeCmd.Flags().StringVarP(&inputJSONFile, "config", "c", "", "Chemin du fichier de configuration JSON")
	AnalyzeCmd.Flags().StringVarP(&outputJSONFile, "output", "o", "", "Chemin du fichier de rapport JSON (optionnel)")
}
