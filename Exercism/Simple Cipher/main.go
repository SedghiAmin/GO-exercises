package main

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type shift int
type vigenere string
