package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap []interface{}

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x)
}

func (h *Heap) Pop() interface{} {
	size := len(*h)
	if size <= 0 {
		return nil
	}
	element := (*h)[size-1] // Index into the slice and obtain the element.
	*h = (*h)[:size-1]      // Remove it from the stack by slicing it off.
	return element

}

func (h *Heap) IsIn(str string) bool {
	for _, elem := range *h {
		if elem == str {
			return true
		}
	}
	return false
}

func (h *Heap) IsEmpty() bool { return len(*h) == 0 }
func (h *Heap) Len() int      { return len(*h) }
func (h *Heap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type XY struct {
	X int
	Y int
}

func (xy *XY) Max() int {
	if xy.X > xy.Y {
		return xy.X
	}
	return xy.Y
}

func (xy *XY) Min() int {
	if xy.X < xy.Y {
		return xy.X
	}
	return xy.Y
}

func (xy1 *XY) Distance(xy2 XY) XY {
	return XY{xy1.X - xy2.X, xy1.Y - xy2.Y}
}

type XYZ struct {
	X int
	Y int
	Z int
}

func (xyz1 *XYZ) Distance(xyz2 XYZ) XYZ {
	return XYZ{xyz1.X - xyz2.X, xyz1.Y - xyz2.Y, xyz1.Z - xyz2.Z}
}

func EqualDistance(xyz1 XYZ, xyz2 XYZ) bool {

	if xyz1.X == xyz2.X && xyz1.Y == xyz2.Y && xyz1.Z == xyz2.Z {
		return true
	}
	return false
}

func (xyz1 *XYZ) Mult(xtz2 XYZ) XYZ {
	return XYZ{xyz1.X * xtz2.X, xyz1.Y * xtz2.Y, xyz1.Z * xtz2.Z}
}

func Add(xyz1 XYZ, xyz2 XYZ) XYZ {
	return XYZ{xyz1.X + xyz2.X, xyz1.Y + xyz2.Y, xyz1.Z + xyz2.Z}
}

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func AtoiRune(a rune) int {
	if a < '0' || a > '9' {
		panic("Cannot convert rune to int")
	}
	return int(a - '0')
}

func LoadFile(filepath string) *bufio.Scanner {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}
