package guru

import (
	"encoding/gob"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Protocol is a representation of a systematic review Protocol in XML.
type Protocol struct {
	Objective          string `xml:"objective"`
	TypeOfStudy        string `xml:"type_of_study"`
	Participants       string `xml:"participants"`
	IndexTests         string `xml:"index_tests"`
	TargetConditions   string `xml:"target_conditions"`
	ReferenceStandards string `xml:"reference_standards"`
}

type Protocols map[string]Protocol

func ReadAndWriteProtocols(protocolsDir, protocolsBinFile string) Protocols {
	gob.Register(Protocol{})
	gob.Register(Protocols{})

	// First, get a list of files in the directory.
	files, err := ioutil.ReadDir(protocolsDir)
	if err != nil {
		panic(err)
	}

	protocols := make(Protocols)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if len(f.Name()) == 0 {
			continue
		}

		p := path.Join(protocolsDir, f.Name())
		source, err := ioutil.ReadFile(p)
		if err != nil {
			panic(err)
		}

		var protocol Protocol
		err = xml.Unmarshal(source, &protocol)
		if err != nil {
			panic(err)
		}

		_, topic := path.Split(p)

		protocols[strings.TrimSpace(topic)] = protocol
	}
	f, err := os.OpenFile(protocolsBinFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = gob.NewEncoder(f).Encode(protocols)
	if err != nil {
		panic(err)
	}
	return protocols
}

func LoadProtocols(protocolsBinFile string) Protocols {
	gob.Register(Protocol{})
	gob.Register(Protocols{})

	f, err := os.OpenFile(protocolsBinFile, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var protocols Protocols
	err = gob.NewDecoder(f).Decode(&protocols)
	if err != nil {
		panic(err)
	}
	return protocols
}
