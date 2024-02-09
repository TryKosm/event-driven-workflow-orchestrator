package main

import (
	"fmt"

	"event-driven-workflow-orchestrator/internal/workflow"
)

func main() {
	ok := workflow.Handle(workflow.Event{ID: "1", Payload: "work"})
	fmt.Println(ok)
}
