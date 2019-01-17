package stringschen

import (
	"bytes"
	"fmt"
	."github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestStrEqu(t *testing.T){
	Convey("测试字符串相等 忽略大小写",t, func() {
		str1:="hello"
		str2:="HELlo"
		So(strings.EqualFold(str1,str2),ShouldBeTrue)
	})
}

func TestStrPreAndSuf(t *testing.T){
	Convey("判断是否有前缀或者后缀",t, func() {
		str:="abcdEFGbdfd"
		So(strings.HasSuffix(str,"bdfd"),ShouldBeTrue)
		So(strings.HasPrefix(str,"abcd"),ShouldBeTrue)
	})
}

func TestStrContf(t *testing.T){
	Convey("测试子串存在性",t, func() {
		str:="abcdEFGbdfd"
		Convey("子串整体是否存在", func() {
			So(strings.Contains(str,"EFG"),ShouldBeTrue)
		})
		Convey("测试是否存在单个字符", func() {
			So(strings.ContainsRune(str,[]rune("EF")[0]),ShouldBeTrue)
		})
		Convey("测试是否存在多个字符集合中的任何一个", func() {
			So(strings.ContainsAny(str,"EF"),ShouldBeTrue)

		})
		Convey("子串存在几个", func() {
			So(strings.Count(str,"EF"),ShouldEqual,1)
		})
	})
	Convey("子串的存在定位",t, func() {
		str:="abc我们打发打发cdfdf"
		So(strings.Index(str,"打发"),ShouldEqual,9)
		So(strings.IndexByte(str,[]byte("df")[0]),ShouldEqual,22)
		So(strings.IndexRune(str,[]rune("打发")[0]),ShouldEqual,9)
		So(strings.IndexFunc(str, func(r rune) bool {
			if r==[]rune("打发")[0]  || r==[]rune("打发")[1] {
				return true
			}else {
				return false
			}

		}),ShouldEqual,9)


	})
	Convey("存在所有相同的函数 用来逆向查找",t, func() {

	})
}

func TestStr2lowOrUpper(t *testing.T){
	Convey("字符串大小写转换",t, func() {
		str:="womenWdfsdDDDDAFDS阿发fdsdf"
		So(strings.ToLower(str),ShouldEqual,"womenwdfsdddddafds阿发fdsdf")
		So(strings.ToUpper(str),ShouldEqual,"WOMENWDFSDDDDDAFDS阿发FDSDF")
	})

}

func TestStrRepeat(t *testing.T){
	Convey("字符串重复",t, func() {
		str:="abcd"
		So(strings.Repeat(str,3),ShouldEqual,"abcdabcdabcd")
	})
}

func TestStrRepAndMapt(t *testing.T){
	Convey("替换",t, func() {
		str:="abcdaffdsdfsdfsd"
		So(strings.Replace(str,"fs","我们",2),ShouldEqual,"dfds")
	})
	Convey("转换",t, func() {
		str:="abcdaffdsdfsdfsd"
		So(strings.Map(func(r rune) rune {
			if strings.Contains("abdfds",string(r)){
				return r+5
			}else{
				return r+3
			}
		},str),ShouldEqual,"dfds")
	})
}

func TestStrTrim(t *testing.T){
	Convey("剔除字符串的空白或者指定字符",t, func() {
		str:="abdfdfffffffffffffabd"
		str2:="   ffffffffffabd   "
		So(strings.Trim(str,"abd"),ShouldEqual,"fdfffffffffffff")
		So(strings.TrimSpace(str2),ShouldEqual,"ffffffffffabd")
		So(strings.TrimFunc(str, func(r rune) bool {
			return strings.Contains("abc",string(r))

		}),ShouldEqual,"dfdfffffffffffffabd")
		Convey("剔除前面整体字符集或者字符集中的任何一个", func() {
			Convey("剔除整体", func() {
				So(strings.TrimPrefix(str,"abdfd"),ShouldEqual,"fffffffffffffabd")

			})
			Convey("剔除集合中的任何一个", func() {
				So(strings.TrimLeft(str,"af"),ShouldEqual,"fffffffffabd")
			})
			Convey("剔除特定字符", func() {
				So(strings.TrimLeftFunc(str, func(r rune) bool {
					return strings.Contains("abc",string(r))

				}),ShouldEqual,"dfdfffffffffffffabd")

			})

		})
	})
	
	
}

func TestStrFenge(t *testing.T){
	Convey("字符串切割相关",t, func() {
		str:="hello my baby"
		Convey("用空白分割", func() {
			So(len(strings.Fields(str)),ShouldEqual,3)

		})
		Convey("用满足特定函数的字符分割", func() {
			So(len(strings.FieldsFunc(str, func(r rune) bool {
				if r==[]rune(" ")[0] {
					return true
				}else{
					return false
				}

			})),ShouldEqual,3)
		})
		Convey("用字符串作为分割符", func() {
			str2:="helloMYworldMYniMyhao!"
			So(strings.Split(str2,"MY")[1],ShouldEqual,"world")
			So(strings.SplitN(str2,"MY",3)[1],ShouldEqual,"world")
			Convey("在分割符后面 分切片", func() {
				So(strings.SplitAfter(str2,"MY")[1],ShouldEqual,"worldMY")
				So(strings.SplitAfterN(str2,"MY",3)[1],ShouldEqual,"worldMY")
			})
		})

	})
	
}

func TestStrJOin(t *testing.T)  {
	Convey("字符串拼接",t, func() {
		strs:=make([]string,4)
		for i:=0;i<4;i++{
			strs[i]=fmt.Sprintf("hleeos %d",i)
		}
		strt:=strings.Join(strs,"!MY!")
		So(strt,ShouldEqual,"oh!")
	})

}

func TestStrReader(t *testing.T)  {
	Convey("字符串转变为readder",t, func() {
		str:="helorfldjslfjalsdfjlksadjfkldasfadsfdsa"
		strd:=strings.NewReader(str)
		So(strd.Len(),ShouldEqual,len(str))
		temp:=make([]byte,10)
		strd.Read(temp)
		strd.ReadRune()
		strd.UnreadRune()
		strd.ReadByte()
		strd.UnreadByte()
		strd.Seek(0,0)
		strd.ReadAt(temp,10)
		var temp2 bytes.Buffer
		strd.WriteTo(&temp2)

	})

}

func TestStrReplace(t *testing.T)  {
	Convey("字符串替换相关操作",t, func() {
		strep:=strings.NewReplacer("oh","HHA","chuntian","QIUTIAN")
		str:="fdfdoh chuntian ddohdd chuntianfdfd"
		now:=strep.Replace(str)
		SkipSo(now,ShouldEqual,"what ")
		Convey("添加一个writer ，对str 进行替换之后 再写入writer" +
			"string=》替换=》writerr", func() {
			var temp bytes.Buffer
			strep.WriteString(&temp,str)
			So(temp.String(),ShouldEqual,"hahahh")


		})
	})

}