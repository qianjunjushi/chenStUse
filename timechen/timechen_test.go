package timechen

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestMonth2str(t *testing.T) {

	Convey("time标准库有月份的常量 具有strng方法 " +
		"可以转换为字符串" +
		"转换一个月到字符串", t, func() {

		Convey("给出一个月常量转变为字符", func(){
			tm1:=time.February
			So(tm1.String(),ShouldEqual,"February")
		})

	})

}


func TestWeekday2str(t *testing.T) {

	Convey("time标准库 具有星期几的常量，具有string方法转换为字符串" +
		"转换一个星期几到字符串", t, func() {

		Convey("给出一个星期几常量转变为字符", func(){
			tm1:=time.Monday
			So(tm1.String(),ShouldEqual,"Monday")
		})

	})

}

func TestLocationGen(t *testing.T){
	Convey("时区 关联一个名字 和一个偏移" +
		"有两个特例时区" +
		"时区有一个名字 可以通过string 方法来得到" +
		"可以用fixedzone 创建一个新的时区", t, func() {
		local:=time.Local
		utc:=time.UTC
		So(local.String(),ShouldEqual,"Local")
		So(utc.String(),ShouldEqual,"UTC")
		nez:=time.FixedZone("nez",3600)
		So(nez.String(),ShouldEqual,"nez")
		So(1, ShouldEqual, 1)
	})

	Convey("获取时间的时区信息",t,func(){
		t1:=time.Now()
		lot1:=t1.Local()

		zone,of:=lot1.Zone()
		So(zone,ShouldEqual,"CST")
		So(of,ShouldEqual,28800)
		//So(lot1.String(),ShouldEqual,"local11")
		//从时间中获取的时区信息 包含时间


	})

}

func TestTimeGen(t *testing.T) {
	Convey("时间 主要由 时区 时间 时间间隔 相关", t, func() {
		Convey("时间部分 时间生成 时间解析 时间操作", func() {

			Convey("时间生成：", func() {

				Convey("给出年月日时分秒 纳秒 还有时区 创建一个时间", func() {
					t1 := time.Date(2019, 1, 7, 13, 23, 07, 1, time.UTC)
					So(t1.Year(), ShouldEqual, 2019)
				})
				Convey("获取当前时间 创建一个时间", func() {
					t1 := time.Now()
					So(t1.Year(), ShouldEqual, 2019)
				})
				Convey("解析字符串 创建一个时间", func() {
					t1string := "2016-01-02T15:04:05Z0700"
					t1, _ := time.Parse(time.RFC3339, t1string)

					So(t1.Year(), ShouldEqual, 2016)

				Convey("解析字符串 创建一个时间 附带ns", func() {
					t1string := "2019-01-02T15:04:05.33339Z0800"
					t1, _ := time.Parse(time.RFC3339Nano, t1string)
					So(t1.Year(), ShouldEqual, 2019)
				})


				Convey("给出unix 时间 和ns  创建一个本地时间",func(){
					t1:=time.Unix(1545264000,399493)
					So(t1.Year(),ShouldEqual,2018)
					So(t1.String(),ShouldEqual,"2018")
				})


				})

			})

			Convey("时间解析",func(){
				Convey("获取时间中的时区信息",func(){
					t1:=time.Now()
					loca1:=t1.Location()
					lot1:=t1.Local()
					So(loca1.String(),ShouldEqual,"Local")
					zone,of:=lot1.Zone()
					So(zone,ShouldEqual,"CST")
					So(of,ShouldEqual,28800)

				})

				Convey("从时间中解析出unix 秒和纳秒",func(){
					t1:=time.Now()
					s:=t1.Unix()
					ns:=t1.Nanosecond()
					 ns6 :=int64(ns)
					So(time.Unix(s,ns6).String(),ShouldEqual,t1.String())
				})

				Convey("从时间中解析出年月日 时分秒",func(){

					year,mon,day,hour,min,sec,nsec := 2019,time.January,9,22,11,18,893
					t1:=time.Date(year,mon,day,hour,min,sec,nsec,time.UTC)
					Convey("分两批导出，一次三个",func(){
						year1,mon1,day1:=t1.Date()
						So(year,ShouldEqual,year1)
						So(mon1,ShouldEqual,mon)
						So(day1,ShouldEqual,day)

						hour1,min1,sec1:=t1.Clock()
						So(hour1,ShouldEqual,hour)
						So(min1,ShouldEqual,min)
						So(sec1,ShouldEqual,sec)


					})

					Convey("单独导出",func(){
						year2:=t1.Year()
						mon2:=t1.Month()
						day2:=t1.Day()
						hour2:=t1.Hour()
						min2:=t1.Minute()
						sec2:=t1.Second()
						ns2:=t1.Nanosecond()
						So(year2,ShouldEqual,year)
						So(mon2,ShouldEqual,mon)
						So(day2,ShouldEqual,day)


						So(hour2,ShouldEqual,hour)
						So(min2,ShouldEqual,min)
						So(sec2,ShouldEqual,sec)
						So(ns2,ShouldEqual,nsec)


					})


				})
				Convey("从时间中导出一年的第几个星期" +
					"1月1号到1月3号可能属于上一年的最后一周，12月29号到12月31号可能属于下一年的第一周。" +
					"，一年的第几天，" +
					"是星期几",func(){
					t1:=time.Now()
					year,numweek:=t1.ISOWeek()
					So(year,ShouldEqual,2019)
					So(numweek,ShouldEqual,2)
				})



			})

			Convey("时间操作",func(){

				Convey("判断时间是否是零 也就是初始值",func(){
					t1:=time.Now()
					So(t1.IsZero(),ShouldBeFalse)
					t2,_:=time.Parse(time.RFC3339,"201939491234")
					So(t2.IsZero(),ShouldBeTrue)
				})

				Convey("把当前时间转换为本地时间，时间真正没变，形式变了" +
					"甚至可以转变为特定时区的时间",func(){
					t1:=time.Now()
					t2:=t1.UTC()
					t3:=t2.Local()

					loc:=time.FixedZone("df",3600)
					t4:=t3.In(loc)
					So(t4.String(),ShouldEqual,"2019")

					//So(t1.String(),ShouldEqual,t3.String())
					//So(t2.String(),ShouldEqual,t1.String())

				})

				Convey("时间比较",func(){
					t1:=time.Now()
					time.Sleep(time.Second)
					t2:=time.Now()
					Convey("比较时间是否相等",func(){
						So(t1.Equal(t2),ShouldBeFalse)


					})
					Convey("比较时间是否超前另一个",func(){
						So(t1.Before(t2),ShouldBeTrue)


					})
					Convey("比较时间是否落后另一个",func(){
						So(t2.After(t1),ShouldBeTrue)

					})


				})



			})




			Convey("时间转换",func(){
				Convey("转换为秒 纳秒",func(){


				})


			})

			Convey("",func(){


			})






		})


	})

}