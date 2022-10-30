package once_demo

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Student
/*
Once 有两点需要注意:
1. 即使函数panic了, 也只会执行一次
2. 如果两次调用Do, 只会执行第一次调用的
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
	s.English -= 1
	time.Sleep(2 * time.Millisecond)
	s.Math -= 1
}
func (s *Student) Inc() {
	s.Math += 1
	time.Sleep(2 * time.Millisecond)
	s.English += 1
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

func Run() {
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
		}()
	}
	//在WaitGroup中, 我们希望Done只执行一次, 就需要用到Once
	o := sync.Once{}
	for i := 0; i < 3; i++ {
		go func(s *Student) {
			wg.Wait()
			//只会执行一次
			o.Do(func() {
				logrus.Info("DONE")
			})
			s.Score()
		}(s)
	}
	time.Sleep(1 * time.Hour)
}
