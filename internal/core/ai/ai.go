package ai

import "math/rand"

type AI struct{}

func (a AI) GetAttacks(list, vulnerabilities []string) []string {
	records := make([]string, 0)

	// logic goes here (now it's random)
	for _, item := range list {
		if item == "injection" || item == "payload" {
			records = append(records, item)

			continue
		}

		if rand.Intn(10) > 8 {
			records = append(records, item)
		}
	}

	return records
}
