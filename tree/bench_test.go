package tree_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gaissmai/go-inet/internal"
	"github.com/gaissmai/go-inet/tree"
)

func BenchmarkTreeInsert(b *testing.B) {
	bench := []int{1000, 10000, 100000, 1000000}

	for _, n := range bench {
		bs := internal.GenBlockMixed(n)
		is := make([]tree.Item, len(bs))
		for i := range bs {
			is[i] = tree.Item{bs[i], nil, nil}
		}

		b.Run(fmt.Sprintf("%7d", n), func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				t := tree.New()
				if err := t.Insert(is...); err != nil {
					b.Errorf("item is duplicate: %s", err)
				}
			}
		})

	}
}

func BenchmarkContainsTree(b *testing.B) {
	bench := []int{1000, 10000, 100000, 1000000}

	for _, n := range bench {
		bs := internal.GenBlockMixed(n)
		is := make([]tree.Item, len(bs))
		for i := range bs {
			is[i] = tree.Item{bs[i], nil, nil}
		}

		t := tree.New()
		_ = t.Insert(is...)

		vx := is[rand.Intn(len(is))]
		b.Run(fmt.Sprintf("%7d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				t.Contains(vx)
			}
		})

	}
}

func BenchmarkLookupTree(b *testing.B) {
	bench := []int{1000, 10000, 100000, 1000000}

	for _, n := range bench {
		bs := internal.GenBlockMixed(n)
		is := make([]tree.Item, len(bs))
		for i := range bs {
			is[i] = tree.Item{bs[i], nil, nil}
		}

		t := tree.New()
		_ = t.Insert(is...)

		vx := is[rand.Intn(len(is))]
		b.Run(fmt.Sprintf("%7d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				t.Lookup(vx)
			}
		})

	}
}

func BenchmarkWalkTree(b *testing.B) {
	bench := []int{1000, 10000, 100000}

	for _, n := range bench {
		bs := internal.GenBlockMixed(n)
		is := make([]tree.Item, len(bs))
		for i := range bs {
			is[i] = tree.Item{bs[i], nil, nil}
		}

		t := tree.New()
		_ = t.Insert(is...)

		var walkFn tree.WalkFunc = func(n *tree.Node, l int) error { return nil }

		b.Run(fmt.Sprintf("%7d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = t.Walk(walkFn)
			}
		})

	}
}

func BenchmarkTreeRemoveItem(b *testing.B) {
	bench := []int{1000, 10000, 100000}

	for _, n := range bench {
		bs := internal.GenBlockMixed(n)
		is := make([]tree.Item, len(bs))
		for i := range bs {
			is[i] = tree.Item{bs[i], nil, nil}
		}

		t := tree.New()
		_ = t.Insert(is...)

		vx := is[rand.Intn(len(is))]
		b.Run(fmt.Sprintf("%7d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = t.Remove(vx)
			}
		})

	}
}

func BenchmarkTreeRemoveBranch(b *testing.B) {
	bench := []int{1000, 10000, 100000}

	for _, n := range bench {
		bs := internal.GenBlockMixed(n)
		is := make([]tree.Item, len(bs))
		for i := range bs {
			is[i] = tree.Item{bs[i], nil, nil}
		}

		t := tree.New()
		_ = t.Insert(is...)

		vx := is[rand.Intn(len(is))]
		b.Run(fmt.Sprintf("%7d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = t.RemoveBranch(vx)
			}
		})

	}
}
