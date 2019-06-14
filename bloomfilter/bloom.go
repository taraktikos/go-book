package main

import (
	"github.com/spaolacci/murmur3"
	"hash"
	"hash/fnv"
)

type Interface interface {
	Add(item []byte)
	Test(item []byte) bool
}

type BloomFilter struct {
	bitSet    []bool        //the bloom filter bitSet
	k         uint          //number of hash values
	n         uint          //number of elements in the filter
	m         uint          //size of the bloom filter bitSet
	hashFuncs []hash.Hash64 //the hash functions
}

func (bf *BloomFilter) Add(item []byte) {
	hashes := bf.hashValues(item)
	for i := uint(0); i < bf.k; i++ {
		position := uint(hashes[i]) % bf.m
		bf.bitSet[position] = true
	}
	bf.n++
}

func (bf *BloomFilter) Test(item []byte) bool {
	hashes := bf.hashValues(item)
	for i := uint(0); i < bf.k; i++ {
		position := uint(hashes[i]) % bf.m
		if !bf.bitSet[position] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) hashValues(item []byte) (result []uint64) {
	for _, hashFunc := range bf.hashFuncs {
		_, err := hashFunc.Write(item)
		if err != nil {
			panic("can't calculate hash")
		}
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}
	return result
}

func New(size uint) *BloomFilter {
	return &BloomFilter{
		bitSet:    make([]bool, size),
		k:         3,
		n:         uint(0),
		m:         size,
		hashFuncs: []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()},
	}
}
