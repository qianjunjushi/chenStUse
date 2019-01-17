package byteschen

import (
	"bytes"
	"errors"
)
/*将字节索引转变为字符串索引*/
func ChangebyteIndex2Strindex(input string,index []int)(resindex []int,err error){

	for _,i:=range index{
		if i>len(input){
			err=errors.New("错误的索引范围")
			resindex=nil
			return
		}
		bytetemp:=[]rune(input[:i])
		//fmt.Println(string(bytetemp))
		resindex=append(resindex,len(bytetemp) )

	}
	return

}
/*将字符串索引转变为字节索引*/
func ChangestrIndex2byteindex(input string,index []int)(resindex []int,err error){

	for _,i:=range index{
		if i>len([]rune(input)){
			err=errors.New("错误的索引范围")
			resindex=nil
			return
		}
		runetemp:=([]rune(input))[:i]
		//fmt.Println(string(runetemp))
		resindex=append(resindex,len([]byte(string(runetemp))) )

	}
	return

}

/*获取每个子切片的索引位置 ，位置是字节索引*/
func GetChildSlicePos(input ,child []byte)(res []int ,err error){
	res=nil
	err=nil
	if len(child)==0{
		err=errors.New("子串不能为空")
		return
	}
	index:=bytes.Index(input,child)
	if index==-1{
		//err=errors.New("不存在该子串")
		return
	}else{
		res=append(res, index)
		input=input[index+len(child):]
	}
	for{
		index=bytes.Index(input,child)
		if index==-1{
			break

		}else{
			input=input[index+len(child):]
			res=append(res, index+len(child)+res[len(res)-1])
		}

	}
	return
}