package service

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/trapajim/rest/api/config"
)

// ListMetaData defines the meta deta of a list
type ListMetaData struct {
	PL          int
	PTS         int
	CP          int
	Faction     string
	Detachments []config.Detachments
}

// ExractMetaData gets the meta data of an army list
func ExractMetaData(list string) ListMetaData {
	metaDataResult := ListMetaData{}
	extractListMeta(list, &metaDataResult)
	extractFaction(list, &metaDataResult)
	extractDetachments(list, &metaDataResult)
	return metaDataResult
}

func extractListMeta(list string, metaDataResult *ListMetaData) {
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
}

type factionCount struct {
	name  string
	count int
}

func extractFaction(list string, metaDataResult *ListMetaData) {
	factions := config.GetFactions()
	result := factionCount{count: 0}
	for _, faction := range factions {
		for _, sub := range faction {
			count := strings.Count(list, sub.Name)
			if count > result.count {
				result.count = count
				result.name = sub.Name
			}
		}
	}
	metaDataResult.Faction = result.name
}

func extractDetachments(list string, metaDataResult *ListMetaData) {
	detachments := config.GetDetachments()
	result := []config.Detachments{}
	for _, detachment := range detachments {
		count := strings.Count(list, detachment.Name)
		if count >= 1 {
			for i := 0; i < count; i++ {
				result = append(result, detachment)
			}

		}
	}
	metaDataResult.Detachments = result
}
