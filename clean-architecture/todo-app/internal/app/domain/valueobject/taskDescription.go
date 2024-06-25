package valueobject

var (
	DESCRIPTION_MAX_LENGTH = 300
)

type TaskDescription string

func (taskDescription *TaskDescription) String() string {
	return string(*taskDescription)
}

type TooLongTaskDescriptionError struct{}

func (e *TooLongTaskDescriptionError) Error() string {
	return "Too Long Task Description Provided"
}

func NewTaskDescription(description string) (*TaskDescription, error) {
	if len(description) < DESCRIPTION_MAX_LENGTH {
		return (*TaskDescription)(&description), nil
	}

	return nil, &TooLongTaskDescriptionError{}
}
