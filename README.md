# Loganizer

Outil CLI Go pour l'analyse concurrente de fichiers de logs.

## Installation

```bash
cd loganizer
go build -o loganizer
```

## Utilisation

### Analyse de logs

```bash
./loganizer analyze --config config.json --output report.json
```

- `--config` ou `-c` : chemin du fichier de configuration JSON (obligatoire)
- `--output` ou `-o` : chemin du fichier de rapport JSON (optionnel)

### Exemple de fichier de configuration

```json
[
  { "id": "web-server-1", "path": "test_logs/access.log", "type": "nginx-access" },
  { "id": "app-backend-2", "path": "test_logs/errors.log", "type": "custom-app" }
]
```

### Exemple de rapport généré

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable ou inaccessible.",
    "error_details": "Fichier introuvable ou inaccessible: /non/existent/log.log (open /non/existent/log.log: no such file or directory)"
  }
]
```

## Equipe
- Louis LUBINEAU
- Pierre VIPREY
- Julien CHAZAL
