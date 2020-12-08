package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var BasePath string


func GetNumber(filename string) int {
	filenameSplit := strings.Split(filename, "_")
	number, err := strconv.Atoi(filenameSplit[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return number
}


func GetStepByNumber(number int) (*Step, error) {
	for n, step := range Steps {
		if step.Number == number {
			return &Steps[n], nil
		}
	}
	return nil, NewError("Not exists")
}


func Init() {
	var step *Step
	var err1 error

	BasePath = "migrates"

	files, err := ioutil.ReadDir(BasePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, file := range files {
		number := GetNumber(file.Name())

		step, err1 = GetStepByNumber(number)

		if err1 == nil {
			step.InsertStepFile(&StepFile{Name: file.Name()})
		} else {
			step := Step{}
			step.InsertStepFile(&StepFile{Name: file.Name()})
			Steps = append(Steps, step)
		}
	}
}