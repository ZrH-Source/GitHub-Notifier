package adapter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/webhooks/v6/github"
)

func GitHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := github.New(github.Options.Secret("Whatever"))
	payload, err := hook.Parse(r, github.PushEvent)

	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
			fmt.Println(err)
		}
	}
	if push, v := payload.(github.PushPayload); v {
		repoName := push.Repository.FullName
		repository := strings.Split(repoName, "/")[1]
		infos := GitInfo{author: push.Pusher.Name, repository: repository, email: push.Pusher.Email, hash: push.Commits[0].ID}
		fmt.Println(infos)
		infos.Update()
	}

}
