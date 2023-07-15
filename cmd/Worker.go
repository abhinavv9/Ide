package cmd

import (
	"fmt"
)

type Job struct {
	UserID string
	Code   string
	Lang   string
	Image  string
}

// Initialize the job channel and start the workers
func (s *Server) startWorkers(numWorkers int) (chan Job, chan struct{}) {
	jobCh := make(chan Job)
	doneCh := make(chan struct{})

	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobCh {
				//
				fmt.Println(job)
				// Handle the job results as needed
			}
		}()
	}

	return jobCh, doneCh
}
