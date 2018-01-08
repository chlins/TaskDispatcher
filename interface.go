package TaskDispatcher

type tasker interface {
	// all tasks needs to implment function run() && close()
	Run()
}
