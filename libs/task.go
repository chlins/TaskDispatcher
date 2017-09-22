package libs

type Task struct {
	tasks []func()
}

func NewTask(task ...func()) *Task {
	return &Task{
		tasks: task,
	}
}

func (t *Task) Run() {
	for _, childTask := range t.tasks {
		go childTask()
	}
}
