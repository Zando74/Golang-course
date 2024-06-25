package valueobject

import "fmt"

type BaseTaskState int

type TaskState BaseTaskState

func (ts *TaskState) Get() BaseTaskState {
	return BaseTaskState(*ts)
}

const (
	Todo BaseTaskState = iota
	InProgress
	InReview
	Done
)

func (s *TaskState) String() string {
	switch BaseTaskState(*s) {
	case Todo:
		return "TODO"
	case InProgress:
		return "IN PROGRESS"
	case InReview:
		return "IN REVIEW"
	case Done:
		return "DONE"
	}
	return "UNKNOWN"
}

type InvalidTaskStateTransitionError struct {
	Current, Desired *TaskState
}

func (e *InvalidTaskStateTransitionError) Error() string {
	return fmt.Sprintf("Transition from %s to %s Not Allowed", e.Current.String(), e.Desired.String())
}

type InvalidTaskStateError struct{}

func (e *InvalidTaskStateError) Error() string {
	return "Invalid Task Value Provided"
}

var transitionRules = map[BaseTaskState][]BaseTaskState{
	Todo: {
		InProgress,
	},
	InProgress: {
		Todo, InReview,
	},
	InReview: {
		Done, Todo,
	},
}

func transitionAllowed(currentState, desiredState BaseTaskState) bool {
	for i := 0; i < len(transitionRules[currentState]); i++ {
		if transitionRules[currentState][i] == desiredState {
			return true
		}
	}
	return false
}

func NewState(state BaseTaskState) (*TaskState, error) {
	if state < 0 || state > 3 {
		return nil, &InvalidTaskStateError{}
	}

	return (*TaskState)(&state), nil
}

func (initialState *TaskState) GoTo(desiredState BaseTaskState) (*TaskState, error) {
	if transitionAllowed(BaseTaskState(*initialState), desiredState) {
		return (*TaskState)(&desiredState), nil
	}
	return nil, &InvalidTaskStateTransitionError{initialState, (*TaskState)(&desiredState)}
}
