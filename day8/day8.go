package day8

import (
	. "adventofcode2022/tools"
	"fmt"
)

func isVisible(pos int, grid []int) bool {
	for i := 0; i < len(grid); i++ {
		if i == pos {
			return true
		} else if i < pos && grid[i] >= grid[pos] {
			i = pos
		} else if i > pos && grid[pos] <= grid[i] {
			return false
		}
	}
	return true
}

func countVisibleTrees(grid, gridT [][]int) int {
	visible := 0
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			if isVisible(j, grid[i]) || isVisible(i, gridT[j]) {
				visible++
			}
		}
	}
	edgeCount := len(grid)*2 + len(grid[0])*2 - 4
	return edgeCount + visible
}

func treeScore(pos int, grid []int) (left, right int) {
	for i := pos - 1; i >= 0; i-- {
		if grid[i] >= grid[pos] {
			left++
			break
		} else if grid[i] < grid[pos] {
			left++
		}
	}
	for i := pos + 1; i < len(grid); i++ {
		if grid[i] >= grid[pos] {
			right++
			break
		} else if grid[i] < grid[pos] {
			right++
		}
	}
	return left, right
}

func getHighestScenicScore(grid, gridT [][]int) (score int) {
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			left, right := treeScore(j, grid[i])
			up, down := treeScore(i, gridT[j])
			score = Max(score, left*right*up*down)
		}
	}
	return score
}

// gets a grid of trees and it's transposer
func getTreesGrid() (grid, gridT [][]int) {
	scanner := LoadFile("input/input8.txt")

	for row := 0; scanner.Scan(); row++ {
		srtIn := scanner.Text()
		if row == 0 {
			grid = make([][]int, len(srtIn))
			gridT = make([][]int, len(srtIn))
		}
		for i, t := range srtIn {
			grid[row] = append(grid[row], AtoiRune(t))
			gridT[i] = append(gridT[i], AtoiRune(t))
		}
	}
	return grid, gridT
}

func Ex1andEx2() {
	fmt.Println("Ex1: ", countVisibleTrees(getTreesGrid()))
	fmt.Println("Ex2: ", getHighestScenicScore(getTreesGrid()))
}
