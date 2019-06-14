package main

import "testing"

func TestIfExists(t *testing.T) {
	bf := New(1024)
	bf.Add([]byte("Hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("sir"))
	bf.Add([]byte("io"))

	tt := []struct {
		item   []byte
		exists bool
	}{
		{item: []byte("Hello"), exists: true},
		{item: []byte("hello"), exists: false},
		{item: []byte("sir"), exists: true},
		{item: []byte("test"), exists: false},
		{item: []byte("hi"), exists: false},
	}

	for _, test := range tt {
		t.Run(string(test.item), func(t *testing.T) {
			if bf.Test(test.item) != test.exists {
				t.Errorf("Elements exists but not foubd")
			}
		})
	}
}
