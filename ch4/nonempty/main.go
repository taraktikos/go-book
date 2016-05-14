package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func main() {
	data := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(data, 2))
}
