package encapsutils

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	convey.Convey("async use stack", t, func() {
		st := NewStack()
		var wg sync.WaitGroup
		do := func() {
			for i := 0; i < 10; i++ {
				if !st.Push(time.Now().Unix()) {
					t.Fatal("stack push with failed")
				}
			}
			for i := 0; i < 10; i++ {
				if st.Pop() == nil {
					t.Fatal("stack pop with nil")
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
		convey.So(st.Size(), convey.ShouldEqual, 0)
		first := time.Now().Unix()
		end := time.Now().Unix()
		st.Push(first)
		for i := 0; i < 10; i++ {
			st.Push(time.Now().Unix())
		}
		st.Push(end)
		convey.So(st.Top(), convey.ShouldEqual, end)
		var last interface{}
		for !st.Empty() {
			last = st.Pop()
		}
		convey.So(last, convey.ShouldEqual, first)
	})
}
