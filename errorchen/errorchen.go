package errorchen

import (
	"errors"
	"fmt"
	"time"
)

//错误一般 使用errors.New
func genError(errstr string) (err error ){
	err = errors.New(errstr)
	if err != nil {
		return
	}
	return
}

/*另外一种产生错误的方式*/
func genErr2(errstr string) (err error) {
	 err = fmt.Errorf("%s", errstr)
	return
}

/*错误转成字符串*/
/*其实没必要专门实现一个这样的函数，因为错误的定义接口就是专为字符串
，但是可以借鉴一下这个自定义格式或者添加更多的控制*/
func err2str(err error)(str string){
	s:=fmt.Sprintf("%s",err)
	str=s
	return
}


/*自定义实现error 类型 可以包含更多的错误信息，当然了 可以把所有的信息都转换为字符串 表示出来 也可以用更多的字段表示*/
/*需要实现 一个 error*/
type errstruct struct{
	tim time.Time
	errstr string

}

func (est errstruct)error() string{
	s:=fmt.Sprintf("%s and %s",est.tim.String(),est.errstr)
	return s

}




