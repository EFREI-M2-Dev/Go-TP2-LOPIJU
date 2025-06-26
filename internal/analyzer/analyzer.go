package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type LogReport struct {
	LogID       string `json:"log_id"`
	FilePath    string `json:"file_path"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	ErrorDetail string `json:"error_details"`
}

type FileNotFoundError struct {
	Path string
	Err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("Fichier introuvable ou inaccessible: %s (%v)", e.Path, e.Err)
}

type ParsingError struct {
	Path   string
	Reason string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("Erreur de parsing du log: %s (%s)", e.Path, e.Reason)
}

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
