package pool

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// This is used for holding the task and the result of running a task
type Task struct {
	site string
	Err  error
}

// ThreadPool is a worker group that runs a number of tasks given a max thread count
type ThreadPool struct {
	Tasks       []*Task
	maxCount int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// NewThreadPool creates a new thread pool with the given tasks and
// a given max thread count.
func NewThreadPool(tasks []*Task, maxCount int) *ThreadPool {
	return &ThreadPool{
		Tasks:       tasks,
		maxCount: maxCount,
		tasksChan:   make(chan *Task),
	}
}

// Run runs all the tasks inside the the thread pool
func (tp *ThreadPool) Run() {
	for i := 0; i < tp.maxCount; i++ {
		go tp.work()
	}

	tp.wg.Add(len(tp.Tasks))
	for _, task := range tp.Tasks {
		tp.tasksChan <- task
	}

	// all workers return
	close(tp.tasksChan)

	tp.wg.Wait()
}

// work loops through every single task in the channel and runs the tasks
func (tp *ThreadPool) work() {
	for task := range tp.tasksChan {
		task.Run(&tp.wg)
	}
}

// NewTask creates a new task to be queued in the thread pool
func NewTask(site string) *Task {
	return &Task{site: site}
}

// Run runs a Task and notify the wait group when done
func (t *Task) Run(wg *sync.WaitGroup) {
	response, err := http.Get(t.site)
	if err != nil {
		t.Err = err
		wg.Done()
		return	// If gets an error there's no need to go forward
	}

	defer response.Body.Close()
	fmt.Println("Site: ", t.site)

	// Validates if server responded with MD5 digest
	if response.Header.Get("WWW-Authenticate") != "" {
		wantedHeaders := []string{"nonce", "realm", "qop"}
		result := map[string]string{}
		responseHeaders := strings.Split(response.Header.Get("WWW-Authenticate"), ",")
		for _, rh := range responseHeaders {
			for _, wh := range wantedHeaders {
				if strings.Contains(rh, wh) {
					result[wh] = strings.Split(rh, `"`)[1]
				}
			}
		}
		fmt.Println("Result: ", result)
	}
	wg.Done()
}

func RunNonParallel(tasks []*Task) {
	for _, task := range tasks {
		response, err := http.Get(task.site)
		if err != nil {
			task.Err = err
			continue
		}

		fmt.Println("Site: ", task.site)
		defer response.Body.Close()

		// Validates if server responded with MD5 digest
		if response.Header.Get("WWW-Authenticate") != "" {
			wantedHeaders := []string{"nonce", "realm", "qop"}
			result := map[string]string{}
			responseHeaders := strings.Split(response.Header.Get("WWW-Authenticate"), ",")
			for _, rh := range responseHeaders {
				for _, wh := range wantedHeaders {
					if strings.Contains(rh, wh) {
						result[wh] = strings.Split(rh, `"`)[1]
					}
				}
			}
			fmt.Println("Result: ", result)
		}
	}
}