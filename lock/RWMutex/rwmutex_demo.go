package RWMutex

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Student
/*
我们用English和Math相同表示合理的并发
读写锁:
多个写操作是不可并发的, 写锁可以阻塞两个同时进行的写goroutine
多个读操作可以并发
读写操作不可并发, 读锁中不可写, 反之写锁中不可读

注意, 这里的读写并不是通常意义上的读写:
写指的是Lock/Unlock中间的部分
读是RLock/RUnlock中间的部分
*/
type Student struct {
	Name    string
	English int64
	Math    int64
	sync.RWMutex
}

func (s *Student) Score() {
	logrus.Infof("English:%d, Math:%d", s.English, s.Math)
}
func (s *Student) Dec() {
	s.Math -= 1
	time.Sleep(2 * time.Millisecond)
	s.English -= 1
}
func (s *Student) Inc() {
	s.Math += 1
	time.Sleep(2 * time.Millisecond)
	s.English += 1
}
func Run() {
	s := &Student{
		Name:    "KuJyo",
		English: 0,
		Math:    0,
	}
	//启动100个goroutine
	for i := 0; i < 100; i++ {
		go func() {
			for r := 0; r < 10; r++ {
				s.Inc()
				s.Dec()
			}
			s.Score()
		}()
	}
	time.Sleep(1 * time.Hour)
}
func Run_() {
	s := &Student{
		Name:    "KuJyo",
		English: 0,
		Math:    0,
	}
	//启动100个goroutine
	for i := 0; i < 20; i++ {
		go func() {
			for r := 0; r < 5; r++ {
				//加写锁, 避免两个goroutine同时写, 必须等写完成才能读
				s.Lock()
				s.Inc()
				s.Dec()
				s.Unlock()
			}
			//加读锁, 保证写操作完成后才会执行Score
			s.RLock()
			s.Score()
			s.RUnlock()
		}()
	}
	time.Sleep(1 * time.Hour)
}
