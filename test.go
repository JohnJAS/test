package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type score struct {
	studentNum   int
	studentName  string
	mathScore    int
	englistScore int
	totalScore   int
}

func (s score) getTotalScore() int {
	return s.mathScore + s.englistScore
}

func Exec() {

	socreList := []score{
		{1001, "John", 85, 90, 0},
		{1002, "Mary", 80, 80, 0},
		{1003, "Mike", 75, 70, 0},
		{1004, "Rose", 85, 76, 0},
		{1005, "Ann ", 89, 62, 0},
		{1006, "Lee ", 55, 88, 0},
		{1007, "Tom ", 68, 74, 0},
	}

	fmt.Println("     学号     姓名    数学    英语    总分")

	for i, x := range socreList {
		i++
		fmt.Println(strconv.Itoa(i), "  ", strconv.Itoa(x.studentNum), "   ", x.studentName, "   ", strconv.Itoa(x.mathScore), "    ", strconv.Itoa(x.englistScore), "   ", strconv.Itoa(x.getTotalScore()))
	}
	var input string

	fmt.Println("Generate cvs file? (Y/N)")
	fmt.Scanln(&input)
	if input == "y" || input == "Y" || input == "yes" || input == "Yes" || input == "YES" {
		filename := "./mydata.csv"
		columns := make([][]string, 100)

		columns[0] = append(columns[0], "", "学号", "姓名", "数学", "英语", "总分")

		for i, x := range socreList {
			i++
			columns[i] = append(columns[i], strconv.Itoa(i), strconv.Itoa(x.studentNum), x.studentName, strconv.Itoa(x.mathScore), strconv.Itoa(x.englistScore), strconv.Itoa(x.getTotalScore()))
		}

		ExportCsv(filename, columns)
		fmt.Println("Success.")
	} else {
		fmt.Println("Exited.")
	}

}

func ExportCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) // 创建文件句柄
	if err != nil {
		log.Fatalf("Generate ["+filePath+"] failed,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}

func Fun() {
	fmt.Println("hahahahahahahahaha...")
}
