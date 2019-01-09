package work_queue

// import "context"

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
	Shut    chan bool
	Done    chan bool
	nworker uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO: initialize struct; start nWorkers workers as goroutines
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{}, maxJobs)
	q.Shut = make(chan bool, nWorkers)
	q.Done = make(chan bool, nWorkers)
	q.nworker = nWorkers

	for i := uint(0); i < nWorkers; i++ {
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	// TODO: Listen on the .Jobs channel for incoming tasks. For each task...
	// TODO: run tasks by calling .Run(),
	// TODO: send the return value back on Results channel.
	// TODO: Exit (return) when .Jobs is closed.
	// ctxobject, cancelfunction = context.WithCancel(context.Background())
	for {
		select {
		case j := <-queue.Jobs:
			queue.Results <- j.Run()
		case <-queue.Shut:
			queue.Done <- true
			return

		}
	}

}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO: close .Jobs and remove all remaining jobs from the channel.

	for i := uint(0); i < queue.nworker; i++ {
		queue.Shut <- true
	}

	for {
		select {
		case <-queue.Jobs:
		default:
			return
		}
	}

}
