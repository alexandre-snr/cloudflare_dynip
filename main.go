package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

func getCronString() string {
	value, exists := os.LookupEnv("CRON_STRING")
	if !exists {
		return "@every 5m"
	}
	return value
}

func scheduleTask() *cron.Cron {
	c := cron.New()
	c.AddFunc(getCronString(), task)
	c.Start()
	return c
}

func idleMainThread() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	<-done
}

func gracefulShutdown(c *cron.Cron) {
	ctx := c.Stop()
	<-ctx.Done()
}

func main() {
	task()

	c := scheduleTask()
	idleMainThread()
	gracefulShutdown(c)
}
