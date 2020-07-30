package QraTest

import (
	"fmt"
)

func Ok(ok bool,data ...interface{}){
	if ok==false{
		panic("fail "+fmt.Sprintln(data...))
	}
}