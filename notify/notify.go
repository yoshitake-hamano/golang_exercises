package notify

var sem = make(chan struct{}, 1)
var waitingSync = make(chan struct{}, 1)
var waiting = make(chan struct{}, 1)

func init() {
	sem <- struct{}{}
}

func Lock() {
	<-sem
}

func Wait() {
	sem <- struct{}{}
	<-waiting
	<-sem
}

func NotifyAll() {
	sem <- struct{}{}

	waitingSync <- struct{}{}
	if len(waiting) == 0 {
		waiting <- struct{}{}
	}
	<-waitingSync
}
