# GitHub-Notifier

Un service web en Go qui permet de mettre à jour automatiquement vos repositories GitHub via des webhooks.

## Fonctionnalités

- Réception des webhooks GitHub
- Mise à jour automatique des repositories (pull)
- Clonage automatique des nouveaux repositories
- Gestion sécurisée des webhooks avec vérification du secret
- Logs détaillés des opérations

## Prérequis

- Go 1.16 ou supérieur
- Git
- ngrok (pour les tests en local)

## Installation

1. Clonez le repository :
```bash
git clone https://github.com/votre-username/GitHub-Notifier.git
cd GitHub-Notifier
```

2. Installez les dépendances :
```bash
go mod download
```

3. Configurez le secret GitHub (optionnel) :
```bash
export GITHUB_WEBHOOK_SECRET="votre_secret_github"
```

## Configuration

1. Configurez ngrok pour exposer votre serveur local :
```bash
ngrok http 80
```

2. Dans les paramètres de votre repository GitHub :
   - Allez dans "Settings" > "Webhooks"
   - Cliquez sur "Add webhook"
   - URL : `http://votre-url-ngrok/github`
   - Content type : `application/json`
   - Secret : Le même que celui configuré dans `GITHUB_WEBHOOK_SECRET`
   - Événements : Sélectionnez "Just the push event"

## Utilisation

1. Démarrez le serveur :
```bash
go run main.go
```

Le serveur démarrera sur le port 80 et écoutera les webhooks GitHub.

## Structure du projet

```
.
├── adapter/
│   ├── common.go    # Gestion des commandes Git
│   └── github.go    # Gestion des webhooks GitHub
├── main.go          # Point d'entrée de l'application
├── go.mod
└── README.md
```

## Sécurité

- Le secret GitHub est configurable via la variable d'environnement `GITHUB_WEBHOOK_SECRET`
- Les webhooks sont vérifiés pour s'assurer qu'ils proviennent bien de GitHub
- Les erreurs sont gérées de manière sécurisée

## Logs

Le service génère des logs pour :
- Les mises à jour de repositories
- Les erreurs de pull/clone
- Les webhooks reçus
- Les erreurs de validation des webhooks

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :
1. Fork le projet
2. Créer une branche pour votre fonctionnalité
3. Commiter vos changements
4. Pousser vers la branche
5. Ouvrir une Pull Request

## Licence

MIT

