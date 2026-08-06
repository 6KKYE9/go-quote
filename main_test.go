package main

import (
	"math/rand"
	"testing"
)

func TestQuoteCount(t *testing.T) {
	if len(quotes) < 5 {
		t.Fatalf("内置名言应不少于 5 条, got %d", len(quotes))
	}
}

func TestPickDeterministic(t *testing.T) {
	// 相同种子应给出相同序列
	rng1 := rand.New(rand.NewSource(42))
	rng2 := rand.New(rand.NewSource(42))
	p1 := rng1.Perm(len(quotes))[:2]
	p2 := rng2.Perm(len(quotes))[:2]
	for i := range p1 {
		if p1[i] != p2[i] {
			t.Fatal("相同种子应产生相同序列")
		}
	}
}

func TestUniqueIndexes(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	idxs := rng.Perm(len(quotes))[:3]
	seen := make(map[int]bool)
	for _, i := range idxs {
		if seen[i] {
			t.Fatal("Perm 应返回不重复下标")
		}
		seen[i] = true
	}
}
