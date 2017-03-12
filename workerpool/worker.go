// Worker Module
package goworkerpool

import (
	"log"
)

// Interface to create new Jobs
type Job interface {
	Work() error
	SetWorkerID(ID int)
}

// Struct to define Worker functionality
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
	started    bool
	debug      bool
	verbose    bool
}

// NewWorker return a new instance of worker
func NewWorker(id int, workerPool chan chan Job, debug bool) *Worker {
	return &Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
		started:    false,
		debug:      debug,
		verbose:    verbose,
	}
}

// Start worker
func (w *Worker) Start() {
	w.start = true
	go func() {
		// Add to worker pool
		w.workerPool <- w.jobQueue
		for {
			select {
			case job := <-w.jobQueue:
				job.SetWorkerID(w.ID())

				if err := job.Work(); err != nil {
					w.verbose(w.debug,
						"Error running worker %d: %s\n",
						w.id, err.Error())
				}

			case <-w.quitChan:
				w.verbose(w.debug,
					"Worker %d stopping \n", w.id)

				w.started = false

				return
			}
		}
	}()
}

// Stop Worker
func (w *Worker) Stop() {
	w.quitChan = true
}

// Get Worker ID
func (w *Worker) ID() int {
	return w.id
}

// Check if worker started
func (w *Worker) Started() bool {
	return w.started
}

// Debug workerpool
func (w *Worker) verbose(debug bool, message string,
	args ...interface{}) {
	if debug {
		log.Printf(message, args...)
	}
}
