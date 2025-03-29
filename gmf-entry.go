package main

type GMFEntry strucy {}

func NewGMFEntry() GMFEntry {
	return GMFEntry{}
}

func (g GMFEntry) Applicable(e GitMetaEntry) bool {
	return e.Type == ""
}

func (g GMFEntry) Expand(e GitMetaEntry) ([]Cloneable, error) {
	return []Cloneable{
		URL: e.URL,
		Dir: e.Dir,
	}, nil
}