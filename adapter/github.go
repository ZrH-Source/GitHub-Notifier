package adapter

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/webhooks/v6/github"
)

var (
	githubSecret = getEnvOrDefault("GITHUB_WEBHOOK_SECRET", "Whatever")
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GitHandler(w http.ResponseWriter, r *http.Request) {
	hook, err := github.New(github.Options.Secret(githubSecret))
	if err != nil {
		log.Printf("Erreur lors de la création du webhook: %v", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	payload, err := hook.Parse(r, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Printf("Événement non supporté reçu")
			http.Error(w, "Événement non supporté", http.StatusBadRequest)
			return
		}
		log.Printf("Erreur lors du parsing du webhook: %v", err)
		http.Error(w, "Erreur lors du traitement de la requête", http.StatusBadRequest)
		return
	}

	push, ok := payload.(github.PushPayload)
	if !ok {
		log.Printf("Type de payload invalide")
		http.Error(w, "Type de payload invalide", http.StatusBadRequest)
		return
	}

	repoName := push.Repository.FullName
	repository := strings.Split(repoName, "/")[1]
	
	if len(push.Commits) == 0 {
		log.Printf("Aucun commit trouvé dans le payload")
		http.Error(w, "Aucun commit trouvé", http.StatusBadRequest)
		return
	}

	infos := GitInfo{
		author:     push.Pusher.Name,
		repository: repository,
		email:      push.Pusher.Email,
		hash:       push.Commits[0].ID,
	}

	log.Printf("Mise à jour reçue pour le repository %s par %s", repository, push.Pusher.Name)
	
	if err := infos.Update(); err != nil {
		log.Printf("Erreur lors de la mise à jour: %v", err)
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mise à jour effectuée avec succès")
}
