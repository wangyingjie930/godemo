/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package interface_test

import (
	"algorithm/base/struct/methodSet_v2"
	"testing"
)

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func TestInterface(test *testing.T) {
	var t T
	var pt *T
	methodSet_v2.PrintMethodSet(&t)
	methodSet_v2.PrintMethodSet(&pt)
}
