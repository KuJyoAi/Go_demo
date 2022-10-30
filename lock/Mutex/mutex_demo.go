package Mutex

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Student
/*
我们用English和Math相同表示合理的并发
互斥锁, 用Lock/UnLock之间包含的代码无法并发执行
与读写锁唯一的区别是, 读写锁是细粒度的操作, 允许多个读并存
*/
type Student struct {
	Name    string
	English int64
	Math    int64
	sync.Mutex
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
				s.Lock()
				s.Inc()
				s.Dec()
				s.Unlock()
			}
			//与读写锁的区别在于, 此处Score必然是顺序执行的, 而读写锁的话, Score可以并发执行
			s.Lock()
			s.Score()
			s.Unlock()
		}()
	}
	time.Sleep(1 * time.Hour)
}
