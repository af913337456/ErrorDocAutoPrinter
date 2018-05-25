package core

import (
	"strings"
	"bufio"
	"fmt"
)

type EachLineMatchPrinter struct {

}

func NewDefaultEachLineMatchPrinter() *EachLineMatchPrinter {
	return &EachLineMatchPrinter{}
}

// 逐行判断的形式
func (p EachLineMatchPrinter) FindLines(
	printer *ErrorDocPrinter,reader *bufio.Reader,fileName,currentSuffix string,handleLine func(line string)) []string {
	var lines []string
	for {
		byt, _, err := reader.ReadLine()
		if err != nil {
			// 读完一个文件
			break
		}
		line := string(byt)
		// 排除注释
		if startWith(line,"//") {
			continue
		}
		if startWith(line,"/*") {
			continue
		}
		if startWith(line,"*") {
			continue
		}
		for _,value := range printer.TargetErrorFuncName {
			if strings.Contains(line,value) {
				// hit，准备生成
				handleLine(line)
				lines = append(lines,line)
			}
		}
	}
	return lines
}


func (p EachLineMatchPrinter) BuildACell(printer ErrorDocPrinter,columns,size int,prefixName,param string) string {
	return prefixName + param
}

func (p EachLineMatchPrinter) ResultLine(line string) {
	//fmt.Println(line)
}

func (p EachLineMatchPrinter) EndOfAFile(printer ErrorDocPrinter,aFileRetLines []string) {
	//fmt.Println(aFileRetLines)
}

func (p EachLineMatchPrinter) EndOfAllFile(printer ErrorDocPrinter,allRetLines []string) {
	fmt.Println(allRetLines)

}

