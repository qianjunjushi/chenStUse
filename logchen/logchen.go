package logchen

import (
	"bytes"
	"fmt"
	"log"
)
/*一般用new 创建一个日志 关联一个writer 然后用printlin 往里写东西*/
func initloger(){
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Llongfile)
	logger.Print("Hello, log file!")
	fmt.Print(&buf)

}