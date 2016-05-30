package main

import (
	"fmt"
	"io"
	"log"
	"time"
	"os"
	"bufio"
	"strings"
)

func main() {
	start := time.Now()
	//graph := parseFromFile()
	graph := map[string][]string {
		"A": {"B", "C", "E"},
		"B": {"A", "C", "D"},
		"C": {"D"},
		"D": {"C"},
		"E": {"F", "D"},
		"F": {"C"},
	}
	elapsed := time.Since(start)
	fmt.Print("Parsing time: ")
	fmt.Println(elapsed)

	fmt.Print("Graph size: ")
	fmt.Println(len(graph))

	start = time.Now()
	//result := bfs(graph, "William_Shakespeare", "Adolf_Hitler", [][]string{})
	result := bfs(graph, "A", "D")
	elapsed = time.Since(start)
	fmt.Print("Searching time: ")
	fmt.Println(elapsed)

	for _, r := range result {
		if len(r) < 10 {
			fmt.Println(r)
		}
	}
}

func bfs(graph map[string][]string, start string, end string) [][]string {
	result := [][]string{}
	q := [][]string {{start}}
	for len(q) > 0 {
		tmpPath := q[0]
		q = q[1:]
		lastNode := tmpPath[len(tmpPath)-1]
		if lastNode == end {
			result = append(result, tmpPath)
		}
		if edges, ok := graph[lastNode]; ok {
			for _, node := range edges {
				if !stringInSlice(node, tmpPath) {
					newPath := append(tmpPath, node)
					q = append(q, newPath)
				}
			}
		}
	}
	return result
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func parseFromFile() map[string][]string {
	var graph = make(map[string][]string)

	start := time.Now()
	inputFilePath := "out/links.txt"
	inputFile, err := os.Open(inputFilePath)
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//var allLinks = make(map[string]bool)
	reader := bufio.NewReader(inputFile)
	i := 1
	line, err := reader.ReadString('\n')
	for err == nil {
		arr := strings.Split(line, " - ")
		graph[arr[0]] = append(graph[arr[0]], strings.Split(arr[1], ", ")...)
		//allLinks[arr[0]] = true
		//for _, link := range strings.Split(arr[1], ", ") {
		//	allLinks[link] = true
		//}
		if i % 200000 == 0 {
			//break
		}
		i++
		line, err = reader.ReadString('\n')
	}
	if err != io.EOF && err != nil {
		log.Fatal(err)
	}
	//fmt.Println(graph)
	elapsed := time.Since(start)
	fmt.Print("Time: ")
	fmt.Println(elapsed)
	fmt.Print("Count ")
	fmt.Println(i)
	return graph
}
