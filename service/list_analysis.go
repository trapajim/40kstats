package service

import (
	"regexp"
	"strconv"
	"strings"
)

// ListMetaData defines the meta deta of a list
type ListMetaData struct {
	PL  int
	PTS int
	CP  int
}

// ExractMetaData gets the meta data of an army list
func ExractMetaData(list string) ListMetaData {
	metaDataResult := ListMetaData{}
	r := regexp.MustCompile("(?s)Total: \\[(.*?)\\]")
	extractedMetaString := r.FindStringSubmatch(list)
	metaData := strings.Split(extractedMetaString[1], ",")
	seperateNumberFromString := regexp.MustCompile("([0-9]+).*?([A-Z||a-z]+)")
	for _, element := range metaData {
		meta := seperateNumberFromString.FindStringSubmatch(element)

		intResult, err := strconv.Atoi(meta[1])
		if err != nil {
			continue
		}
		switch strings.ToLower(meta[2]) {
		case "pl":
			metaDataResult.PL = intResult
		case "pts":
			metaDataResult.PTS = intResult
		case "cp":
			metaDataResult.CP = intResult
		}
	}
	return metaDataResult
}
