package iohandler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type BatchProcessor func(input *string) (*string, error)

func ProcessStdinBatched(processor BatchProcessor, maxBytes int, splitString string) error {
	return ProcessStreamingBatches(os.Stdin, os.Stdout, processor, maxBytes, splitString)
}

func ProcessStreamingBatches(r io.Reader, w io.Writer, processor BatchProcessor, maxBytes int, splitString string) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(w)
	defer writer.Flush()

	var buffer bytes.Buffer

	flushBuffer := func() error {
		processed, err := processBuffer(&buffer, splitString, processor)
		buffer.Reset()

		if err != nil {
			return err
		}
		writer.Write(processed.Bytes())
		processed.Reset()
		return nil
	}

	for scanner.Scan() {
		line := scanner.Text()
		if buffer.Len()+len(line)+1 > maxBytes {
			if err := flushBuffer(); err != nil {
				return err
			}
		}

		buffer.WriteString(line)
		buffer.WriteString(splitString)
	}

	if err := flushBuffer(); err != nil {
		return err
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan error: %w", err)
	}

	return nil
}

func processBuffer(buffer *bytes.Buffer, splitString string, processor BatchProcessor) (*bytes.Buffer, error) {
	lines := strings.Split(buffer.String(), splitString)
	if len(lines) == 0 {
		return &bytes.Buffer{}, nil
	}
	var skippedIndices []int
	var skippedLines []string
	var processTextBuilder strings.Builder

	for i, line := range lines {
		if isASCII(line) {
			skippedIndices = append(skippedIndices, i)
			skippedLines = append(skippedLines, line)
		} else {
			processTextBuilder.WriteString(line)
			processTextBuilder.WriteString(splitString)
		}
	}

	processeText := processTextBuilder.String()
	output, err := processor(&processeText)

	if err != nil {
		return nil, fmt.Errorf("Failed processing the text: %w", err)
	}

	result := assembleWithSkippedLines(skippedIndices, skippedLines, output, splitString)

	return result, nil
}

func assembleWithSkippedLines(skippedIndices []int, skippedLines []string, processed *string, splitString string) *bytes.Buffer {
	var result bytes.Buffer
	processedLines := strings.Split(*processed, splitString)
	skippedIndex := 0
	processedIndex := 0
	totalLines := len(processedLines) + len(skippedLines)

	for i := range totalLines {
		if skippedIndex < len(skippedIndices) && i == skippedIndices[skippedIndex] {
			result.WriteString(skippedLines[skippedIndex])
			skippedIndex++
		} else {
			result.WriteString(processedLines[processedIndex])
			processedIndex++
		}
		result.WriteString(splitString)
	}

	return &result
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
