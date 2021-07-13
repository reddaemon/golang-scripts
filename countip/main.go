package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
Программа парсит лог и считает кол-во запросов с уникальных адресов
 */
func main() {
	f, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("cannot open file: %s", err)
	}
	defer f.Close()

	ipmap := make(map[string]string)
	counter := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ipaddressReg := regexp.MustCompile(`([0-9]{1,3}\.){3}[0-9]{1,3}`)
		ipaddressData := ipaddressReg.FindAllString(line, -1)[0]
		ipmap[ipaddressData] = line
		counter[ipaddressData]++
	}

	for k, v := range counter {
		fmt.Println(k, v)
	}
}
