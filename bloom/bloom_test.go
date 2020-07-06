package bloom

import "testing"

func TestBloom(t *testing.T) {
	filter := NewBloomFilter(10)
	filter.Add("Hello_World")
	filter.Add("aaabbbcccc")
	if !filter.PossiblyHave("Hello_World") || !filter.PossiblyHave("aaabbbcccc") {
		t.Error("The bloom filter should contains the value")
	}
	if filter.PossiblyHave("xxxuuuaa") || filter.PossiblyHave("random") {
		t.Error("The bloom filter should not contains the value")
	}
}