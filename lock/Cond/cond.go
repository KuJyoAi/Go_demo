package Cond

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Student
/*
Cond 可以按条件唤醒一组goroutine
等待模板:
c := c.NewCond(&sync.Mutex{})
c.L.Lock()
if condition {
	c.Wait()
}
//code1
c.L.Unlock();
当Wait()不在执行时, 使用Signal()或者Broadcast()唤醒一个和或者全部,
唤醒全部goroutine时按照加入的队列唤醒
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
	c := sync.NewCond(&sync.Mutex{}) //Cond对象
	o := sync.Once{}
	done := false //完成标记
	for i := 0; i < 10; i++ {
		go func(s *Student) {
			//准备打印, 等待唤醒, 而且每次打印会增加1
			c.L.Lock() //这里配合c.Wait()必须锁住
			for !done {
				c.Wait() //等待完成, 注意, 不在wait只是唤醒的必要条件, 还需要用signal或者broadcast唤醒
			}
			o.Do(func() {
				logrus.Info("I'm the first")
				c.Broadcast() //有一个被唤醒, 由他唤醒所有的唤醒所有的goroutine
			})
			s.Score()
			s.Inc()
			c.L.Unlock()

		}(s)
	}
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			for r := 0; r < 5; r++ {
				s.Inc()
				s.Dec()
			}
		}()
	}
	wg.Wait()
	logrus.Info("DONE")
	done = true //表示完成
	c.Signal()  //唤醒一个goroutine
	time.Sleep(1 * time.Hour)
}
