package main

import (
	"github.com/kosmgco/tldr/routes"
	"github.com/kosmgco/tldr/task"
	"time"
)

func main() {
	go func() {
		for {
			task.Run()
			time.Sleep(time.Hour * 24)
		}
	}()
	routes.Start()
}
