package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var validCandidates []string
	subjectFrequencyMap := map[rune]int{}
	lowercaseSubject := strings.ToLower(subject)

	for _, character := range lowercaseSubject {
		subjectFrequencyMap[character]++
	}

	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}

		lowercaseCandidate := strings.ToLower(candidate)
		if lowercaseSubject == lowercaseCandidate {
			continue
		}

		candidateFrequencyMap := map[rune]int{}
		for _, character := range lowercaseCandidate {
			candidateFrequencyMap[character]++
		}

		if reflect.DeepEqual(subjectFrequencyMap, candidateFrequencyMap) {
			validCandidates = append(validCandidates, candidate)
		}
	}
	return validCandidates
}
