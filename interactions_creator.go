package slackbot

import "github.com/slack-go/slack"

type CreateInteractionsHandlerOption interface {
	Apply(InteractionCallbackFunctionMap)
}

func CreateInteractionsHandler(options ...CreateInteractionsHandlerOption) *InteractionsHandler {
	callbackFunctions := InteractionCallbackFunctionMap{}
	for _, opt := range options {
		opt.Apply(callbackFunctions)
	}

	return &InteractionsHandler{
		callbacks: callbackFunctions,
	}
}

// InteractionFunctionDialogCancellation set callback functions for DialogCancellation.
func InteractionFunctionDialogCancellation(
	f ...InteractionCallbackFunction,
) InteractionFunctionDialogCancellationStruct {
	return InteractionFunctionDialogCancellationStruct{
		functions: f,
	}
}

type InteractionFunctionDialogCancellationStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionDialogCancellationStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeDialogCancellation] = f.functions
}

// InteractionFunctionDialogSubmission set callback functions for DialogSubmission.
func InteractionFunctionDialogSubmission(
	f ...InteractionCallbackFunction,
) InteractionFunctionDialogSubmissionStruct {
	return InteractionFunctionDialogSubmissionStruct{
		functions: f,
	}
}

type InteractionFunctionDialogSubmissionStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionDialogSubmissionStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeDialogSubmission] = f.functions
}

// InteractionFunctionDialogSuggestion set callback functions for DialogSuggestion.
func InteractionFunctionDialogSuggestion(
	f ...InteractionCallbackFunction,
) InteractionFunctionDialogSuggestionStruct {
	return InteractionFunctionDialogSuggestionStruct{
		functions: f,
	}
}

type InteractionFunctionDialogSuggestionStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionDialogSuggestionStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeDialogSuggestion] = f.functions
}

// InteractionFunctionInteractionMessage set callback functions for InteractionMessage.
func InteractionFunctionInteractionMessage(
	f ...InteractionCallbackFunction,
) InteractionFunctionInteractionMessageStruct {
	return InteractionFunctionInteractionMessageStruct{
		functions: f,
	}
}

type InteractionFunctionInteractionMessageStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionInteractionMessageStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeInteractionMessage] = f.functions
}

// InteractionFunctionMessageAction set callback functions for MessageAction.
func InteractionFunctionMessageAction(
	f ...InteractionCallbackFunction,
) InteractionFunctionMessageActionStruct {
	return InteractionFunctionMessageActionStruct{
		functions: f,
	}
}

type InteractionFunctionMessageActionStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionMessageActionStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeMessageAction] = f.functions
}

// InteractionFunctionBlockActions set callback functions for BlockActions.
func InteractionFunctionBlockActions(
	f ...InteractionCallbackFunction,
) InteractionFunctionBlockActionsStruct {
	return InteractionFunctionBlockActionsStruct{
		functions: f,
	}
}

type InteractionFunctionBlockActionsStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionBlockActionsStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeBlockActions] = f.functions
}

// InteractionFunctionBlockSuggestion set callback functions for BlockSuggestion.
func InteractionFunctionBlockSuggestion(
	f ...InteractionCallbackFunction,
) InteractionFunctionBlockSuggestionStruct {
	return InteractionFunctionBlockSuggestionStruct{
		functions: f,
	}
}

type InteractionFunctionBlockSuggestionStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionBlockSuggestionStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeBlockSuggestion] = f.functions
}

// InteractionFunctionViewSubmission set callback functions for ViewSubmission.
func InteractionFunctionViewSubmission(
	f ...InteractionCallbackFunction,
) InteractionFunctionViewSubmissionStruct {
	return InteractionFunctionViewSubmissionStruct{
		functions: f,
	}
}

type InteractionFunctionViewSubmissionStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionViewSubmissionStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeViewSubmission] = f.functions
}

// InteractionFunctionViewClosed set callback functions for ViewClosed.
func InteractionFunctionViewClosed(
	f ...InteractionCallbackFunction,
) InteractionFunctionViewClosedStruct {
	return InteractionFunctionViewClosedStruct{
		functions: f,
	}
}

type InteractionFunctionViewClosedStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionViewClosedStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeViewClosed] = f.functions
}

// InteractionFunctionShortcut set callback functions for Shortcut.
func InteractionFunctionShortcut(
	f ...InteractionCallbackFunction,
) InteractionFunctionShortcutStruct {
	return InteractionFunctionShortcutStruct{
		functions: f,
	}
}

type InteractionFunctionShortcutStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionShortcutStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeShortcut] = f.functions
}

// InteractionFunctionWorkflowStepEdit set callback functions for WorkflowStepEdit.
func InteractionFunctionWorkflowStepEdit(
	f ...InteractionCallbackFunction,
) InteractionFunctionWorkflowStepEditStruct {
	return InteractionFunctionWorkflowStepEditStruct{
		functions: f,
	}
}

type InteractionFunctionWorkflowStepEditStruct struct {
	functions []InteractionCallbackFunction
}

func (f InteractionFunctionWorkflowStepEditStruct) Apply(m InteractionCallbackFunctionMap) {
	m[slack.InteractionTypeWorkflowStepEdit] = f.functions
}
