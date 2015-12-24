package main

import (
	"fmt"
	"sync"
	"time"
)

type LockedInt struct {
	x int
	// a pointer to a basic mutex lock
	sync.Mutex
	// the '0' value of a lock is considered a
	// valid unlocked lock. This allows us to use
	// the "x: 5" notation when creating our struct
	// without giving a name to the sync.Mutex field

	// we cannot name this field and simulataneously use
	// the "a.Lock()" and "a.Unlock()" notation provided
	// by the mutex

	// thus, the following syntax will create a LockedInt
	// with a useable lock:
	// LockedInt{x: 5}

	// alternatively, we can achieve the same behavior as
	// follows:
	// LockedInt{5, sync.Mutex{}}
	// where sync.Mutex{} gives us a zero-value mutex
	// this is transitively done with the previous syntax
}

func main() {
	sharedStackData()
}

func sharedStackData() {
	shared_lint := LockedInt{x: 5}
	shared_lint.Lock()

	// spawns an anonymous function off in a new thread
	go func() {
		fmt.Println("child thread trying to acquire lock at:", time.Now())
		shared_lint.Lock()
		fmt.Println("child thread lock acquired at:", time.Now(), "with x:", shared_lint.x)
		shared_lint.x = 10
		time.Sleep(2 * time.Second)
		shared_lint.Unlock()
	}()

	time.Sleep(2 * time.Second)
	shared_lint.Unlock()
	// wait just long enough that the other thread gets the lock
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main thread trying to acquire lock at:", time.Now())
	shared_lint.Lock()
	defer shared_lint.Unlock()
	fmt.Println("main thread lock acquired at:", time.Now(), "with x:", shared_lint.x)
}