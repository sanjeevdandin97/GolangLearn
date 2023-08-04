package go_mutex

import (
	"fmt"
	"sync"
	"time"
)

type Hits struct {
	n  int
	mx sync.Mutex
}

var CommonParam = 0

func INIT(systemArgument *string) {
	if *systemArgument == "MTX" {
		goMutex()
	} else if *systemArgument == "RWMTX" {
		goReadWriteMutex()
	} else if *systemArgument == "MTXI" {
		goMutexWithSTructs()
	}
}

// MUTEX BASIC
func goMutex() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		w.Add(1)
		go Increment(&w, &m)
	}

	w.Wait()
	fmt.Println("CommonParam: ", CommonParam)
}

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	CommonParam += 1
	mutex.Unlock()
	wg.Done()
}

// READ/ WRITE MUTEX
func goReadWriteMutex() {
	CommonMap := map[int]int{}
	// RWMutex initialization
	mux := &sync.RWMutex{}

	go write(CommonMap, mux)
	go read(CommonMap, mux)
	go read(CommonMap, mux)

	// Stop the program from exiting
	// Error: this loop will spin, using 100% CPU (SA5002)
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

// MUTEX WITH INSTRUCTS
func goMutexWithSTructs() {
	hits := Hits{}
	hits.Inc()

	time.Sleep(time.Second * 10)
}

func (h *Hits) Inc() {
	//Obtain the exclusive lock
	h.mx.Lock()
	//Increment the value
	h.n++
	//Release the exclusive lock
	h.mx.Unlock()
}
