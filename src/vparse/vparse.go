package vparse

import (
	"bufio"
	"io"
	"strings"
)

type Node struct {
	Type       string
	Properties map[string]string
	Children   []Node
}

func parseNode(scanner *bufio.Scanner, nodeType string) Node {
	node := Node{nodeType, make(map[string]string), make([]Node, 0)}

	var lastProperty string
	var expectColon bool = false

SC:
	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, " ") {
			value := line[1:]
			if expectColon {
				value = strings.SplitN(value, ":", 2)[1]
			}

			node.Properties[lastProperty] += value

			expectColon = false
			continue
		}

		if !strings.Contains(line, ":") {
			lastProperty = line
			expectColon = true
			continue
		}

		splitLine := strings.SplitN(line, ":", 2)

		switch splitLine[0] {

		case "END":
			break SC

		case "BEGIN":
			node.Children = append(node.Children, parseNode(scanner, splitLine[1]))
			break

		default:
			node.Properties[splitLine[0]] = splitLine[1]
			lastProperty = splitLine[0]

		}

	}

	return node
}

func Parse(reader io.Reader) []Node {
	scanner := bufio.NewScanner(reader)
	return parseNode(scanner, "").Children
}
