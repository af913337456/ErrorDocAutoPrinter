package core

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"path/filepath"
)

type configer struct{

}

func NewConfiger () *configer {
	return &configer{}
}

func (s *configer) Load (filename string, v interface{}) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Load config json failed ===> filename : "+filename+" -- "+err.Error())
		return false
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil{
		fmt.Println("read json failed ===> filename : "+filename+" -- "+err.Error())
		return false
	}
	return true
}

func (s *configer) BindDefaultConfig(printer *ErrorDocPrinter) bool {
	configer := NewConfiger()
	isSuccess  := configer.Load(findConfigFile("DefaultConfig.json"), printer)
	if !isSuccess{
		return false
	}
	jsonBytes,_ := json.Marshal(printer)
	fmt.Println(string(jsonBytes))
	return true
}

func findConfigFile(fileName string) string {
	if _, err := os.Stat("./" + fileName); err == nil {

		fileName, _ = filepath.Abs("./" + fileName)

	} else if _, err := os.Stat("../" + fileName); err == nil {

		fileName, _ = filepath.Abs("../" + fileName)

	} else if _, err := os.Stat("../../" + fileName); err == nil {

		fileName, _ = filepath.Abs("../../" + fileName)

	}else if _, err := os.Stat("../../../" + fileName); err == nil {

		fileName, _ = filepath.Abs("../../../" + fileName)

	}else if _, err := os.Stat("../../../../" + fileName); err == nil {

		fileName, _ = filepath.Abs("../../../../" + fileName)

	} else if _, err := os.Stat("config/"+fileName); err == nil {

		fileName, _ = filepath.Abs("config/"+fileName)

	} else if _, err := os.Stat(fileName); err == nil {

		fileName, _ = filepath.Abs(fileName)

	}
	return fileName
}