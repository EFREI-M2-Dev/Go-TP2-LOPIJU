package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"loganizer/internal/analyzer"
)

func ExportReports(reports []analyzer.LogReport, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de rapport: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(reports); err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON: %w", err)
	}
	return nil
}
