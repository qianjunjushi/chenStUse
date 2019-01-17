package byteschen

import (
	"bytes"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestByteCompare(t *testing.T) {
	Convey("字节切片比较相关操作",t,func(){
		Convey("比较两个切片", func() {
			var a,b []byte
			if bytes.Compare(a, b) < 0 {
				// a less b
			}
			if bytes.Compare(a, b) <= 0 {
				// a less or equal b
			}
			if bytes.Compare(a, b) > 0 {
				// a greater b
			}
			if bytes.Compare(a, b) >= 0 {
				// a greater or equal b
			}
			// Prefer Equal to Compare for equality comparisons.
			if bytes.Equal(a, b) {
				// a equal b
			}
			if !bytes.Equal(a, b) {
				// a not equal b
			}
		})

	})

}


func TestChangebyte2Rune(t *testing.T) {
	Convey("将byte切片转为编码整数切片",t, func() {
		str:="春风又绿江南岸"
		a:=[]byte(str)
		r:=bytes.Runes(a)
		So(string(r),ShouldEqual,str)
	})
}

func TestHasSufixorPrefix(t *testing.T) {
	Convey("测试是否有前缀后缀 切片",t,func(){
		str:="春风又绿江南岸"
		a:=[]byte(str)
		So(bytes.HasPrefix(a,[]byte("春风")),ShouldBeTrue)
		So(bytes.HasSuffix(a,[]byte("江南岸")),ShouldBeTrue)
	})
}

func TestChildSlice(t *testing.T) {
	Convey("子切片相关",t, func() {
		str:="春风又绿江南岸江南岸江南"
		a:=[]byte(str)
		b:=[]byte("江南")
		So(bytes.Contains(a,b),ShouldBeTrue)
		So(bytes.Count(a,b),ShouldEqual,3)
		So(bytes.IndexByte(a,'a'),ShouldEqual,-1)
		So(bytes.IndexRune(a,'南'),ShouldEqual,15)
		//测试是否存在任何一个
		So(bytes.IndexAny(a,"绿南岸江又"),ShouldEqual,6)
		Convey("遍历字符串，直到满足函数为真的字符", func() {
			So(bytes.IndexFunc(a, func(r rune) bool {
				if r==[]rune("岸")[0]{
					return true
				}else {
					return false
				}

			}),ShouldEqual,18)


		})



	})
}

func TestGetChildSlicePos(t *testing.T) {
	Convey("测试所有子串出现的位置" +
		"注意字符串可以转rune 或byte " +
		" byte 不能转rune 只能转string",t, func() {
		str:="春风又绿江南岸江南岸江南"
		a:=[]byte(str)
		b:=[]byte("江南")
		res,err:=GetChildSlicePos(a,b)
		if err==nil{
			So(res[0],ShouldEqual,12)
			fmt.Println(res)
			//rs:=[]rune(str)
			res2,err:=ChangebyteIndex2Strindex(str,res)
			if err==nil{
				fmt.Println(res2)
				So(res2[0],ShouldEqual,4)
				res3,err2:=ChangestrIndex2byteindex(str,res2)
				if err2==nil{
					So(res3[0],ShouldEqual,12)
					fmt.Println(res3)
				}
			}

		}else{
			fmt.Println("err happened ")
			fmt.Println(err)
		}
	})
}

func TestChangelowup(t *testing.T) {
	Convey("大小写转换",t, func() {
		str:="abcDEF"
		a:=[]byte(str)
		So(string(bytes.ToUpper(a)),ShouldEqual,"ABCDEF")
		So(string(bytes.ToLower(a)),ShouldEqual,"abcdef")

	})
}

func TestRepeabytes(t *testing.T) {
	Convey("给出一个重复的切片",t,func() {
		str:="春风又绿江南岸"
		a:=[]byte(str)
		b:=bytes.Repeat(a,3)
		So(string(b),ShouldEqual,str)
	})
}

func TestReplacestr(t *testing.T) {
	Convey("替换子串str",t, func() {
		str:="春风又绿江南又绿岸"
		a:=[]byte(str)
		b:=bytes.Replace(a,[]byte("又绿"),[]byte("abcdfdfs"),2)
		So(string(b),ShouldEqual,str)
	})
}

func TestMapstr(t *testing.T) {
	Convey("对str中每个字符进行函数转换为另一个字符",t, func() {
		str:="春风又绿江南又绿岸"
		a:=[]byte(str)
		b:=bytes.Map(func(r rune) rune {
			return r+3
		},a)
		So(string(b),ShouldEqual,str)
	})
}

func TestStrTrimGen(t *testing.T)  {
	Convey("剔除字符前后特定字符或者空白",t, func() {
		str:="我的心的太乱abddfd我的"
		a:=[]byte(str)
		b:=bytes.Trim(a,"的我")
		SkipSo(string(b),ShouldEqual,"心太乱abddfd")
		str2:="   我的心的太乱abddfd   "
		a2:=[]byte(str2)
		c:=bytes.TrimSpace(a2)
		SkipSo(string(c),ShouldEqual,str2)
		str3:="我的心的太乱abddfd的我我"
		a3:=[]byte(str3)
		c3:=bytes.TrimFunc(a3, func(r rune) bool {
			if r==[]rune("我的")[0]  ||  r==[]rune("我的")[1]{
				return true
			}
			return false

		})
		SkipSo(string(c3),ShouldEqual,"心的太乱abddfd")
		Convey("单侧剔除的两种方式的区别", func() {
			str5:="abcdFDAFDSA"
			a5:=[]byte(str5)
			Convey("trimleft 剔除子串中单独任何一个字符", func() {
				b5:=bytes.TrimLeft(a5,"bdca")
				So(string(b5),ShouldEqual,"FDAFDSA")

			})
			Convey("trimprefix 剔除子串作为一个整体", func() {
				b6:=bytes.TrimLeft(a5,"abcd")
				So(string(b6),ShouldEqual,"FDAFDSA")

			})

		})
	})


}

func TestStrFengeGEn(t *testing.T)  {
	Convey("字符串分割相关",t, func() {
		str:="痛苦_ 来源于_ 过去的放肆"
		a:=[]byte(str)
		b:=bytes.Fields(a)
		So(len(b),ShouldEqual,3)
		c:=bytes.FieldsFunc(a, func(r rune) bool {
			if r ==[]rune("_")[0]{
				return true
			}
			return false
		})
		So(len(c),ShouldEqual,3)

	})
	Convey("上面考虑的是单个字节或者字符分割， " +
		"下面考虑多个字节分割",t, func() {
		str:="痛苦_ 来源于_ 过去的放肆"
		a:=[]byte(str)
		b:=bytes.Split(a,[]byte("_ "))
		So(len(b),ShouldEqual,3)
		// n代表的是分割结果是n个子切片
		b2:=bytes.SplitN(a,[]byte("_ "),2)
		So(len(b2),ShouldEqual,2)
		Convey("包含分割串的分割", func() {
			b3:=bytes.SplitAfter(a,[]byte("_ "))
			So(len(b3),ShouldEqual,3)
			b4:=bytes.SplitAfterN(a,[]byte("_ "),2)
			So(len(b4),ShouldEqual,2)
		})
	})

	Convey("字符串分割的逆向操作，字符拼接",t,func(){
		a:=make([][]byte,10)
		for i:=0; i<len(a);i++{
			//b=make([]byte,10)
			a[i]=[]byte(fmt.Sprintf("hello %d",i))
		}
		c:=bytes.Join(a,[]byte("你好"))
		So(string(c),ShouldEqual,"jjjj")
	})

}

func TestSTrReaderGe(t *testing.T)  {
	Convey("字节当成reader 相关操作",t, func() {
		str:="春蚕到死丝方尽，蜡炬成灰泪始干"
		bs:=[]byte(str)
		rd:=bytes.NewReader(bs)
		So(rd.Len(),ShouldEqual,len(str))
		temp:=make([]byte,10)
		n,_:=rd.Read(temp)
		So(n,ShouldEqual,len(temp))
		b,_:=rd.ReadByte()
		rd.UnreadByte()
		fmt.Println(b)
		r,s,_:=rd.ReadRune()
		rd.UnreadRune()
		fmt.Println(r,s)
		rd.Seek(10,0,)
		rd.ReadAt(temp,10)
		var temp2 bytes.Buffer
		rd.WriteTo(&temp2)


	})

}

func TestBfugener(t *testing.T)  {
	Convey("字节的缓冲相关" +
		"缓冲能读能写 很方便",t, func() {
		var temp bytes.Buffer
		srtemp2:=make([]byte,10)
		temp2:=bytes.NewBuffer(srtemp2)
		temp3:=bytes.NewBufferString("用这个字符串填充buffer")
		//清空缓冲
		temp.Reset()
		temp2.Truncate(0)
		bs:=temp3.Bytes()
		So(string(bs),ShouldEqual,"用这个字符串填充buffer")
		st:=temp3.String()
		So(st,ShouldEqual,"用这个字符串填充buffer")
		Convey("增加缓冲的容量，使其能够存放n个字节", func() {
			temp3.Grow(40)
			temp3.Write([]byte("dddddddddd"))
			temp3.Read(srtemp2)
			//跳过n个字节，就像读取过一样
			temp3.Next(3)
			c,_:=temp3.ReadByte()
			temp3.UnreadByte()
			fmt.Println(c)
			r,s,_:=temp3.ReadRune()
			temp3.UnreadByte()
			fmt.Println(r,s)
		})
		temp3.Write([]byte("安心工作，好好锻炼，心慢下来，把病痛当成契机"))
		//line,_:=temp3.ReadBytes(([]byte("慢"))[0])
		//SkipSo(string(line),ShouldEqual,"lll")
		line2,_:=temp3.ReadString(([]byte("慢"))[0])
		SkipSo(string(line2),ShouldEqual,"lll")
		temp3.WriteRune([]rune("helo")[0])
		temp3.WriteByte([]byte("helo")[0])
		temp3.WriteString("nihao ")
		temp3.WriteTo(temp2)
		temp3.ReadFrom(temp2)
	})
}
