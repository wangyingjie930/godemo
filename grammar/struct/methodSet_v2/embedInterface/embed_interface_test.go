/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package embedInterface

import (
	"algorithm/base/struct/methodSet_v2"
	"testing"
)

type TestInterface interface {
	M1()
	M2()
	M3()
}

type T struct {
	TestInterface
}

func (*T) M4() {}

func TestEmbedInterface(test *testing.T) {
	var t T
	var pt *T
	methodSet_v2.PrintMethodSet(&t)  //M1 M2 M3
	methodSet_v2.PrintMethodSet(&pt) //M1 M2 M3 M4
}
