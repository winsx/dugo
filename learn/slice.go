package learn

import "fmt"

func TrySlice() {
	var slice = make([]int, 20)
	sliceTmp := []int{1, 2, 3, 4, 5}
	for i, v := range sliceTmp {
		slice[i] = v
	}
	slice = append(slice, sliceTmp...)
	fmt.Printf("%v len:%d cap:%d \n", slice, len(slice), cap(slice))
}

func TrySliceRemove() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("orig:%v len:%d, cap:%d, removed:%v\n", s, len(s), cap(s), remove(s, 2))
}

func TrySliceRemove2() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("orig:%v len:%d, cap:%d, removed:%v\n", s, len(s), cap(s), remove2(s, 2))
}

func TrySliceNonEmpty() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)           // `["one", "", "three"]
	fmt.Printf("%q\n", nonEmpty(data)) // `["one", "three", "three"]
}

func TrySliceNonEmpty2() {
	data := []string{"one", "", "three", "", "", "", ""}
	fmt.Printf("%q\n", data)            // `["one", "", "three"]
	data2 := nonEmpty2(data)
	fmt.Printf("%q len:%d cap:%d\n", data2, len(data2), cap(data2)) // `["one", "three", "three"]
}

func remove(slice []int, i int) []int {
	if i < len(slice) {
		copy(slice[i:], slice[i+1:])
		return slice[:len(slice) - 1]
	}
	return slice
}

func remove2(slice []int, i int) []int {
	length := len(slice)
	if i < length {
		slice[i] = slice[length - 1]
		return slice[:length - 1]
	}
	return slice
}

func nonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings
}

func nonEmpty2(strings []string) []string {
	out := strings[:0] // zero-Length slice of original
	fmt.Printf("out len:%d cap:%d ", len(out), cap(out))
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
