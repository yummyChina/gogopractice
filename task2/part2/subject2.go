package part2

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Name string
	Job  func()
}

type Scheduler struct {
	tasks     []Task
	status    map[int]time.Duration
	startTime time.Time
	wg        sync.WaitGroup
	mu        sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		status: make(map[int]time.Duration),
	}
}

func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) Start() {
	s.startTime = time.Now()
	fmt.Printf("开始执行任务调度（共%d个任务）\n", len(s.tasks))
	s.wg.Add(len(s.tasks))
	for _, task := range s.tasks {
		go func(t Task) {
			defer s.wg.Done()
			start := time.Now()
			t.Job()
			duration := time.Since(start)
			s.mu.Lock()
			s.status[t.ID] = duration
			s.mu.Unlock()
		}(task)

	}
	s.wg.Wait()

}

func (s *Scheduler) PrintStats() {
	fmt.Println("\n任务执行统计:")
	for id, duration := range s.status {
		fmt.Printf("任务ID：%-4d 执行时间:%v\n", id, duration.Round(time.Millisecond))
	}
}
