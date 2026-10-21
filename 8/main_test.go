package main

import (
	"testing"
	"testing/synctest"
)

func TestWaitWithoutTasks(t *testing.T) {
	wg := NewWaitGroup()
	wg.Wait()
}

func TestWaitForAllTasks(t *testing.T) {
	wg := NewWaitGroup()
	results := make([]int, 10)
	wg.Add(len(results))

	for i := range results {
		go func(index int) {
			defer wg.Done()
			results[index] = index + 1
		}(i)
	}

	wg.Wait()

	for i, value := range results {
		if value != i+1 {
			t.Errorf("results[%d] = %d, want %d", i, value, i+1)
		}
	}
}

func TestWaitAfterTasksDone(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)
	wg.Done()
	wg.Done()

	wg.Wait()
}

func TestWaitMultipleWaiters(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		wg := NewWaitGroup()
		wg.Add(1)

		returned := make([]bool, 3)
		for i := range returned {
			go func(index int) {
				wg.Wait()
				returned[index] = true
			}(i)
		}

		synctest.Wait()
		for i, done := range returned {
			if done {
				t.Errorf("waiter %d returned before Done", i)
			}
		}

		wg.Done()
		synctest.Wait()
		for i, done := range returned {
			if !done {
				t.Errorf("waiter %d did not return after Done", i)
			}
		}
	})
}

func TestAddNegativeCounter(t *testing.T) {
	wg := NewWaitGroup()

	defer func() {
		if recover() == nil {
			t.Error("Add(-1) must panic when the counter is zero")
		}
	}()

	wg.Add(-1)
}
