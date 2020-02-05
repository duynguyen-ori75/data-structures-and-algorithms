package main

import (
	"bloom"
	"fmt"
	"strconv"
)

func main() {
	filter := bloom.NewBloomFilter(10)
	filter.Add("Hello_World")
	filter.Add("aaabbbcccc")
	fmt.Printf("Is 'Hello_World' in the filter: %s\n", strconv.FormatBool(filter.PossiblyHave("Hello_World")))
	fmt.Printf("Is 'aaabbbcccc' in the filter: %s\n", strconv.FormatBool(filter.PossiblyHave("aaabbbcccc")))
	fmt.Printf("Is 'xxxuuuaa' in the filter:  %s\n", strconv.FormatBool(filter.PossiblyHave("xxxuuuaa")))
	fmt.Printf("Is 'random' in the filter:  %s\n", strconv.FormatBool(filter.PossiblyHave("random")))
}
