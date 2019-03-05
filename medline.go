package guru

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	PMID = "PMID"
	TI   = "TI"
	AB   = "AB"
	MH   = "MH"
	PT   = "PT"
)

type MedlineDocument struct {
	PMID string   // PubMed ID
	TI   string   // Title
	AB   string   // Abstract
	MH   []string // MeSH Headings
	PT   []string // Publication Types
}

type MedlineDocuments []MedlineDocument

func UnmarshalMedline(r io.Reader) MedlineDocuments {
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
				doc.MH = append(doc.MH, content...)
			case PT:
				doc.PT = append(doc.PT, content...)
			}
			pair := strings.Split(line, "-")
			item = strings.TrimSpace(pair[0])
			content = []string{strings.TrimSpace(pair[1])}
		}
	}
	docs = append(docs, doc)
	return docs
}

func (m MedlineDocument) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n"))
	b.WriteString(fmt.Sprintf("PMID- %s\n", m.PMID))
	b.WriteString(fmt.Sprintf("TI  - %s\n", m.TI))
	b.WriteString(fmt.Sprintf("AB  - %s\n", m.AB))
	for _, mh := range m.MH {
		b.WriteString(fmt.Sprintf("MH  - %s\n", mh))
	}
	for _, pt := range m.PT {
		b.WriteString(fmt.Sprintf("PT  - %s\n", pt))
	}
	return b.String()
}
