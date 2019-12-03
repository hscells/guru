package guru

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	PMID = "PMID"
	TI   = "TI"
	BTI  = "BTI"
	CTI  = "CTI"
	AB   = "AB"
	MH   = "MH"
	PT   = "PT"
	AU   = "AU"
	DCOM = "DCOM"
)

type MedlineDocument struct {
	PMID string   // PubMed ID
	TI   string   // Title
	AB   string   // Abstract
	DCOM string   // Date Completed
	MH   []string // MeSH Headings
	PT   []string // Publication Types
	AU   []string // Authors
}

type MedlineDocuments []MedlineDocument

func UnmarshalMedline(r io.Reader) MedlineDocuments {

	const (
		Bnil int = iota
		Bpmid
		Bti
		Bbti
		Bcti
		Bab
		Bmh
		Bpt
		Bau
		Bdcom
	)
	s := bufio.NewScanner(r)
	var (
		item int
		docs MedlineDocuments
		doc  MedlineDocument
	)
	c := new(strings.Builder)
	for s.Scan() {
		// There are a number of articles without titles.
		// These variables will store book and collection title values.
		var bti, cti string
		line := s.Bytes()
		if len(line) == 0 && len(doc.PMID) > 0 {
			if len(doc.TI) == 0 && len(bti) > 0 {
				doc.TI = bti
			} else if len(doc.TI) == 0 && len(cti) > 0 {
				doc.TI = cti
			}
			if len(doc.TI) == 0 {
				log.Printf("parsing %s with no title\n", doc.PMID)
			}
			docs = append(docs, doc)
			doc = MedlineDocument{}
		} else if len(line) >= 5 && bytes.Equal(line[:6], []byte("      ")) {
			_, err := c.Write(bytes.Replace(line, []byte("     "), []byte(""), 1))
			if err != nil {
				panic(err)
			}
		} else {
			switch item {
			case Bpmid:
				doc.PMID = c.String()
			case Bti:
				doc.TI = c.String()
			case Bbti:
				bti = c.String()
			case Bcti:
				cti = c.String()
			case Bab:
				doc.AB = c.String()
			case Bdcom:
				doc.DCOM = c.String()
			case Bmh:
				doc.MH = append(doc.MH, c.String())
			case Bpt:
				doc.PT = append(doc.PT, c.String())
			case Bau:
				doc.AU = append(doc.AU, c.String())
			case Bnil:
				break
			}
			pair := bytes.Split(line, []byte("-"))
			if len(pair) <= 1 {
				continue
			}
			p0 := bytes.TrimSpace(pair[0])
			if bytes.Equal(p0, []byte(PMID)) {
				item = Bpmid
			} else if bytes.Equal(p0, []byte(TI)) {
				item = Bti
			} else if bytes.Equal(p0, []byte(BTI)) {
				item = Bbti
			} else if bytes.Equal(p0, []byte(CTI)) {
				item = Bcti
			} else if bytes.Equal(p0, []byte(AB)) {
				item = Bab
			} else if bytes.Equal(p0, []byte(MH)) {
				item = Bmh
			} else if bytes.Equal(p0, []byte(PT)) {
				item = Bpt
			} else if bytes.Equal(p0, []byte(AU)) {
				item = Bau
			} else if bytes.Equal(p0, []byte(DCOM)) {
				item = Bdcom
			} else {
				item = Bnil
			}

			c = new(strings.Builder)
			_, err := c.Write(bytes.TrimSpace(bytes.Join(pair[1:], []byte("-"))))
			if err != nil {
				panic(err)
			}
		}
	}
	docs = append(docs, doc)
	return docs
}

func UnmarshalAbstract(r io.Reader) MedlineDocuments {
	s := bufio.NewScanner(r)
	var (
		docs MedlineDocuments
		doc  MedlineDocument
		i    int
	)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			i++
			continue
		}

		if len(line) >= 5 && line[0:5] == "PMID:" {
			doc.PMID = strings.TrimSpace(strings.Split(line, " ")[1])
			docs = append(docs, doc)
			doc = MedlineDocument{}
			i = -1
			continue
		}

		switch i {
		case 2:
			doc.TI = doc.TI + line
		case 5:
			doc.AB = doc.AB + line
		}
	}
	return docs
}

func (m MedlineDocument) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n"))
	b.WriteString(fmt.Sprintf("PMID- %s\n", m.PMID))
	b.WriteString(fmt.Sprintf("TI  - %s\n", m.TI))
	b.WriteString(fmt.Sprintf("AB  - %s\n", m.AB))
	b.WriteString(fmt.Sprintf("DCOM  - %s\n", m.DCOM))
	for _, au := range m.AU {
		b.WriteString(fmt.Sprintf("AU  - %s\n", au))
	}
	for _, pt := range m.PT {
		b.WriteString(fmt.Sprintf("PT  - %s\n", pt))
	}
	for _, mh := range m.MH {
		b.WriteString(fmt.Sprintf("MH  - %s\n", mh))
	}
	return b.String()
}
