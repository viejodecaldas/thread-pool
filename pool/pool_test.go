package pool

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask("adjust.com")

	if task == nil {
		t.Fail()
	}

	if task.site != "adjust.com" {
		t.Fail()
	}
}

func TestNewThreadPool(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask("adjust.com"))
	tasks = append(tasks, NewTask("google.com"))
	threadPool := NewThreadPool(tasks, 5)

	if threadPool == nil {
		t.Fail()
	}

	if threadPool.maxCount != 5 {
		t.Fail()
	}

	if len(threadPool.Tasks) != 2 {
		t.Fail()
	}
}

func TestRunNonParallelSuccess(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask("http://adjust.com"))
	tasks = append(tasks, NewTask("http://google.com"))

	RunNonParallel(tasks)

	for _, task := range tasks {
		if task.Err != nil {
			t.Fail()
		}
	}
}

func TestRunNonParallelNonFormattedURL(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask("adjust.com"))

	RunNonParallel(tasks)

	for _, task := range tasks {
		if task.Err == nil {
			t.Fail()
		}
	}
}

func TestRunParallelSuccess(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask("http://adjust.com"))
	tasks = append(tasks, NewTask("http://google.com"))
	tasks = append(tasks, NewTask("http://facebook.com"))
	tasks = append(tasks, NewTask("http://twitter.com"))
	p := NewThreadPool(tasks, 3)
	p.Run()

	for _, task := range tasks {
		if task.Err != nil {
			t.Fail()
		}
	}
}

func TestRunParallelWrongSiteFormat(t *testing.T) {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask("adjust.com"))
	tasks = append(tasks, NewTask("google.com"))
	tasks = append(tasks, NewTask("facebook.com"))
	tasks = append(tasks, NewTask("twitter.com"))
	p := NewThreadPool(tasks, 2)
	p.Run()

	for _, task := range tasks {
		if task.Err == nil {
			t.Fail()
		}
	}
}