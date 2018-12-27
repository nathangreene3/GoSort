package main

import (
	"math/rand"
	"strings"
)

// person is a first and last name collection.
type person struct {
	first string
	last  string
}

// people is a collection of persons.
type people []*person

// length returns the number of people.
func (ppl people) length() int {
	return len(ppl)
}

// less returns the less-than comparison by last names. When equal, first names are compared.
func (ppl people) less(i, j int) bool {
	if ppl[i].last == ppl[j].last {
		return ppl[i].first < ppl[j].first
	}
	return ppl[i].last < ppl[j].last
}

func (ppl people) swap(i, j int) {
	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func (ppl people) at(i int) interface{} {
	return ppl[i]
}

func (ppl people) randomize() {
	rand.Shuffle(ppl.length(), func(i, j int) { ppl.swap(i, j) })
}

func (ppl people) String() string {
	if ppl.length() == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(ppl[0].String())
	for i := 1; i < ppl.length(); i++ {
		b.WriteString(" ")
		b.WriteString(ppl[i].String())
	}
	return b.String()
}

func (p *person) String() string {
	return strings.TrimSpace(p.first + " " + p.last)
}
