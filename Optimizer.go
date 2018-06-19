package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
	"io"
)

type OBJOptimizer struct {
	f       *os.File
	Fstart  int
	Fend    int
	Fmiddle int
}
type OBJError struct {
	content string
}

func (e *OBJError) Error() string {
	return e.content
}
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(0)
	}
}
func NewOBJOptimizer(filename string) *OBJOptimizer {
	f, err := os.OpenFile("T.obj", os.O_APPEND|os.O_RDWR, 0660)
	checkErr(err)

	scanner := bufio.NewScanner(f)
	fstartline := 0
	fendline := 1
	currentline := 0
	for scanner.Scan() {
		currentline++
		str := scanner.Text()
		if strings.Contains(str, "f") {
			//F的开始行号
			if fstartline == 0 {
				fstartline = currentline
			}
			fendline = currentline
		}
	}
	mid := (fendline - fstartline) / 2
	mid = int(math.Floor(float64(mid + fstartline)))
	//defer f.Close()
	f.Seek(0,0)
	return &OBJOptimizer{f: f, Fstart: fstartline, Fend: fendline, Fmiddle: mid}
}
//ExtractFace is extract front face through readlines from mid to end.
func (e *OBJOptimizer) ExtractFace() {
	var data string
	OriginalBytes := bytes.NewBufferString(data)
	dataWriter := bufio.NewWriter(OriginalBytes)

	//fi, err := os.Open("T.obj")
	scanner := bufio.NewScanner(e.f)
	fmt.Println(scanner)

	lines := 1
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if lines < e.Fstart || lines >= e.Fmiddle {
			dataWriter.WriteString(scanner.Text()+"\n")
		}
		lines++


	}
	dataWriter.Flush()
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
	//Write to New obj file.
	fl,err:=os.OpenFile("T.bak.obj",os.O_CREATE|os.O_RDWR,0666)
	checkErr(err)
	dataReader:=bufio.NewReader(OriginalBytes)
	i,err:=io.Copy(fl,dataReader)
	fmt.Println("Write bytes:",i)
	checkErr(err)
	defer fl.Close()

}

func main() {
	fmt.Println(os.Getwd())
	newOpt := NewOBJOptimizer("T.obj")
	newOpt.ExtractFace()
	fmt.Printf("%v   %T", newOpt, newOpt)
}
