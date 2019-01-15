package iochen

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"strings"
	"testing"
	"time"
)
import . "github.com/smartystreets/goconvey/convey"

func TestReadFromReader(t *testing.T) {
	Convey("测试从reader中读取内容" +
		"这种读取只是读取满切片即可", t,func() {
		str:="how old are you ,i am seven"
		strRd:=strings.NewReader(str)
		temp:=make([]byte,100)
		_,err:=strRd.Read(temp)
		if err!=nil{
			str2:=string(temp)
			So(str2,ShouldEqual,str)
		}

	})

}


func TestWrite2temp(t *testing.T) {
	Convey("测试写函数",t, func() {
		var tempbuf bytes.Buffer
		str:="清如秋菊何妨瘦"
		temp:=make([]byte,100)
		temp2:=make([]byte,100)
		temp=([]byte) (str)

		tempbuf.Write(temp)
		tempbuf.Read(temp2)
		So(string(temp2),ShouldEqual,str)

	})
}
/*pip  没有缓存  有内容读才不会阻塞，所以一般也没什么好的应用场景
读完了才能写，写完了才能读，一条信息 必须处理完了才能继续下一次*/
func TestPip(t *testing.T) {
	rp, wp := io.Pipe()
	for i := 0; i < 20; i++ {
		go generate(wp)
	}
	time.Sleep(1 * time.Second)
	data := make([]byte, 64)
	for x:=1;x<20;x++{
		n, err := rp.Read(data)
		if nil != err {
			log.Fatal(err)
		}
		if 0 != n {
			log.Println("main loop", n, string(data))
		}
		time.Sleep(1 * time.Second)
	}
}

func TestTeeReade(t *testing.T) {
	SkipConvey("测试读写过滤保存器",t , func() {
		str1:="白云上卧，世无知音"
		//str2:="浅尝辄止"
		strRd1:=strings.NewReader(str1)
		var temp bytes.Buffer
		rd3:=io.TeeReader(strRd1,&temp)
		temp1:=make([]byte,32)
		rd3.Read(temp1)
		SkipSo(string(temp1),ShouldEqual,str1)
		So(temp.String(),ShouldEqual,str1)


	})
}

func TestMulReader(t *testing.T) {
	Convey("多个reader 合并为一个," +
		"注意，这个和理解的不一样" +
		"其实是要读几次，每次读一个源",t,func(){
		str1:="白云上卧，世无知音"
		str2:="浅尝辄止"
		str3:="靡不有初鲜克有终"
		str4:="观天之道，执天之行"
		strRd1:=strings.NewReader(str1)
		strRd2:=strings.NewReader(str2)
		strRd3:=strings.NewReader(str3)
		strRd4:=strings.NewReader(str4)
		rdtotal:=io.MultiReader(strRd1,strRd2,strRd3,strRd4)

		temp1:=make([]byte,100)
		rdtotal.Read(temp1)
		temp1[len(str1)]='a'
		rdtotal.Read(temp1)
		//So(n,ShouldEqual,3)
		So(string(temp1),ShouldEqual,(str1))




	})
}


func TestMulWrite(t *testing.T) {
	Convey("测试多个写，同时写",t,func(){
		var temp1 bytes.Buffer
		var temp2 bytes.Buffer
		var temp3 bytes.Buffer
		var temp4 bytes.Buffer
		wd:=io.MultiWriter(&temp1,&temp2,&temp3,&temp4)
		str:="你好啊"
		wd.Write([]byte(str))
		So(temp1.String(),ShouldEqual,str)
		//So(temp2.String(),ShouldEqual,"你哈")

	})




}

func TestNCopy(t *testing.T) {
	Convey("将一个reader中指定长度内容复制到一个writer中",t,func(){
		str1:="不积小流无以成江海"
		strRd1:=strings.NewReader(str1)
		var temp1 bytes.Buffer
		io.CopyN(&temp1,strRd1,10)
		So(temp1.String(),ShouldEqual,str1)
	})

}

func TestCopy(t *testing.T) {
	Convey("将一个reader中所有内容复制到一个writer中",t,func(){
		str1:="不积跬步无以至千里"
		strRd1:=strings.NewReader(str1)
		var temp1 bytes.Buffer
		io.Copy(&temp1,strRd1)
		So(temp1.String(),ShouldEqual,str1)
	})

}

func TestWrintestr(t *testing.T) {
	Convey("将一个字符串写到一个writer中",t,func(){
		str:="本来无一物，何处惹尘埃"
		var temp bytes.Buffer
		io.WriteString(&temp,str)
		So(temp.String(),ShouldEqual,str)
	})
}

func TestatleatReader(t *testing.T) {
	Convey("普通的reader 读不到就算了，这个有个保底",t,func(){
		str:="菩提本无树"
		strd:=strings.NewReader(str)
		temp:=make([]byte,100)
		io.ReadAtLeast(strd,temp,10)
		So(string(temp),ShouldEqual,str)
	})

}

func TestFullReader(t *testing.T) {
	Convey("普通的reader 读不到就算了，这个要读满",t,func(){
		str:="明镜亦非台"
		strd:=strings.NewReader(str)
		temp:=make([]byte,100)
		n,_:=io.ReadFull(strd,temp)
		So(n,ShouldEqual,len(str))
		So(string(temp),ShouldEqual,str)
	})

}


func	TestReadall(t *testing.T) {
	Convey("前面都是根据存放的地方决定读多少，这个readall" +
		"是把源内容全部读出来",t,func(){
		str:="旧缘渐断，新缘不生"
		strd:=strings.NewReader(str)
		temp,_:=ioutil.ReadAll(strd)
		So(string(temp),ShouldEqual,str)





	})

}

func TestReadAbsfile(t * testing.T){
	Convey("一次性读取一个文件的所有内容",t,func(){
		fiepath:="G:\\test\\12.jpg"
		temp,err:=ioutil.ReadFile(fiepath)
		if err!=nil{
			fmt.Println("oh no")
		}	else{
			So(len(temp),ShouldEqual,3)
		}
	})
}


func TestWriteall2file(t *testing.T) {
	Convey("一次性把所有内容写到文件中去",t, func() {
		fiepath1:="G:\\test\\133.jpg"
		fiepath:="G:\\test\\12.jpg"
		temp,_:=ioutil.ReadFile(fiepath)
		ioutil.WriteFile(fiepath1,temp,0666)
		So(1,ShouldEqual,1)



	})
}

func TestListDir(t *testing.T) {
	Convey("只遍历文件夹下面的文件不递归，返回文件名数组",t, func() {
		path:="G:\\test"
		files,_:=ListDir(path,"")
		So(len(files),ShouldEqual,2)

	})

}

func TestWalkDir(t *testing.T) {
	Convey("只遍历文件夹下面的文件递归，返回文件名数组",t, func() {
		path:="G:\\暂用文件夹"
		files,_:=WalkDir(path,"")
		So(len(files),ShouldEqual,2)

	})
}

func TestTempDir(t *testing.T) {
	Convey("在指定问价夹下面创建临时文件夹名字",t, func() {
		dir:="G:\\"
		tempdir,_:=ioutil.TempDir(dir,"")
		So(tempdir,ShouldEqual,"G:\\hahdf")
	})
}
func TestTempFile(t *testing.T) {
	Convey("在指定问价夹下面创建临时文件,并返回文件名字",t, func() {
		dir:="G:\\"
		fileinf,_:=ioutil.TempFile(dir,"0777")
		So(fileinf.Name(),ShouldEqual,"G:\\hahdf")
	})
}