package core


import (
	"path/filepath"
	"os"
	"strings"
	"bufio"
	"regexp"
	"fmt"
	"errors"
)

const PrinterEmptyColumnsWork = "--空缺--"

type ErrorDocPrinter struct {
	TargetFileSuffix    []string
	TargetErrorFuncName []string
	ParamsColumnNames   []string
	FilterFileName      []string
	ParamsSplitChar     string
	IPrinter IErrorDocPrinter
	retLines []string
	currentFileName string
	currentLineNum  int
}

func NewDefaultErrorDocPrinter(ip IErrorDocPrinter) *ErrorDocPrinter {
	printer := &ErrorDocPrinter{}
	if !NewConfiger().BindDefaultConfig(printer) {
		return nil
	}
	printer.IPrinter = ip
	return printer
	//&ErrorDocPrinter{
	//	TargetFileSuffix:[]string{".go"},
	//	TargetErrorFuncName:[]string{"util.GetCommonErr","util.GetErrWithTips"},
	//	FilterFileName:[]string{"errorPrinter","91porn","default_account"},
	//	ParamsColumnNames:[]string{" 错 误 码 "," 含 义 ","提 示"},
	//	ParamsSplitChar:",",
	//	IPrinter:ip,
	//}
}

// 逐行判断的形式
func (p ErrorDocPrinter) printErrorDoc(rootPath string) error {
	if p.TargetFileSuffix == nil || len(p.TargetFileSuffix) == 0{
		return errors.New("invalid TargetFileSuffix")
	}
	suffixMap := make(map[string]int,len(p.TargetFileSuffix))
	for _,value := range p.TargetFileSuffix {
		suffixMap[value] = 1
	}
	endOfFileError := errors.New("end of read")
	err := filepath.Walk(rootPath, func(fileName string, info os.FileInfo, err error) error {
		if info == nil {
			return endOfFileError
		}
		if info.IsDir() {
			return nil
		}
		for _,filter := range p.FilterFileName {
			if strings.Contains(fileName,filter) {
				return nil
			}
		}
		// 读取文件
		arr := strings.Split(fileName,".")
		arrLen := len(arr)
		if arrLen <= 1 {
			// 没有后缀的文件，下一位
			return nil
		}
		currentSuffix := "."+arr[arrLen-1]
		var aFileLines []string
		if suffixMap[currentSuffix] == 1 {
			// 包含
			file,err := os.Open(fileName)
			defer file.Close()
			if err != nil {
				return errors.New("读取文件错误 "+fileName+" --- "+err.Error()+"；请解决后重新覆盖生成")
			}
			reader := bufio.NewReader(file)
			p.currentFileName = fileName
			counter := 0
			p.IPrinter.FindLines(&p,reader,fileName,currentSuffix, func(line string) {
				relLines := p.readFuncParams(line,currentSuffix)
				aFileLines = append(aFileLines,relLines...)
				p.retLines = append(p.retLines,relLines...)
				counter++
			})
			if counter > 0 {
				// 最后的提示性输出
				p.IPrinter.EndOfAFile(p,aFileLines)
				fmt.Println(fmt.Sprintf("文件："+fileName+" 共找到：%d 项",counter))
				fmt.Println("==========================END A FILE==========================\n")
			}
			return nil
		}
		return nil
	})
	if err == nil || (err != nil && err.Error() == endOfFileError.Error()) {
		allSize := len(p.retLines)
		if allSize > 0{
			p.IPrinter.EndOfAllFile(p,p.retLines)
			fmt.Println(fmt.Sprintf("目录："+rootPath+" 共找到：%d 项",allSize))
		}
	}
	return err
}

func (p ErrorDocPrinter) readFuncParams(funcStr,suffix string) []string {
	// text := `return util.GetCommonErr(1,"wrong","错误") 455454 util.GetCommonErr("2222check order failed")`
	var retLines []string
	for _,funcName := range p.TargetErrorFuncName {
		reg    := regexp.MustCompile(funcName+`\([\s\S]+?\)`)
		retArr := reg.FindAllString(funcStr, -1)

		for _,  match := range retArr {
			match = strings.Replace(match,funcName+"(","",-1)
			match = strings.Replace(match,")","",-1)
			params := strings.Split(match,p.ParamsSplitChar) // 逗号分割参数

			ret := p.buildNewLine(params,suffix)
			retLines = append(retLines,ret)
		}
	}
	return retLines
}

func (p ErrorDocPrinter) buildNewLine(params []string,suffix string) string {
	ret := ""
	columns := 0
	size  := len(p.ParamsColumnNames)
	pLen  := len(params)
	limit := pLen
	if size > pLen {
		limit = size
	}
	for i:=0 ;i < limit ; i++ {
		v := PrinterEmptyColumnsWork
		if i < pLen {
			v = params[i]
		}
		// 去除 "
		v = strings.Replace(v,"\"","",-1)
		if i >= size {
			ret = ret + "--空缺：" + v
			continue
		}
		if v == "err.Error(" && suffix == ".go"{
			// go 语言的错误输出
			v = "这是一个错误输出的函数"
		}
		cell := p.IPrinter.BuildACell(p,columns,limit,p.ParamsColumnNames[columns],v)
		if cell != "" {
			ret = ret + cell
			columns ++
		}
	}
	p.IPrinter.ResultLine(ret)
	return ret
}

func startWith(str,sub string) bool {
	str = strings.Replace(str," ","",-1)
	str = strings.Replace(str,"	","",-1) // tab 符号
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

func substr(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}







