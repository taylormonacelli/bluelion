package bluelion

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
)

type Config struct {
	Verbose        bool
	LogFormat      string
	InputFilePath  string
	OutputFilePath string
}

func Main(config Config) int {
	slog.Debug("bluelion", "test", true)

	err := dowork(config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 1
	}
	return 0
}

func dowork(config Config) error {
	inputFilePath := config.InputFilePath
	outputFilePath := config.OutputFilePath

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("error opening the input file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var blocks []string
	var currentBlock []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			currentBlock = append(currentBlock, line)
		} else {
			if len(currentBlock) > 0 {
				blocks = append(blocks, strings.Join(currentBlock, "\n"))
				currentBlock = nil
			}
		}
	}

	if len(currentBlock) > 0 {
		blocks = append(blocks, strings.Join(currentBlock, "\n"))
	}

	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		sort.Strings(lines)
		blocks[i] = strings.Join(lines, "\n")
	}

	sortedResult := strings.Join(blocks, "\n\n")

	sortedResult = strings.TrimRight(sortedResult, "\n") + "\n"

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating the output file: %v", err)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(sortedResult)
	if err != nil {
		return fmt.Errorf("error writing the sorted data to the output file: %v", err)
	}

	return nil
}
