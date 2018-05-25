package core

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"strconv"
)

// todo NarkDown 风格

type MarkDownErrorDocPrinter struct {

}

func NewDefaultMarkDownErrorDocPrinter() *MarkDownErrorDocPrinter {
	return &MarkDownErrorDocPrinter{}
}

func (p MarkDownErrorDocPrinter) FindLines(
	printer *ErrorDocPrinter,reader *bufio.Reader,fileName,currentSuffix string,handleLine func(line string)) []string {
	// 正则匹配 todo
	var lines []string
	printer.currentLineNum = 0
	for {
		byt, _, err := reader.ReadLine()
		if err != nil {
			// 读完一个文件
			break
		}
		line := string(byt)
		// 排除注释
		printer.currentLineNum++
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

var tipsMap = make(map[string]int)
var diffMap = make(map[string]string)

var codeArr []int64

func quickSort(arr *[]int64,left,right int) {
	if arr == nil {
		return
	}
	if right == len(*arr) {
		right--
	}
	if left < 0 || left >= len(*arr) {
		return
	}
	hight := right
	low   := left
	base  := (*arr)[left]
	if low < hight {
		for ;low < hight; {
			for ;low < hight && base <= (*arr)[hight]; {
				hight--
				break
			}
			(*arr)[low] = (*arr)[hight]
			for ;low < hight && base >= (*arr)[low]; {
				low++
				break
			}
			(*arr)[hight] = (*arr)[low]
		}
		(*arr)[low] = base
		quickSort(arr,left,low-1)
		quickSort(arr,low+1,right)
	}
}

func (p MarkDownErrorDocPrinter) BuildACell(printer ErrorDocPrinter,columns,size int,prefixName,param string) string {
	/**
	| Name | Academy | score |
	| - | - | - |
	| Harry Potter | Gryffindor| 90 |
	| Hermione Granger | Gryffindor | 100 |
	| Draco Malfoy | Slytherin | 90 |
	*/
	if columns == 0 {
		code,err := strconv.ParseInt(param,10,64)
		if err == nil {
			codeArr = append(codeArr,code)
		}
		return "|" + param
	}
	count := tipsMap[param]
	if columns == 1 {
		// 保存提示列
		if count != 0 {
			count++
			diffMap[fmt.Sprintf("param: -- %s -- times:%d",param,count-1)] =
				fmt.Sprintf(" 与 %s 的第 %d 行提示重复",printer.currentFileName,printer.currentLineNum)
		}else{
			count = 1
		}
		tipsMap[param] = count
	}
	if columns == size - 1 {
		return "|" + param + "|"
	}
	// 找出提示一样，但是 code 不一样的
	return "|" + param
}

func (p MarkDownErrorDocPrinter) ResultLine(line string) {
	//fmt.Println(line)
}

func (p MarkDownErrorDocPrinter) EndOfAFile(printer ErrorDocPrinter,aFileRetLines []string) {

}

func (p MarkDownErrorDocPrinter) EndOfAllFile(printer ErrorDocPrinter,allRetLines []string) {
	//fmt.Println(retLines)
	var final []string
	header := "|"
	for _,name := range printer.ParamsColumnNames {
		header = header + name + "|"
	}
	size := len(printer.ParamsColumnNames)

	middle := ""
	for i:=0;i<size;i++ {
		middle = middle + "| - "
	}
	middle = middle + "|"
	final = append(final,header,middle)

	quickSort(&codeArr,0,len(codeArr))
	codeArrSize := len(codeArr)
	for i:=0; i<codeArrSize ;i++ {
		codeAtr := strconv.Itoa((int)(codeArr[i]))
		index := 0
		for _,line := range allRetLines {
			if strings.Contains(line,"|"+codeAtr+"|") {
				final = append(final,line)
				// 减去一个，减少循环次数
				//retLines = append(retLines[:index],retLines[index+1:]...)
				index--
				break
			}
			index++
		}
	}

	// 生成文件
	fileName := "errorInfo.md"
	file,err := os.Create(fileName)
	defer file.Close()
	if err!=nil {
		fmt.Println(err)
	}
	for _,line := range final {
		fmt.Println(line)
		file.WriteString(line+"\n")
	}
	fmt.Println("=====================================")
	for name,tip := range diffMap {
		fmt.Println(name + tip)
	}
	fmt.Println("=====================================")
}

