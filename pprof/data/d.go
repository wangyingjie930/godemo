/**
  @author: wangyingjie
  @since: 2023/5/14
  @desc:
**/

package data

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
