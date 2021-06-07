package adapter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/webhooks/v6/github"
)

type gitInfo struct {
	author     string
	repository string
	email      string
	hash       string
}

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
		str := push.Repository.FullName
		repository := strings.Split(str, "/")[1]
		infos := gitInfo{author: push.Pusher.Name, repository: repository, email: push.Pusher.Email, hash: push.Commits[0].ID}
		fmt.Println(infos)
	}

}
