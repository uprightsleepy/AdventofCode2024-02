package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	locationIds, historianList, err := parseInputFile("values.txt")
	if err != nil {
		fmt.Println("an error occurred while parsing the input file")
	}

	sort.Ints(locationIds)
	sort.Ints(historianList)

	frequencyMap := make(map[int]int)
	for _, num := range historianList {
		frequencyMap[num]++
	}

	resultMap := make(map[int]int)
	for _, value := range locationIds {
		resultMap[value] = frequencyMap[value]
	}

	total := 0
	for key, value := range resultMap {
		total += key * value
	}

	fmt.Printf("The total similarity score is: %d\n", total)
}

func parseInputFile(filepath string) (locationIds []int, historianList []int, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("an error occurred while parsing the file {%s}: %v", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("malformed file received")
		}

		leftVal, _ := strconv.Atoi(fields[0])
		rightVal, _ := strconv.Atoi(fields[1])

		locationIds = append(locationIds, leftVal)
		historianList = append(historianList, rightVal)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("an error occurred while reading the file: %v", err)
	}

	return locationIds, historianList, nil
}
