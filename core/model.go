package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

var Steps []Step

type StepFile struct {
	Name string
	MD5 string
}

type Step struct {
	Up StepFile
	Down StepFile
	Number int
}

func (s *StepFile)IsValid()bool {
	return s.Name != "" && s.MD5 != ""
}

func (s *StepFile)CaculateMD5() {
	FullPath := BasePath + "/" + s.Name
	data, err := ioutil.ReadFile(FullPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	s.MD5 = hex.EncodeToString(md5Ctx.Sum(nil))
}

func (s *Step)InsertStepFile(stepFile *StepFile) {
	s.Number = GetNumber(stepFile.Name)

	filenameSplit := strings.Split(stepFile.Name, "_")
	arrow := filenameSplit[1]

	stepFile.CaculateMD5()

	if arrow == "up" {
		s.Up = *stepFile
	}else if arrow == "down" {
		s.Down = *stepFile
	}
}

func GetMaxNumber() int {
	if len(Steps) == 0 {
		return 0
	}

	max := Steps[0].Number

	for _, step := range Steps {
		if step.Number > max {
			max = step.Number
		}
	}

	return max
}