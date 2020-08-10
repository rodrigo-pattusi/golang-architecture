package architecture

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

type Db map[int]Person

func (m Db) Save(n int, p Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	acc := NewMockAccessor(ctrl)

	p := Person{
		First: "James",
	}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	Put(acc, 1, p)

	ctrl.Finish()
}

func ExamplePut() {
	mdb := Db{}
	p := Person{
		First: "James",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)
	// Output: {James}
}
