package cmd

import (
    "fmt"

    "github.com/abhinavv9/codee/types"
)


// Initialize the job channel and start the workers
func (s *Server) startWorkers(numWorkers int) (chan Job, chan struct{}) {
	jobCh := make(chan types.Job)
	doneCh := make(chan struct{})
    

	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobCh {
                fmt.Println(job)		        	
            }
		}()
	}

	return jobCh, doneCh
}
