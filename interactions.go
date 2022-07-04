package slackbot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/slack-go/slack"
)

type (
	InteractionCallbackFunction    func(m *slack.InteractionCallback) error
	InteractionCallbackFunctionMap map[slack.InteractionType][]InteractionCallbackFunction
)

type InteractionsHandler struct {
	callbacks InteractionCallbackFunctionMap
}

var _ http.Handler = &InteractionsHandler{}

func (h *InteractionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	message, err := parseBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// NOTE: Acknowledgment response
	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(""))
	if err != nil {
		return
	}

	if callbackFuncs, ok := h.callbacks[message.Type]; ok {
		useCallback(message, callbackFuncs...)
	}
}

func parseBody(body io.ReadCloser) (*slack.InteractionCallback, error) {
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll(body): %w", err)
	}

	defer body.Close()

	jsonStr, err := url.QueryUnescape(string(bodyByte)[8:])
	if err != nil {
		return nil, fmt.Errorf("QueryUnescape: %w", err)
	}

	var message slack.InteractionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &message, nil
}

func GetCallbackID(message slack.InteractionCallback) CallbackID {
	var callbackID CallbackID

	switch message.Type {
	case slack.InteractionTypeWorkflowStepEdit:
		callbackID = CallbackID(message.CallbackID)
	case slack.InteractionTypeViewSubmission:
		callbackID = CallbackID(message.View.CallbackID)
	case
		slack.InteractionTypeBlockActions,
		slack.InteractionTypeBlockSuggestion,
		slack.InteractionTypeDialogCancellation,
		slack.InteractionTypeDialogSubmission,
		slack.InteractionTypeDialogSuggestion,
		slack.InteractionTypeInteractionMessage,
		slack.InteractionTypeMessageAction,
		slack.InteractionTypeShortcut,
		slack.InteractionTypeViewClosed:
		log.Println("no CallbackID:", message.Type)
	}

	return callbackID
}

func useCallback(message *slack.InteractionCallback, functions ...InteractionCallbackFunction) {
	var wg sync.WaitGroup

	wg.Add(len(functions))

	for _, function := range functions {
		m := message
		f := function

		go func() {
			defer wg.Done()

			err := f(m)
			if err != nil {
				log.Println(err)
			}
		}()
	}

	wg.Wait()
}
