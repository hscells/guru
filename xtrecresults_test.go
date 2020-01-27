package guru_test

import (
	"github.com/hscells/guru"
	"github.com/hscells/trecresults"
	"os"
	"testing"
)

func TestWriteCompressedTrecResultFile(t *testing.T) {
	fR, err := os.OpenFile("testdata/tar18_training.res", os.O_RDONLY, 0664)
	if err != nil {
		t.Fatal(err)
	}
	resultFile, err := trecresults.ResultsFromReader(fR)
	if err != nil {
		t.Fatal(err)
	}
	fC, err := os.OpenFile("testdata/tar18_training.xres", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		t.Fatal(err)
	}
	_, err = guru.WriteCompressedTrecResultFile(fC, resultFile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadCompressedTrecResultFile(t *testing.T) {
	f, err := os.OpenFile("testdata/tar18_training.xres", os.O_RDONLY, 0664)
	if err != nil {
		t.Fatal(err)
	}
	resultFile, err := guru.ReadCompressedTrecResultFile(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resultFile.Results))
	for k, v := range resultFile.Results {
		t.Log(k, len(v))
	}
}

func TestReadTrecResultFile(t *testing.T) {
	f, err := os.OpenFile("testdata/tar18_training.res", os.O_RDONLY, 0664)
	if err != nil {
		t.Fatal(err)
	}
	resultFile, err := trecresults.ResultsFromReader(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resultFile.Results))
	for k, v := range resultFile.Results {
		t.Log(k, len(v))
	}
}
