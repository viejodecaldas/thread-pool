package main

import (
	"flag"
	"fmt"
	"github.com/viejodecaldas/thread-pool/pool"
	"strconv"
	"strings"
)

const defaultCount = 10

func main() {
	parallel := flag.Bool("parallel", false, "-parallel")

	flag.Parse()

	args := flag.Args()

	count, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid pool count")
		count = defaultCount
	}

	if count <= 0 || count >= 10 {
		fmt.Println("Not a valid pool count, defaulting to ", defaultCount)
		count = defaultCount
	}

	tasks := make([]*pool.Task, 0)
	for _, site := range args {
		// I don't care the result just verify if is an url
		_, err := strconv.Atoi(site)
		if err == nil {
			continue
		}
		if !strings.Contains(site, "http://") {
			site = "http://" + site
		}
		tasks = append(tasks, pool.NewTask(site))
	}

	if !*parallel {
		pool.RunNonParallel(tasks)
	} else {
		p := pool.NewThreadPool(tasks, count)
		p.Run()
	}

	// Print errors if any
	for _, task := range tasks {
		if task.Err != nil {
			fmt.Println("An error occurred: ", task.Err)
		}
	}
}