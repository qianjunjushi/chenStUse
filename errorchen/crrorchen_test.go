package errorchen

import (
"testing"
. "github.com/smartystreets/goconvey/convey"
	"time"
)

func TestGenErrorandback(t *testing.T) {

	Convey("目的：测试产生错误的函数,并且将错误转换为字符串", t, func() {

		Convey("给出一个字符串，返回一个相应的错误", func() {
			strtest:="出错啦"
			Convey("操作1：是给出错误", func() {
				strtest="出错1"
				Convey("值应该是", func(){
					err1:=genError(strtest)
					So(err2str(err1),ShouldEqual,"出错1")
				})
				Convey("值应该是2", func(){
					err2:=genErr2(strtest)
					err2.Error()
					So(err2str(err2),ShouldEqual,"出错1")
				})

			})







		})

	})

}


func TestCustError(t *testing.T) {

	Convey("目的：测试自定义错误", t, func() {

		Convey("生成一个自定义错误", func() {
			est:=errstruct{time.Now(),"helor owrkd"}
			errs:=est.error()
			So(errs,ShouldEqual,"helor owrkd")

			Convey("判断最后返回字符串", nil)

		})

	})

}
