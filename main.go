package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var m []string
var size int
var xInit int
var yInit int
var bigSize int
var bigX int
var bigY int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	fscanner.Scan()
	NbLine, err := strconv.Atoi(fscanner.Text())
	check(err)
	fscanner.Scan()
	NbCol := len(fscanner.Text())
	check(err)

	m = make([]string, NbLine)

	for index := 0; index < NbLine; index++ {
		m[index] = fscanner.Text()
		if !fscanner.Scan() {
			break
		}
	}
	myBsq(NbCol, NbLine)
}

func myBsq(NbCol int, NbLine int) {

	bigSize = 0

	for y := 0; y < NbLine; y++ {

		for x := 0; x < NbCol; x++ {
			myAlgo(y, x, NbLine, NbCol)
			checkSizes()
		}
	}

	putX()

	for _, line := range m {
		fmt.Println(line)
	}
}

func myAlgo(y int, x int, NbLine int, NbCol int) int {

	size = 1
	xInit = x
	yInit = y

	if m[y][x] == 'o' {
		return 0
	}

	for y+1 < NbLine && x+1 < NbCol && m[y+1][xInit] == '.' && m[yInit][x+1] == '.' && checkColumn(yInit, x+1, size+1) == 1 && checkLine(y+1, xInit, size+1) == 1 {
		x++
		y++
		size++
	}

	return size
}

func checkColumn(yInit, x, size int) int {

	for index := 0; index < size; index++ {

		if m[yInit][x] == 'o' {
			return 0
		}

		yInit++
	}

	return 1
}

func checkLine(y, xInit, size int) int {

	for index := 0; index < size; index++ {

		if m[y][xInit] == 'o' {
			return 0
		}

		xInit++
	}

	return 1
}

func checkSizes() {

	if size > bigSize {
		bigSize = size
		bigX = xInit
		bigY = yInit
	}
}

func putX() {

	bigXCopy := bigX
	for y := 0; y < bigSize; y++ {

		bigX = bigXCopy

		for i := 0; i < bigSize; i++ {
			prev := m[bigY]
			lol := bigX + 1
			s := fmt.Sprintf("%s%c%s", prev[0:bigX], 'x', prev[lol:])
			m[bigY] = s
			bigX++
		}

		bigY++
	}
}
