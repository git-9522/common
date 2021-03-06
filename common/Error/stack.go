package Error

import (
	"runtime"
	"casino_common/common/log"
)

func PrintPanicStack() {
	if x := recover(); x != nil {
		log.T("%v", x)
		for i := 0; i < 10; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				log.Error("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			}
		}
	}
}
