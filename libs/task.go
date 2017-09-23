package libs

//Task is a common struct
type Task struct {
	tasks []func()
}

//NewTask is a construct func of task
func NewTask(task ...func()) *Task {
	return &Task{
		tasks: task,
	}
}

//Run is a starting func of task
func (t *Task) Run() {
	for _, childTask := range t.tasks {
		go childTask()
	}
}
