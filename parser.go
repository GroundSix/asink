package main

type Parser interface {
	Parse() map[string]string
}

