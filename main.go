package main

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
)

func readDimensionMatrix(reader io.Reader) (uint32, error) {
	var dimensionMatrix uint32

	_, err := fmt.Fscan(reader, &dimensionMatrix)
	if err != nil {
		return 0, err
	}

	return dimensionMatrix, nil
}

func CanMatrixSortWithCycles(reader io.Reader) (bool, error) {
	dimensionMatrix, err := readDimensionMatrix(reader)
	if err != nil {
		return false, err
	}

	capContainers := make([]uint64, dimensionMatrix)
	countBallsByColor := make([]uint64, dimensionMatrix)

	for rowIdx := uint32(0); rowIdx < dimensionMatrix; rowIdx++ {
		for colIdx := uint32(0); colIdx < dimensionMatrix; colIdx++ {
			var curValue uint64

			_, err := fmt.Fscan(reader, &curValue)
			if err != nil {
				return false, err
			}

			capContainers[rowIdx] += curValue
			countBallsByColor[colIdx] += curValue
		}
	}

	var colIdx uint32

	for rowIdx := uint32(0); rowIdx < dimensionMatrix; rowIdx++ {
		for colIdx = rowIdx; colIdx < dimensionMatrix; colIdx++ {
			if capContainers[rowIdx] == countBallsByColor[colIdx] {
				countBallsByColor[rowIdx], countBallsByColor[colIdx] = countBallsByColor[colIdx], countBallsByColor[rowIdx]

				break
			}
		}

		if colIdx == dimensionMatrix {
			return false, nil
		}
	}

	return true, nil
}

func CanMatrixSortWithMap(reader io.Reader) (bool, error) {
	dimensionMatrix, err := readDimensionMatrix(reader)
	if err != nil {
		return false, err
	}

	capContainers := make(map[uint64]uint32, dimensionMatrix)
	countBallsByColor := make([]uint64, dimensionMatrix)

	for rowIdx := uint32(0); rowIdx < dimensionMatrix; rowIdx++ {
		capCurContainer := uint64(0)

		for colIdx := uint32(0); colIdx < dimensionMatrix; colIdx++ {
			var curValue uint64

			_, err := fmt.Fscan(reader, &curValue)
			if err != nil {
				return false, err
			}

			capCurContainer += curValue
			countBallsByColor[colIdx] += curValue
		}

		countContainersWithSameCap, ok := capContainers[capCurContainer]
		if ok {
			capContainers[capCurContainer] = countContainersWithSameCap + 1
		} else {
			capContainers[capCurContainer] = 1
		}
	}

	for _, curCountBallsByColor := range countBallsByColor {
		countContainers, ok := capContainers[curCountBallsByColor]
		if !ok || countContainers == 0 {
			return false, nil
		}

		if ok {
			capContainers[curCountBallsByColor] = countContainers - 1
		}
	}

	return true, nil
}

func CanMatrixSortWithSort(reader io.Reader) (bool, error) {
	dimensionMatrix, err := readDimensionMatrix(reader)
	if err != nil {
		return false, err
	}

	capContainers := make([]uint64, dimensionMatrix)
	countBallsByColor := make([]uint64, dimensionMatrix)

	for rowIdx := uint32(0); rowIdx < dimensionMatrix; rowIdx++ {
		for colIdx := uint32(0); colIdx < dimensionMatrix; colIdx++ {
			var curValue uint64

			_, err := fmt.Fscan(reader, &curValue)
			if err != nil {
				return false, err
			}

			capContainers[rowIdx] += curValue
			countBallsByColor[colIdx] += curValue
		}
	}

	slices.SortFunc(capContainers, cmp.Compare[uint64])
	slices.SortFunc(countBallsByColor, cmp.Compare[uint64])

	for idx, curCap := range capContainers {
		if curCap != countBallsByColor[idx] {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	result, err := CanMatrixSortWithCycles(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}

	if result {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
