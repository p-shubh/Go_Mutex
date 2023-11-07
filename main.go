/* package main

import (
	"fmt"
	"sync"
	"time"
)

type Project struct {
	ID int
	// Add other project fields here
}

 func processRequest(projectID int, mu *sync.Mutex, projectMap map[int]*Project, projectCh map[int]chan struct{}) {
	mu.Lock()
	if projectMap[projectID] == nil {
		projectMap[projectID] = &Project{ID: projectID}
		projectCh[projectID] = make(chan struct{})
	}
	mu.Unlock()

	go func(projectID int) {
		// Check if another request is currently processing this project
		mu.Lock()
		if len(projectCh[projectID]) > 1 {
			fmt.Printf("Delaying request for Project %d\n", projectID)
		}
		mu.Unlock()

		// Wait for the previous request to finish (if any)
		<-projectCh[projectID]

		// Perform some processing (e.g., database update)
		fmt.Printf("Processing request for Project %d\n", projectID)
		time.Sleep(2 * time.Second) // Simulate processing time

		// Release the lock for the next request to proceed
		close(projectCh[projectID])

		// Optional: Update the projectMap or perform any other action
		mu.Lock()
		delete(projectCh, projectID)
		mu.Unlock()
	}(projectID)
}

func main() {
	var (
		mu         sync.Mutex
		projectMap = make(map[int]*Project)
		projectCh  = make(map[int]chan struct{})
	)

	// Simulate concurrent requests
	for i := 1; i <= 5; i++ {
		projectID := 1 // Change this to the desired project ID
		processRequest(projectID, &mu, projectMap, projectCh)
	}

	// Keep the program running for demonstration purposes
	select {}
} */

package main

import (
	"fmt"
	"sync"
	"time"
)

var projectLocks = make(map[int]*sync.Mutex)
var lockMutex sync.Mutex

func HandleProjectManagerDashboardInfoByProjectID(projectID int) {
	lockMutex.Lock()
	if _, exists := projectLocks[projectID]; !exists {
		projectLocks[projectID] = &sync.Mutex{}
	}
	mutex := projectLocks[projectID]
	lockMutex.Unlock()

	mutex.Lock()
	defer mutex.Unlock()

	// Your code to handle the project dashboard info goes here
	fmt.Printf("Handling project %d\n", projectID)
	time.Sleep(2 * time.Second)
}

func main() {
	// Simulate concurrent requests to the function with the same project ID
	projectID := 1
	for i := 1; i <= 2; i++ {
		go HandleProjectManagerDashboardInfoByProjectID(projectID)
	}

	// Simulate requests to the function with different project IDs
	for i := 1; i <= 2; i++ {
		go HandleProjectManagerDashboardInfoByProjectID(i + 1)
	}

	// Wait for all goroutines to finish
	time.Sleep(5 * time.Second)
}
