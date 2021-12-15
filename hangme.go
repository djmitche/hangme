package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"

	_ "plugin"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: hangme <iterations>  # 0 means exit immediately\n")
		os.Exit(99)
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse argument as number: %v\n", err)
		os.Exit(1)
	}
	if count == 0 {
		os.Exit(0)
	}
	fmt.Printf("%s\n", runtime.Version())

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			hang(count)
		}()
	}
	wg.Wait()
}

func hang(count int) {
	for i := 0; i < count; i++ {
		timer := time.NewTimer(2 * time.Second)
		started := make(chan struct{})
		go func() {
			select {
			case <-started:
				return
			case <-timer.C:
				fmt.Printf("cmd.Start did not return in timeout\n")
			}
		}()
		cmd := exec.Command("ps", "-x", "-o", "etime", "-p", fmt.Sprintf("%d", os.Getpid()))
		err := cmd.Start()
		close(started)
		if err != nil {
			panic(err)
		}
		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	}
	return
}
