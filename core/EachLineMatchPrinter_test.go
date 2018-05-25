package core

import (
	"testing"
	"fmt"
	"strings"
)


func TestDocPrinter(t *testing.T) {
	p := NewDefaultErrorDocPrinter(NewDefaultMarkDownErrorDocPrinter())
	if p == nil {
		return
	}
	fmt.Println(p.printErrorDoc("../../errorDocPrinter"))
}

//	return util.GetCommonErr("只有管理员才能修改评论")
func TestStartWith(t *testing.T) {
	if startWith2(`		//	return util.GetCommonErr("只有管理员才能修改评论")`,"//") {
		fmt.Println("yes")
	}
}

func startWith2(str,sub string) bool {
	str = strings.Replace(str," ","",-1)
	str = strings.Replace(str,"	","",-1)
	l := strings.Count(sub,"")
	temp := substr(str,0,l)
	if strings.Contains(temp,sub) {
		if sub == " " {
			if !strings.Contains(str,"#") {
				return false
			}
		}
		return true
	}
	return false
}







