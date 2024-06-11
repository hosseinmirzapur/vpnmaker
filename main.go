package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	// read file data from base.json
	fileData, err := os.ReadFile("base.json")
	if err != nil {
		log.Fatalln(err)
	}

	// marshal file data into the BaseConfig struct
	var bc BaseConfig
	err = json.Unmarshal(fileData, &bc)
	if err != nil {
		log.Fatalln(err)
	}

	// set IPs and Ports from the result.csv file
	setIPsAndPorts(&bc)

	// populate BaseConfig with values from warp1.txt and warp2.txt
	populateBC(&bc)

	// make the created.json file
	fileBytes, err := json.MarshalIndent(bc, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("created.json", fileBytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// print out a QR code of the config inside terminal
	// file, err := os.Create("created.png")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	qrc, err := qrcode.New(string(fileBytes))
	if err != nil {
		log.Fatalln(err)
	}

	w, err := standard.New("qrcode.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	if err = qrc.Save(w); err != nil {
		log.Fatalln(err)
	}
}

func setIPsAndPorts(bc *BaseConfig) {
	resultFile, err := os.OpenFile("result.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer resultFile.Close()

	results := []*Result{}
	err = gocsv.UnmarshalFile(resultFile, &results)
	if err != nil {
		log.Fatalln(err)
	}

	// generate 2 random different numbers between 0 and 6
	var i, j int
	for i == j {
		i = rand.IntN(6)
		j = rand.IntN(6)
	}

	// first outbound
	port, _ := strconv.Atoi(strings.Split(results[i].IpPort, ":")[1])
	bc.Outbounds[0].ServerPort = port
	bc.Outbounds[0].Server = strings.Split(results[i].IpPort, ":")[0]

	// second outbound
	port, _ = strconv.Atoi(strings.Split(results[j].IpPort, ":")[1])
	bc.Outbounds[1].ServerPort = port
	bc.Outbounds[1].Server = strings.Split(results[j].IpPort, ":")[0]
}

func populateBC(bc *BaseConfig) {
	setupWarpAcount(bc, "warp1.txt", 0)
	setupWarpAcount(bc, "warp2.txt", 1)
}

func setupWarpAcount(bc *BaseConfig, warpAccountPath string, outboundsIndex int) {
	warpAccount, err := parseTxtFile(warpAccountPath)
	if err != nil {
		log.Fatalln(err)
	}
	bc.Outbounds[outboundsIndex].LocalAddress[1] = warpAccount.IPV6
	bc.Outbounds[outboundsIndex].PrivateKey = warpAccount.PrivateKey
	reservedArray := []int{}
	for _, r := range warpAccount.Reserved {
		num, _ := strconv.Atoi(r)
		reservedArray = append(reservedArray, num)
	}
	bc.Outbounds[outboundsIndex].Reserved = reservedArray

}

func parseTxtFile(filename string) (*WarpAccount, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines) // Split by lines

	wa := WarpAccount{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		switch {
		case line == "":
			continue
		case wa.IPV6 == "":
			wa.IPV6 = line[1 : len(line)-1]
		case wa.PrivateKey == "":
			lineToBeAdded := strings.SplitN(line, ":", 2)[1]
			wa.PrivateKey = lineToBeAdded[1 : len(lineToBeAdded)-1]
		case len(wa.Reserved) == 0:
			lineToBeAdded := strings.SplitN(line, ":", 2)[1]
			wa.Reserved = strings.Split(lineToBeAdded[1:len(lineToBeAdded)-1], ",")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &wa, nil
}
