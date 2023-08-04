# Go Mutex

Mutex is a programming concept used to provide synchronization in concurrent programs. To avoid conflicts, Mutex in Golang ensures that no two goroutines can access a variable at a given time.

## Introduction

Golang provides us with a very powerful tool like Goroutines that is used to do multi-Threading.

When Two or more Goroutines are dealing with the same variable in a program, which is manipulated frequently, It causes a condition called "Racing".

Racing happens when two or more Goroutines try to update the value of a shared resource.

As this updation causes the value to be some random value making our application behavior undesirable.

To avoid such issues to happen, we can use locks. Mutex in Golang is used to do locking. This is available as a standard package called "sync".

A Mutex locks a shared resource and releases the lock once the updation of the shared resource is completed. This avoids the race condition.

**_Example:_**
Below is an example showing the use of Mutex in Golang, where we try to manipulate a common parameter and print it.

```go
package main

import (
	"fmt"
	"sync"
)

var CommonParam = 0

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	CommonParam += 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		w.Add(1)
		go Increment(&w, &m)
	}

	w.Wait()
	fmt.Println("CommonParam: ", CommonParam)
}
```

**_Output:_** `CommonParam:  100`

## Reader/ Writer Mutex

An RWMutex is a reader/writer mutual exclusion lock. In some conditions we might need locks only for writes and data reading can be done concurrently. In this scenario, we can use RWMutex.

**_Example:_**
The below example shows the use of RWMutex.

```go
import (
	"fmt"
	"sync"
)

func main() {
	CommonMap := map[int]int{}
	// RWMutex initialization
	mux := &sync.RWMutex{}

	go write(CommonMap, mux)
	go read(CommonMap, mux)
	go read(CommonMap, mux)

	//Stop the program from exiting
	for {
	}
}

func write(CommonMap map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 20; i++ {
			mux.Lock()
			fmt.Println("Writing : ", i)
			CommonMap[i] = i
			mux.Unlock()
		}
	}
}

func read(CommonMap map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for _, value := range CommonMap {
			fmt.Println("Reading value : ", value)
		}
		mux.RUnlock()
	}
}
```

**_Output:_**

```sh
Writing :  0
Reading value :  0
Reading value :  0
Writing:  1
Reading value :  0
...
...
...
```

## Adding Mutexes into Structs

It's good practice to keep the mutex close to the field that it wants to protect.

If the main purpose of the mutex is to protect fields in the struct.

It's very convenient to add the mutex as a field of that struct which naturally protects the fields from concurrent access.

The below Example demonstrates the use of mutexes instructs.

```go
package main

import (
	"sync"
	"Time"
)

var Hits struct {
	n int
	mx sync.Mutex
}

func main() {
	hits := Hits{}
        hits.Inc()

	time.Sleep(time.Second * 10)
}

func (h *Hits) Inc(entry string) {
	//Obtain the exclusive lock
	h.mx.Lock()
	//Increment the value
	h.n++
	//Release the exclusive lock
	h.mx.Unlock()
}
```

## Common Pitfalls

Mutexes in Golang provides a great way to handle race condition.

But with great power comes great responsibilities and implementing mutexes badly can cause conditions like starvation.

When locking shared resources we can miss unlocking it or keep the lock for an extended time.

We must explicitly use the lock for required things only.

Using `defer` statements is an excellent option and even better is to have the unlocking as per requirements.

## Resources

[What is Mutex in Golang? - Scaler](https://www.scaler.com/topics/golang/mutex-in-golang/)
