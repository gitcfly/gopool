package gopool

import (
	"fmt"
	"runtime"
	"sync"
)

type GoPool struct {
	TaskCap  int
	GoLimit  int
	TaskWait *sync.WaitGroup
	TaskErr  error
	TaskChan chan func() error
}

func NewPool(goLimit int, taskCap int) *GoPool {
	if taskCap <= 0 {
		taskCap = 100
	}
	if goLimit <= 0 {
		goLimit = runtime.NumCPU()
	}
	pool := &GoPool{
		GoLimit:  goLimit,
		TaskCap:  taskCap,
		TaskWait: &sync.WaitGroup{},
		TaskChan: make(chan func() error, 100),
	}
	pool.start()
	return pool
}

func (p *GoPool) Execute(task func() error) {
	p.TaskWait.Add(1)
	p.TaskChan <- task
}

func (p *GoPool) Wait() error {
	close(p.TaskChan)
	p.TaskWait.Wait()
	return p.TaskErr
}

// 私有方法，勿调用
func (p *GoPool) start() *GoPool {
	for i := 0; i < p.GoLimit; i++ {
		go func() {
			for task := range p.TaskChan {
				p.warpDo(task)
			}
		}()
	}
	return p
}

// 私有方法，勿调用
func (p *GoPool) warpDo(task func() error) {
	defer func() {
		if err := recover(); err != nil {
			p.TaskErr = fmt.Errorf("recover err=%v", err)
		}
		p.TaskWait.Done()
	}()
	if err := task(); err != nil {
		p.TaskErr = err
	}
}
