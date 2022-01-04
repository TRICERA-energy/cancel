package main

import (
	"fmt"
	"time"

	"github.com/GoAethereal/cancel"
)

func main() {

	{
		sig := cancel.New()
		sig.Cancel()
		<-sig.Done()
		fmt.Println("Termination manual")
	}

	{
		sig := cancel.New().Timeout(3 * time.Second)
		<-sig.Done()
		fmt.Println("Termination timeout")
	}

	{
		sig := cancel.New().Deadline(time.Now().Add(4 * time.Second))
		<-sig.Done()
		fmt.Println("Termination deadline")
	}

	{
		parent := cancel.New()
		child := cancel.New().Propagate(parent)

		parent.Cancel()
		<-child.Done()

		fmt.Println("Termination parent")
	}

	{
		sig := cancel.New().Timeout(2 * time.Second).Deadline(time.Now().Add(2 * time.Second))
		<-sig.Done()
		fmt.Println("Termination mixed")
	}
}
