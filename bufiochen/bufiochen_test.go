package bufiochen

import (
	"bufio"
	"bytes"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestNewreader(t *testing.T) {
	str:="处免费的设计费垃圾发来的数据奥拉夫就是打发第三方sdffa" +
		"fdsafdsafdsadsfsdfdsfsddddddddddddddddddddddddd大幅度" +
		"dfdddddddddddddddddddddddddddddddddddfdsadfasfddsfads"


	Convey("根据给定的ioreader 创建一个带缓冲的reader",t, func() {
		strRd:=strings.NewReader(str)
		bufiord:= bufio.NewReader(strRd)
		temp:=make([]byte,30)
		bufiord.Read(temp)
		So(string(temp),ShouldEqual,str)
	})
	Convey("根据给定的ioreader 指定缓冲大小 创建一个带缓冲的reader",t, func() {
		strRd:=strings.NewReader(str)
		bufiord:= bufio.NewReaderSize(strRd,1024)
		temp:=make([]byte,30)
		bufiord.Read(temp)
		So(string(temp),ShouldEqual,str)
	})
}

func TestGenReader(t *testing.T) {
	str:="处免费的设计费垃圾发来的数据奥拉夫就是打发第三方sdffa" +
		"fdsafdsafdsadsfsdfdsfsdddd_ddddddddddddddddddddd大幅度" +
		"dfdddddddddddddddddddddddddddddddddddfdsadfasfddsfads"
	Convey("缓冲io的通用reader操作",t,func(){
		strRd:=strings.NewReader(str)
		bufiord:= bufio.NewReader(strRd)
		Convey("清空缓冲区,重设基层reader", func() {
			bufiord.Reset(strRd)
		})
		Convey("获取缓冲中的字节数", func() {
			n:=bufiord.Buffered()
			So(n,ShouldEqual,0)
		})
		Convey("不影响下次读取，预先读取n个缓冲字节", func() {
			temp:=make([]byte,10)
			bufiord.Read(temp)
			peektemp,_:=bufiord.Peek(10)
			So(string(peektemp),ShouldEqual,"peekstr")
		})

	})
	Convey("读取字节 字符 行",t, func() {
		strRd:=strings.NewReader(str)
		bufiord:= bufio.NewReader(strRd)
		Convey("读取字节", func() {
			b,_:=bufiord.ReadByte()
			bufiord.UnreadByte()
			So(b,ShouldEqual,3)
		})
		Convey("读取字符", func() {
			b,n,_:=bufiord.ReadRune()
			bufiord.UnreadRune()
			fmt.Println(b,n)
			So(b,ShouldEqual,3344)
		})

	})

	Convey("按照分隔符读取，readbytes 和readslice类似返回字节切片" +
		"readline不要用，readstring 也是分隔符读取，返回的是string",t, func() {
		strRd:=strings.NewReader(str)
		bufiord:= bufio.NewReader(strRd)
		//temp,_:=bufiord.ReadSlice(byte('_'))
		//SkipSo(string(temp),ShouldEqual,str)
		//temp2,_:=bufiord.ReadBytes(byte('_'))
		//So(string(temp2),ShouldEqual,str)
		temp3,_:=bufiord.ReadString(byte('_'))
		So(temp3,ShouldEqual,str)

	})
}

func TestWrit2(t *testing.T) {
	Convey("缓冲reader 附加实现了writeto 接口",t, func() {
		str:="床前明月光，疑似地上霜"
		strd:=strings.NewReader(str)
		strbfd:=bufio.NewReader(strd)
		var temp bytes.Buffer
		strbfd.WriteTo(&temp)
		So(temp.String(),ShouldEqual,str)

	})
}

func TestNewWriter(t *testing.T) {
	Convey("在创建新的缓冲writer,可以默认也可以指定缓冲大小",t, func() {
		var temp bytes.Buffer
		w1:=bufio.NewWriter(&temp)
		w2:=bufio.NewWriterSize(&temp,1024)
		fmt.Print(w1,w2)
		Convey("缓存操作 reset  查看缓冲剩余大小、使用大小", func() {
			w1.Reset(&temp)
			w1.WriteString("heloof")
			n:=w1.Buffered()
			So(n,ShouldEqual,6)
			n=w1.Available()
			So(n,ShouldEqual,199)


		})
		Convey("写操作，写字符串 字节 字符", func() {
			w1.WriteString("hello")
			w1.WriteByte('h')
			w1.WriteRune('你')
			w1.Flush()

		})

	})
}

func TestReadfrom(t *testing.T) {
	Convey("writer实现了一个 readfrom",t, func() {
		str :="心无所有，因住不所有"
		strrd:=strings.NewReader(str)
		var temp bytes.Buffer
		wter:=bufio.NewWriter(&temp)
		wter.ReadFrom(strrd)
		So(temp.String(),ShouldEqual,str)


	})
}

func TestScanGen(t *testing.T) {
	Convey("缓冲库实额外实现了一个预读解析功能" +
		"分割之后 可以用文本读取 也可以按照字节切片",t, func() {
		// An artificial input source.
		const input = "1234 5678 1234567901234567890"
		scanner := bufio.NewScanner(strings.NewReader(input))
		// Create a custom split function by wrapping the existing ScanWords function.
		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			advance, token, err = bufio.ScanWords(data, atEOF)
			if err == nil && token != nil {
				_, err = strconv.ParseInt(string(token), 10, 32)
			}
			return
		}

		Convey("自定义分割函数", func() {
			// Set the split function for the scanning operation.
			scanner.Split(split)
			// Validate the input
			for scanner.Scan() {
				fmt.Printf("%s\n", scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input: %s", err)
			}

		})
		Convey("行分割函数", func() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fmt.Println(scanner.Text()) // Println will add back the final '\n'
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		})
		Convey("字分割", func() {
			// An artificial input source.
			const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
			scanner := bufio.NewScanner(strings.NewReader(input))
			// Set the split function for the scanning operation.
			scanner.Split(bufio.ScanWords)
			// Count the words.
			count := 0
			for scanner.Scan() {
				count++
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading input:", err)
			}
			fmt.Printf("%d\n", count)
		})


	})
}