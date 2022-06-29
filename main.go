package main

import (
	"fmt"
	excelize "github.com/xuri/excelize/v2"
	"math/rand"
	"time"
)

const (
	add         = "+"
	sub         = "-"
	mul         = "×"
	div         = "÷"
	questionLen = 1600
	minNumber   = 0
	maxNumber   = 100
	timeFormat  = "2006-01-02 15:04:05"
)

var (
	operators         = []string{add, sub}
	operatorsMulOrDiv = []string{mul, div}
)

func main() {
	rand.Seed(time.Now().Unix())
	questionList := make([]string, 0, questionLen)
	qAndA := make([]string, 0, questionLen)
	for i := 0; i < questionLen; i++ {
		if rand.Intn(2) == 1 {
			questionList, qAndA = addOrSub(questionList, qAndA)
			continue
		}
		questionList, qAndA = mulOrDiv(questionList, qAndA)
	}
	// 打印题目
	//output(questionList, qAndA)
	// 生成文件
	outputExcel(questionList, "作业")
	outputExcel(qAndA, "答案")
}

func mulOrDiv(questionList, qAndA []string) ([]string, []string) {
	a, b := randNum9(), randNum9()
	answer := run(a, b, mul)
	questionList = append(questionList, fmt.Sprint(a, mul, b, "="))
	qAndA = append(qAndA, fmt.Sprint(answer))
	return questionList, qAndA
}

func addOrSub(questionList, qAndA []string) ([]string, []string) {
	firstNum, secondNum, lastNum := initNum()
	firstOper, secondOper := initOperator()
	stepOne := run(firstNum, secondNum, firstOper)
	if stepOne < minNumber || stepOne > maxNumber {
		return addOrSub(questionList, qAndA)
	}
	answer := run(stepOne, lastNum, secondOper)
	if answer < minNumber || answer > maxNumber {
		return addOrSub(questionList, qAndA)
	}
	questionList = append(questionList, fmt.Sprint(firstNum, firstOper, secondNum, secondOper, lastNum, "="))
	//qAndA = append(qAndA, fmt.Sprint(firstNum, firstOper, secondNum, secondOper, lastNum, "=", answer))
	qAndA = append(qAndA, fmt.Sprint(answer))
	return questionList, qAndA
}

func output(questionList, qAndA []string) {
	for k, v := range questionList {
		fmt.Println(k+1, "、", v)
	}
	fmt.Println("--------   答 案   --------")
	for k, v := range qAndA {
		fmt.Println(k+1, "、", v)
	}
}

func outputExcel(questionList []string, name string) {
	f := excelize.NewFile()
	for i := 0; i < len(questionList)/50; i++ {
		f.SetActiveSheet(f.NewSheet(fmt.Sprintf("Sheet%v", i+1)))
		var sp = questionList[i*50 : (i+1)*50]
		if len(sp) > 25 {
			f = fQList(sp[:25], f, fmt.Sprintf("Sheet%v", i+1), i)
		} else {
			f = fQList(sp, f, fmt.Sprintf("Sheet%v", i+1), i)
			break
		}
		f = sQList(sp[25:], f, fmt.Sprintf("Sheet%v", i+1), i)
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs(fmt.Sprintf("%s%s.xlsx", name, time.Now().Format(timeFormat))); err != nil {
		fmt.Println(err)
	}
}

func fQList(questionList []string, f *excelize.File, sheet string, i int) *excelize.File {
	f.SetCellValue(sheet, fmt.Sprintf("A%d", 1), fmt.Sprintf("第%d页", i+1))
	for k, v := range questionList {
		f.SetCellValue(sheet, fmt.Sprintf("B%d", k+2), fmt.Sprintf("%s", v))
	}
	return f
}

func sQList(questionList []string, f *excelize.File, sheet string, i int) *excelize.File {
	for k, v := range questionList {
		//f.SetCellValue(sheet, fmt.Sprintf("D%d", k+2), fmt.Sprintf("%d", (k+1)+i*25+25))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", k+2), fmt.Sprintf("%s", v))
	}
	return f
}

func questionOpt(f *excelize.File) {
	f.SetCellValue("Sheet1", "A1", fmt.Sprintf("%s", "序号"))
	f.SetCellValue("Sheet1", "B1", fmt.Sprintf("%s", "题目"))

}

func randNum() int {
	return rand.Intn(100)
}

func randNum9() int {
	return rand.Intn(9) + 1
}

func randomOperator() string {
	return operators[rand.Intn(len(operators))]
}

func initOperator() (string, string) {
	return randomOperator(), randomOperator()
}

func initNum() (int, int, int) {
	return randNum(), randNum(), randNum()
}

func run(a, b int, oper string) int {
	switch oper {
	case add:
		return a + b
	case mul:
		return a * b
	case div:
		return a / b
	}
	return a - b
}
