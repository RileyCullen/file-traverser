package traversabledirectory

// Lists possible types that a text buffer action can be.
type textBufferActionType string

const (
	noop = "noop"
	// Relative movement (i.e., <num>[j | k] actions).
	relativeLineChange textBufferActionType = "relative-line-change"
	// Exact line movement (i.e., :<num> actions).
	explicitLineChange textBufferActionType = "explicit-line-change"
)

type textBufferAction struct {
	actionType textBufferActionType
	buffer     []string
}

func newNoopTextBufferAction() *textBufferAction {
	return &textBufferAction{
		actionType: noop,
	}
}

func newRelativeLineChangeTextBufferAction(buffer []string) *textBufferAction {
	return &textBufferAction{
		actionType: relativeLineChange,
		buffer:     buffer,
	}
}

func newExplicitLineChangeTextBufferAction() *textBufferAction {
	return &textBufferAction{
		actionType: explicitLineChange,
		buffer:     []string{},
	}
}
