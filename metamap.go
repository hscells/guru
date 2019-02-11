package guru

import "github.com/hscells/metawrap"

func MetaMapConceptRatio(concepts1 []string, concepts2 []string) (float64, float64, int, int) {
	n := 0.0
	for _, c1 := range concepts1 {
		for _, c2 := range concepts2 {
			if c1 == c2 {
				n++
			}
		}
	}
	return n, n / float64(len(concepts1)), len(concepts1), len(concepts2)
}

func MetaMapCUIs(text string, client metawrap.HTTPClient) (c []string, err error) {
	candidates, err := client.Candidates(text)
	if err != nil {
		return
	}
	seen := make(map[string]bool)
	for _, candidate := range candidates {
		if _, ok := seen[candidate.CandidateCUI]; !ok {
			c = append(c, candidate.CandidateCUI)
			seen[candidate.CandidateCUI] = true
		}
	}
	return
}
