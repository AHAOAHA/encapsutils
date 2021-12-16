package encapsutils

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	convey.Convey("async use queue", t, func() {
		qe := NewQueue()
		var wg sync.WaitGroup
		do := func() {
			for i := 0; i < 10; i++ {
				if !qe.Push(time.Now().Unix()) {
					t.Fatal("queue push failed")
				}
			}
			for i := 0; i < 10; i++ {
				if qe.Pop() == nil {
					t.Fatal("queue pop with nil")
				}
			}
		}
		goruntine := rand.Int() % 100
		for i := 0; i < goruntine; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				do()
			}()
		}
		wg.Wait()
		convey.So(qe.Size(), convey.ShouldEqual, 0)
		first := time.Now().Unix()
		end := time.Now().Unix()
		qe.Push(first)
		for i := 0; i < 10; i++ {
			qe.Push(time.Now().Unix())
		}
		qe.Push(end)
		convey.So(qe.Front(), convey.ShouldEqual, first)
		convey.So(qe.Back(), convey.ShouldEqual, end)
		var last interface{}
		for !qe.Empty() {
			last = qe.Pop()
		}
		convey.So(last, convey.ShouldEqual, end)
	})
}
