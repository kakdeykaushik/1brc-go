package main

import (
	"fmt"
	"math"
	"sort"
)

type Entry struct {
	min   float64
	max   float64
	sum   float64
	count float64
}

type Store struct {
	mapp map[string]Entry
}

func NewStore() *Store {
	return &Store{make(map[string]Entry)}
}

func (s *Store) Add(city string, temperature float64) {
	e, ok := s.mapp[city]
	if !ok {
		e = Entry{min: 100, max: -100, sum: 0, count: 0}
	}

	if temperature < e.min {
		e.min = temperature
	}
	if temperature > e.max {
		e.max = temperature
	}

	e.count++
	e.sum += temperature

	s.mapp[city] = e
}

func (s *Store) String() string {

	var keys []string
	for key := range s.mapp {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	res := "{"
	sep := ", "
	for _, k := range keys {
		v := s.mapp[k]
		res = res + fmt.Sprintf("%s=%.1f/%.1f/%.1f", k, v.min, round(v.sum/v.count), v.max) + sep
	}

	res = res[:len(res)-2]

	res += "}"
	return res
}

func round(value float64) float64 {
	return math.Ceil(value*10) / 10
}
