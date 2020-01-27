package guru

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/hscells/trecresults"
	"io"
)

var MaxTopicWidth = 10
var MaxDocIDWidth = 8

// ReadCompressedTrecResultFile reads a compressed trec results (run) file.
// File format is as follows:
//
// HEADER
// ------
// |byte        |byte       |'\n'              |
// |topic width |docID width|new line character|
//
// BODY
// ----
// |`topic width` bytes|8 bytes                        |k*`docID width` bytes|'0000000\n'|
// |topic              |number of documents (uint64), k|document IDs         |padding    |
//
// The HEADER section is written once and defines how the file should be read.
// The BODY section should be written to a single line such that multiple BODY sections
// may be concatenated together manually and the file can still be decompressed.
func ReadCompressedTrecResultFile(r io.Reader) (*trecresults.ResultFile, error) {
	resultFile := trecresults.NewResultFile()

	// Read first three bytes: [topic width, doc ID width, '\n'].
	p := make([]byte, 3)
	_, err := r.Read(p)
	if err != nil {
		return nil, err
	}
	topicWidth := int(p[0])
	docIDWidth := int(p[1])

	var (
		tok   int
		topic string
		k     int
	)

	// Read next set of bytes: [topic (string of size 'topic width'), number of results (uint64)].
	p = make([]byte, topicWidth+8)
	for {
		_, err = r.Read(p)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			return resultFile, nil
		}

		if tok == 0 {
			topic = fixedWidthStringify(p[:topicWidth])                 // First half of bytes is the topic.
			k = int(int64(binary.BigEndian.Uint64(p[topicWidth:])))     // Second half is the number of document IDs.
			resultFile.Results[topic] = make(trecresults.ResultList, k) // Allocate memory for the document IDs.
			p = make([]byte, (k*docIDWidth)+docIDWidth)                 // Allocate the bytes for the next read, which will contain the document IDs.
			tok = 1
		} else {
			for i, j := 0, 0; i < len(p)-docIDWidth; i += docIDWidth { // Iterate over each [doc ID (string of size 'doc ID width')].
				v := p[i : i+docIDWidth]
				if v[0] == 0 {
					fmt.Println(v)
					break
				}
				resultFile.Results[topic][j] = &trecresults.Result{
					DocId: fixedWidthStringify(v),
				}
				j++
			}
			p = make([]byte, topicWidth+8) // Ready to read the next topic, so allocate the bytes for it.
			tok = 0
		}
	}
}

// WriteCompressedTrecResultFile writes a result file in memory to disk.
// See `ReadCompressedTrecResultFile` for file format.
func WriteCompressedTrecResultFile(w io.Writer, res trecresults.ResultFile) (int, error) {
	// The first three bytes are the topic width, the doc ID width, and a newline character.
	n, err := w.Write([]byte{byte(MaxTopicWidth), byte(MaxDocIDWidth), '\n'})
	if err != nil {
		return n, err
	}
	for topic, resultList := range res.Results {

		// Write the topic at the start of the line.
		l, err := w.Write(fixWidthString(topic, MaxTopicWidth))
		if err != nil {
			return n, err
		}
		n += l

		p := make([]byte, 8)
		binary.BigEndian.PutUint64(p, uint64(len(resultList)))
		l, err = w.Write(p)
		if err != nil {
			return n, err
		}
		l += n

		// Write each docID in fixed bytes.
		var buff bytes.Buffer
		for _, result := range resultList {
			_, err := buff.Write(fixWidthString(result.DocId, MaxDocIDWidth))
			if err != nil {
				return n, err
			}
		}
		// That's the end of a result list, the next one will appear on the next line.
		endDoc := make([]byte, MaxDocIDWidth)
		endDoc[len(endDoc)-1] = '\n'
		_, err = buff.Write(endDoc)
		if err != nil {
			return n, err
		}

		l, err = w.Write(buff.Bytes())
		if err != nil {
			return n, err
		}
		n += l

	}
	return n, nil
}

// fixWidthString converts a string to a fixed width byte slice.
// The byte slice is padded with 0s.
func fixWidthString(s string, width int) []byte {
	b := make([]byte, width)
	for i := 0; i < width; i++ {
		if i < len(s) {
			b[i] = s[i]
		} else {
			b[i] = 0
		}
	}
	return b
}

// fixedWidthStringify converts a byte slice to a string, removing padded 0s.
func fixedWidthStringify(b []byte) string {
	var s string
	for _, v := range b {
		if v == 0 {
			return s
		}
		s += string(v)
	}
	return s
}
