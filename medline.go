package guru

import (
	"bufio"
	"io"
	"strings"
)

const (
	PMID = "PMID"
	TI   = "TI"
	AB   = "AB"
	MH   = "MH"
)

type MedlineDocument struct {
	PMID string
	TI   string
	AB   string
	MH   []string
}

type MedlineDocuments []MedlineDocument

func UnmarshallMedline(r io.Reader) MedlineDocuments {
	s := bufio.NewScanner(r)
	var (
		item    string
		content []string
		docs    MedlineDocuments
		doc     MedlineDocument
	)
	for s.Scan() {
		line := s.Text()
		if line == "" && len(doc.PMID) > 0 {
			docs = append(docs, doc)
			doc = MedlineDocument{}
		} else if len(line) < 6 {
			continue
		} else if line[:6] == "      " {
			content = append(content, strings.TrimSpace(line))
		} else {
			switch item {
			case PMID:
				doc.PMID = strings.Join(content, " ")
			case TI:
				doc.TI = strings.Join(content, " ")
			case AB:
				doc.AB = strings.Join(content, " ")
			case MH:
				doc.MH = content
			}
			pair := strings.Split(line, "-")
			item = strings.TrimSpace(pair[0])
			content = []string{strings.TrimSpace(pair[1])}
		}
	}
	docs = append(docs, doc)
	return docs
}
