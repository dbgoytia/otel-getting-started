package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

// fibonacci returns the nth fibonacci number
func Fibonacci(number uint) (uint64, error) {

	if number < 1 {
		return uint64(0), nil
	}

	var prev, n0, n1 uint64 = 0, 0, 1

	for i := uint(2); i <= number; i++ {
		prev = n1
		n1 = n1 + n0
		n0 = prev
	}

	return n1, nil
}

func main() {
	l := log.New(os.Stdout, "", 0)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
