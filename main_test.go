package QraTest

import (
	"testing"
)

func TestOk(t *testing.T) {
	Ok(true)
	func(){
		defer func(){
			v:=recover()
			println(v)
		}()
		Ok(false)
	}()
}