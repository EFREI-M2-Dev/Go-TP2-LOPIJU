# Loganizer

Outil CLI Go pour l'analyse concurrente de fichiers de logs.

## Présentation

**Loganizer** est un outil en ligne de commande permettant d'analyser plusieurs fichiers de logs en parallèle, d'en extraire un rapport synthétique, et de gérer les erreurs de façon robuste. Il est destiné aux administrateurs système et développeurs.

## Fonctionnalités principales
- Analyse concurrente de plusieurs logs (goroutines, WaitGroup)
- Gestion avancée des erreurs personnalisées (fichier introuvable, parsing)
- Interface CLI ergonomique (Cobra)
- Import/export JSON (configuration et rapport)
- Code modulaire (internal/)

## Installation

```bash
git clone <repo_url>
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

## Architecture du projet

- `cmd/` : commandes CLI (analyze, root)
- `internal/config` : lecture de la configuration JSON
- `internal/analyzer` : analyse concurrente, erreurs personnalisées, rapport
- `internal/reporter` : export JSON des résultats

## Gestion des erreurs personnalisées
- **FileNotFoundError** : fichier de log introuvable ou inaccessible
- **ParsingError** : erreur de parsing lors de l'analyse du log

L'affichage console distingue ces erreurs pour faciliter le diagnostic.

## Documentation du code
Chaque package et fonction principale est documenté avec des commentaires GoDoc.

## Équipe
- Louis Dupont
- Alice Martin
- Bob Leroy

---

Pour toute question, contactez l'équipe via le dépôt GitHub.

