package main

import (
	"fmt"
	"sync"
)

type TaskIDKey struct {
	TaskID    int
	SubTaskID int
}

var mutexMap = make(map[TaskIDKey]*sync.Mutex)
var mutexMapLock sync.Mutex

func HandleToCheckSaveEntriesInTaskStatusANDUpdateInProjectTasksCalculationAlgo(task_id, sub_task_id int) {
	key := TaskIDKey{TaskID: task_id, SubTaskID: sub_task_id}

	mutexMapLock.Lock()
	mutex, exists := mutexMap[key]
	if !exists {
		mutex = &sync.Mutex{}
		mutexMap[key] = mutex
	}
	mutexMapLock.Unlock()

	mutex.Lock()
	defer mutex.Unlock()

	// Your actual query execution code goes here
	// This code will be executed by only one goroutine at a time for the same key
	// ...

	// Simulate some work
	fmt.Printf("Executing for task_id %d and sub_task_id %d\n", task_id, sub_task_id)
}

func CallToMain2() {
	task_id := 1
	sub_task_id := 2

	for i := 1; i <= 5; i++ {
		go HandleToCheckSaveEntriesInTaskStatusANDUpdateInProjectTasksCalculationAlgo(task_id, sub_task_id)
	}

	fmt.Println("Waiting for goroutines to complete...")
	// fmt.Scanln()
}
