package main

import "fmt"

type WaitGroup struct {
	sem   chan struct{}
	count int
	done  chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem:  make(chan struct{}, 1),
		done: make(chan struct{}),
	}
}

func main() {
	wg := NewWaitGroup()
	const tasks = 10
	wg.Add(tasks)

	for i := 1; i <= tasks; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Задача %d выполнена\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все задачи завершены")
}

func (wg *WaitGroup) Add(delta int) {
	wg.sem <- struct{}{}
	defer func() {
		<-wg.sem
	}()

	newCount := wg.count + delta

	if newCount < 0 {
		panic("negative WaitGroup counter")
	}

	if wg.count == 0 && newCount > 0 {
		wg.done = make(chan struct{})
	}

	wg.count = newCount

	if delta < 0 && newCount == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.sem <- struct{}{}

	if wg.count == 0 {
		<-wg.sem
		return
	}

	done := wg.done

	<-wg.sem
	<-done
}
