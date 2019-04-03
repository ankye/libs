package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex
var times = 10000000

func add(x *int) {
	for i := 0; i < times; i++ {
		*x++
	}
	wg.Done()
}

func sub(x *int) {
	for i := 0; i < times; i++ {
		*x--
	}
	wg.Done()
}

func addMutex(x *int) {
	for i := 0; i < times; i++ {
		lock.Lock()
		*x++
		lock.Unlock()
	}
	wg.Done()
}

func subMutex(x *int) {
	for i := 0; i < times; i++ {
		lock.Lock()
		*x--
		lock.Unlock()
	}
	wg.Done()
}

func addAtomic(x *int32) {
	for i := 0; i < times; i++ {
		atomic.AddInt32(x, 1)
	}
	wg.Done()
}

func subAtomic(x *int32) {
	for i := 0; i < times; i++ {
		atomic.AddInt32(x, -1)
	}
	wg.Done()
}

func main() {
	x := 0
	start := time.Now()
	wg.Add(2)
	go add(&x)
	go sub(&x)
	wg.Wait()
	fmt.Println("No lock: ", x)
	elasped := time.Since(start)
	fmt.Println(elasped)

	start = time.Now()
	wg.Add(2)
	x = 0
	go addMutex(&x)
	go subMutex(&x)
	wg.Wait()
	fmt.Println("Mutex lock with condition race: ", x)
	elasped = time.Since(start)
	fmt.Println(elasped)

	start = time.Now()
	wg.Add(2)
	var y int32 = 0
	go addAtomic(&y)
	go subAtomic(&y)
	wg.Wait()
	fmt.Println("Atomic CAS with condition race: ", y)
	elasped = time.Since(start)
	fmt.Println(elasped)

	start = time.Now()
	wg.Add(1)
	x = 0
	go addMutex(&x)
	wg.Wait()
	fmt.Println("Mutex lock without condition race: ", x)
	elasped = time.Since(start)
	fmt.Println(elasped)

	start = time.Now()
	wg.Add(1)
	y = 0
	go addAtomic(&y)
	wg.Wait()
	fmt.Println("Atomic CAS without condition race: ", y)
	elasped = time.Since(start)
	fmt.Println(elasped)
}
