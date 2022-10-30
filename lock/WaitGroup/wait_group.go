package wg

import (
	"github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
	"time"
)

// Student
/*

 */
type Student struct {
	Name    string
	English int64
	Math    int64
}

func (s *Student) Score() {
	logrus.Infof("English:%d, Math:%d", s.English, s.Math)
}
func (s *Student) Dec() {
	atomic.AddInt64(&s.Math, -1)
	time.Sleep(2 * time.Millisecond)
	atomic.AddInt64(&s.English, -1)
}
func (s *Student) Inc() {
	atomic.AddInt64(&s.Math, 1)
	time.Sleep(2 * time.Millisecond)
	atomic.AddInt64(&s.English, 1)
}
func Run_() {
	s := &Student{
		Name:    "KuJyo",
		English: 0,
		Math:    0,
	}
	//启动100个goroutine
	wg := sync.WaitGroup{}
	wg.Add(20) //需要等待20个goroutine计算完毕(下面)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done() //表示goroutine执行完毕
			for r := 0; r < 5; r++ {
				s.Inc()
				s.Dec()
			}
			s.Score()
			time.Sleep(2000 * time.Millisecond)
		}()
	}
	//wg.Wait() //等待上面的执行完毕
	for i := 0; i < 3; i++ {
		go func(s *Student) {
			wg.Wait()
			logrus.Info("DONE")
			s.Score()
		}(s)
	}
	time.Sleep(1 * time.Hour)
}
