/**
  @author: wangyingjie
  @since: 2023/4/23
  @desc:
**/

package Container

//var App *Container

var App = &Container{bindings: map[string]interface{}{}}

type Container struct {
	bindings map[string]interface{}
}

func (c *Container) Single(abstract string, instance interface{}) {
	c.bindings[abstract] = instance
}

func (c *Container) App(abstract string) interface{} {
	return c.bindings[abstract]
}
