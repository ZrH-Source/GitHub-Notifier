package adapter

import (
	"fmt"
	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

func GitHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := github.New(github.Options.Secret("hp_1CUoHFud3LKEh6i3rWBxlevFxmcjhL3UtkJ3"))
	payload, err := hook.Parse(r, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
			fmt.Println(err)
		}
	}
	if push, v := payload.(github.PushPayload); v {
		fmt.Printf("%+v", push)
	}
}
