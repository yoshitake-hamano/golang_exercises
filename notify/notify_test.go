package notify

import (
	"sync"
	"testing"
)

func TestLockAndNotifyAll(t *testing.T) {
	Lock()
	NotifyAll()
}

func TestLockWaitNotifyAll(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		Lock()
		Wait()
		NotifyAll()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Lock()
		Wait()
		NotifyAll()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Lock()
		NotifyAll()
		wg.Done()
	}()
	wg.Wait()
}

// func TestLockWaitWaitNotifyAll(t *testing.T) {
// 	Lock()
// 	Wait()
// 	Wait()
// 	NotifyAll()
// }
