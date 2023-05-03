/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package alias

import (
	"algorithm/base/struct/methodSet_v2"
	"testing"
)

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type Interface interface {
	M1()
	M2()
}

type NewT T
type NewInterface Interface

type AliasT = T
type AliasInterface = Interface

func TestAlias(test *testing.T) {
	var t T                //M1
	var pt *T              //M1 M2
	var it Interface       //M1 M2
	var newt NewT          //empty
	var newpt *NewT        //empty
	var newit NewInterface //M1 M2

	var aliast AliasT          //M1
	var aliaspt *AliasT        //M1 M2
	var aliasit AliasInterface //M1 M2

	methodSet_v2.PrintMethodSet(&t)
	methodSet_v2.PrintMethodSet(&pt)
	methodSet_v2.PrintMethodSet(&it)

	methodSet_v2.PrintMethodSet(&newit)
	methodSet_v2.PrintMethodSet(&newt)
	methodSet_v2.PrintMethodSet(&newpt)

	methodSet_v2.PrintMethodSet(&aliast)
	methodSet_v2.PrintMethodSet(&aliaspt)
	methodSet_v2.PrintMethodSet(&aliasit)
}
