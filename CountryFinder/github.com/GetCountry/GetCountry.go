package main

import (
	"fmt"
	"os"
	"github.com/Geo-GoLang/geobed"
	"strconv"
	"bufio"
	"strings"
	"log"
)



func main() {
		
	var g = geobed.NewGeobed()
	var firstLine bool = true
	file, err := os.Open("c:/temp/636057301304931794_script.txt")	
		
    if err != nil {
        log.Fatal(err)
    }	
    defer file.Close()

	fileHandle, _ := os.Create("c:/temp/Part-3.txt") 
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		if firstLine == true {
			firstLine = false
			fmt.Fprintln(writer,  scanner.Text() + "\n\r")
			continue 
		}
		stringSlice := strings.Split(scanner.Text(), ",")
		lat, _ := strconv.ParseFloat(stringSlice[7], 64)
		long, _ := strconv.ParseFloat(stringSlice[8], 64)
		r :=  g.ReverseGeocode(lat, long)
		stringSlice[2] =  r.Country
		fmt.Fprintln(writer,  strings.Join(stringSlice, ",") +"\r")
		fmt.Printf(r.Country + ", ")
    }

	writer.Flush()
	
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
}

