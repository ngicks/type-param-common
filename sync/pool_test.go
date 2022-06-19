package sync_test

import (
	"bytes"
	"math/rand"
	sync_ "sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ngicks/type-param-common/sync"
)

func TestPool(t *testing.T) {
	var newCallCount, curCallCount, prevCallCount int64
	p := sync.NewPool(func() *bytes.Buffer {
		atomic.AddInt64(&newCallCount, 1)
		return bytes.NewBuffer([]byte{})
	})

	bufList := make([]*bytes.Buffer, 0)
	for i := 0; i < 10; i++ {
		if got := p.Get(); got == nil {
			t.Fatalf("must be nil but %v", got)
		} else {
			bufList = append(bufList, got)
		}
	}
	if curCallCount = atomic.LoadInt64(&newCallCount); curCallCount < prevCallCount+10 {
		t.Fatalf("new is not called correctly")
	}
	prevCallCount = curCallCount

	for _, buf := range bufList {
		p.Put(buf)
	}

	bufList = bufList[:]

	for i := 0; i < 10; i++ {
		if got := p.Get(); got == nil {
			t.Fatalf("must be nil but %v", got)
		} else {
			bufList = append(bufList, got)
		}
	}
	if curCallCount = atomic.LoadInt64(&newCallCount); curCallCount < prevCallCount {
		t.Fatalf("new is not called correctly")
	}
	prevCallCount = curCallCount

	if got := p.Get(); got == nil {
		t.Fatalf("must be nil but %v", got)
	} else {
		bufList = append(bufList, got)
	}
	if curCallCount = atomic.LoadInt64(&newCallCount); curCallCount < prevCallCount+1 {
		t.Fatalf("new is not called correctly")
	}
}

func TestPoolWithoutNew(t *testing.T) {
	p := sync.NewPool[*bytes.Buffer](nil)

	for i := 0; i < 20; i++ {
		if got := p.Get(); got != nil {
			t.Fatalf("must be nil but %v", got)
		}
	}

	// If race build flag is enabled, Put drops random x elements.
	// We can't rely on how many we put them into the pool.
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))
	p.Put(bytes.NewBuffer([]byte{}))

	var someStored bool
	for i := 0; i < 3; i++ {
		if got := p.Get(); got != nil {
			someStored = true
		}
	}
	if !someStored {
		t.Fatalf("non is stored")
	}
}

func TestPoolRace(t *testing.T) {
	p := sync.NewPool(func() *bytes.Buffer {
		return bytes.NewBuffer([]byte{})
	})

	getWithRandomSleep := func() {
		b := p.Get()
		time.Sleep(time.Duration(rand.Int63n(100)))
		p.Put(b)
	}

	wg := sync_.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			getWithRandomSleep()
			wg.Done()
		}()
	}

	wg.Wait()
}
