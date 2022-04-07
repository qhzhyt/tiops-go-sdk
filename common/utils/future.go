package utils

import (
	"go.uber.org/atomic"
	"sync"
	"time"
)

type Future struct {
	p     atomic.Float64
	store map[string]interface{}
	err   error
	v     interface{}
	ch    chan bool
	done  bool
	msg   string
	ok    bool
	mu sync.Mutex
}

func (f *Future) Get() (interface{}, error) {
	<-f.ch
	return f.v, f.err
}

func (f *Future) closeChan() {
	select {
	case  <- f.ch:

	default:
		close(f.ch)
	}
}

func (f *Future) SetResult(value interface{}) {
	f.v = value
	f.done = true
	f.ok = true
	f.closeChan()
}

func (f *Future) SetError(err error) {
	f.err = err
	f.done = true
	f.closeChan()
}

func (f *Future) SetMessage(msg string) {
	f.msg = msg
}

func (f *Future) GetMessage() string {
	return f.msg
}

func (f *Future) GetError() error {
	return f.err
}


func (f *Future) Progress() float64 {
	return f.p.Load()
}

func (f *Future) SetProgress(p float64) {
	f.p.Store(p)
}

func (f *Future) AddProgress(dp float64) {
	f.p.Add(dp)
}

func (f *Future) OnProgress(f0 func(p float64)) *Future {
	p0 := f.Progress()
	for !f.IsDone() {
		if p0 != f.Progress() {
			p0 = f.Progress()
			f0(p0)
		}
		time.Sleep(time.Millisecond * 500)
	}
	return f
}

func (f *Future) OnFinished(f0 func(f1 *Future)) *Future {
	go func() {
		f.Get()
		f0(f)
	}()
	return f
}

func (f *Future) OnResult(f0 func(v interface{})) *Future {
	go func() {
		v, err := f.Get()
		if err == nil {
			f0(v)
		}
	}()
	return f
}

func (f *Future) OnError(f0 func(e error)) *Future {
	go func() {
		_, err := f.Get()
		if err != nil {
			f0(err)
		}
	}()
	return f
}

func (f *Future) IsDone() bool {
	return f.done
}

func (f *Future) IsSuccess() bool {
	return f.ok
}

func (f *Future) PutData(k string, v interface{}) {
	if f.store == nil {
		f.mu.Lock()
		defer f.mu.Unlock()
		if f.store == nil {
			f.store = map[string]interface{}{}
		}
	}
	f.store[k] = v
}

func (f *Future) GetData(k string) interface{} {
	if f.store == nil {
		return nil
	}
	return f.store[k]
}

func (f *Future) Go(f0 func(f1 *Future)) *Future {
	go f0(f)
	return f
}

func RunFuture(f func(future *Future)) *Future {
	future := &Future{ch: make(chan bool)}
	go f(future)
	return future
}
