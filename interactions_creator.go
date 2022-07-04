package slackbot

import "github.com/slack-go/slack"

type CreateInteractionsHandlerOption interface {
	Apply(InteractionCallbackFunctionMap)
}

func CreateInteractionsHandler(options ...CreateInteractionsHandlerOption) InteractionsHandler {
	callbackFunctions := InteractionCallbackFunctionMap{}
	for _, opt := range options {
		opt.Apply(callbackFunctions)
	}

	return InteractionsHandler{
		callbacks: callbackFunctions,
	}
}

// InteractionFunctionDialogCancellation set callback functions for DialogCancellation.
type InteractionFunctionDialogCancellation []InteractionCallbackFunction

func (f InteractionFunctionDialogCancellation) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeDialogCancellation] = f
}

// InteractionFunctionDialogSubmission set callback functions for DialogSubmission.
type InteractionFunctionDialogSubmission []InteractionCallbackFunction

func (f InteractionFunctionDialogSubmission) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeDialogSubmission] = f
}

// InteractionFunctionDialogSuggestion set callback functions for DialogSuggestion.
type InteractionFunctionDialogSuggestion []InteractionCallbackFunction

func (f InteractionFunctionDialogSuggestion) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeDialogSuggestion] = f
}

// InteractionFunctionInteractionMessage set callback functions for InteractionMessage.
type InteractionFunctionInteractionMessage []InteractionCallbackFunction

func (f InteractionFunctionInteractionMessage) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeInteractionMessage] = f
}

// InteractionFunctionMessageAction set callback functions for MessageAction.
type InteractionFunctionMessageAction []InteractionCallbackFunction

func (f InteractionFunctionMessageAction) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeMessageAction] = f
}

// InteractionFunctionBlockActions set callback functions for BlockActions.
type InteractionFunctionBlockActions []InteractionCallbackFunction

func (f InteractionFunctionBlockActions) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeBlockActions] = f
}

// InteractionFunctionBlockSuggestion set callback functions for BlockSuggestion.
type InteractionFunctionBlockSuggestion []InteractionCallbackFunction

func (f InteractionFunctionBlockSuggestion) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeBlockSuggestion] = f
}

// InteractionFunctionViewSubmission set callback functions for ViewSubmission.
type InteractionFunctionViewSubmission []InteractionCallbackFunction

func (f InteractionFunctionViewSubmission) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeViewSubmission] = f
}

// InteractionFunctionViewClosed set callback functions for ViewClosed.
type InteractionFunctionViewClosed []InteractionCallbackFunction

func (f InteractionFunctionViewClosed) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeViewClosed] = f
}

// InteractionFunctionShortcut set callback functions for Shortcut.
type InteractionFunctionShortcut []InteractionCallbackFunction

func (f InteractionFunctionShortcut) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeShortcut] = f
}

// InteractionFunctionWorkflowStepEdit set callback functions for WorkflowStepEdit.
type InteractionFunctionWorkflowStepEdit []InteractionCallbackFunction

func (f InteractionFunctionWorkflowStepEdit) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeWorkflowStepEdit] = f
}
