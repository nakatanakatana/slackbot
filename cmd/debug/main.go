package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/nakatanakatana/slackbot"
	"github.com/slack-go/slack"
)

const (
	APIBaseURL = "/api/v1"
)

func debugCallback(m *slack.InteractionCallback) error {
	for _, a := range m.ActionCallback.BlockActions {
		log.Printf("%#v\n", a)
	}

	return nil
}

type Config struct {
	SigningSecret slackbot.SigningSecret `split_words:"true" required:"true"`
}

func main() {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("%#v", cfg)

	verifier := slackbot.NewSecretsVerifierMiddleware(cfg.SigningSecret)
	handler := slackbot.CreateInteractionsHandler(
		slackbot.InteractionFunctionBlockActions(debugCallback),
	)

	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("%s/interaction", APIBaseURL), verifier(handler))
	log.Printf("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
