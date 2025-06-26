// Package analyzer fournit les fonctions d'analyse de logs et les erreurs personnalisées.
package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// LogReport représente le résultat de l'analyse d'un log.
type LogReport struct {
	LogID       string `json:"log_id"`
	FilePath    string `json:"file_path"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	ErrorDetail string `json:"error_details"`
}

// FileNotFoundError est une erreur personnalisée pour un fichier introuvable ou inaccessible.
type FileNotFoundError struct {
	Path string
	Err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("Fichier introuvable ou inaccessible: %s (%v)", e.Path, e.Err)
}

// ParsingError est une erreur personnalisée pour une erreur de parsing de log.
type ParsingError struct {
	Path   string
	Reason string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("Erreur de parsing du log: %s (%s)", e.Path, e.Reason)
}

// AnalyzeLog analyse un fichier de log et retourne un LogReport.
// Elle simule une analyse avec un délai aléatoire et une erreur de parsing aléatoire.
func AnalyzeLog(id, path string) LogReport {
	if _, err := os.Stat(path); err != nil {
		fileErr := &FileNotFoundError{Path: path, Err: err}
		return LogReport{
			LogID:       id,
			FilePath:    path,
			Status:      "FAILED",
			Message:     "Fichier introuvable ou inaccessible.",
			ErrorDetail: fileErr.Error(),
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := time.Duration(50+r.Intn(151)) * time.Millisecond
	time.Sleep(delay)

	if r.Float32() < 0.1 {
		parseErr := &ParsingError{Path: path, Reason: "Format inattendu"}
		return LogReport{
			LogID:       id,
			FilePath:    path,
			Status:      "FAILED",
			Message:     "Erreur de parsing.",
			ErrorDetail: parseErr.Error(),
		}
	}

	return LogReport{
		LogID:       id,
		FilePath:    path,
		Status:      "OK",
		Message:     "Analyse terminée avec succès.",
		ErrorDetail: "",
	}
}
