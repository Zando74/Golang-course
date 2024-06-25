package valueobject

var (
	TITLE_MAX_LENGTH = 100
)

type TaskTitle string

type TooLongTaskTitleError struct{}

func (taskTitle *TaskTitle) String() string {
	return string(*taskTitle)
}

func (e *TooLongTaskTitleError) Error() string {
	return "Too Long Task Title Provided"
}

func NewTaskTitle(title string) (*TaskTitle, error) {
	if len(title) < TITLE_MAX_LENGTH {
		return (*TaskTitle)(&title), nil
	}

	return nil, &TooLongTaskTitleError{}
}
