package core

import (
	"bufio"
)

type IErrorDocPrinter interface {
	//HandleFileContent(printer ErrorDocPrinter,reader *bufio.Reader,fileName,currentSuffix string)
	FindLines(printer *ErrorDocPrinter,reader *bufio.Reader,fileName,currentSuffix string,handleLine func(line string)) []string
	BuildACell(printer ErrorDocPrinter,columns,size int,prefixName,param string) string
	ResultLine(line string)
	EndOfAFile(printer ErrorDocPrinter,aFileRetLines []string)
	EndOfAllFile(printer ErrorDocPrinter,allRetLines []string)
}

















