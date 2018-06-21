package main

// InputData contains either the byte slice of raw data or the file name and path to the saved data
type InputData struct {
	isFile   bool
	fileName string
	data     []byte
}
