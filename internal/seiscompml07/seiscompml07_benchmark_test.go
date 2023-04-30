package seiscompml07

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gclitheroe/exp/internal/quake"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/protobuf/proto"
)

// setup creates XML, JSON, and Protobuf testing files in direcory dir.
// The files contain the same quake information in different formats
// cleanup() will never be nil and should always be called even if
// err is non nil.
func setup() (dir string, cleanup func() error, err error) {
	dir, err = os.MkdirTemp("", "")
	if err != nil {
		return dir, func() error { return nil }, err
	}

	c := func() error {
		return os.RemoveAll(dir)
	}

	var ep EventParameters
	var f *os.File
	var b []byte

	if f, err = os.Open(filepath.Join("testdata", "2015p768477.xml")); err != nil {
		return dir, c, err
	}

	defer f.Close()

	if b, err = io.ReadAll(f); err != nil {
		return dir, c, err
	}

	if ep, err = Unmarshal(b); err != nil {
		return dir, c, err
	}

	if len(ep.Events) != 1 {
		err = fmt.Errorf("should have found 1 event")
		return dir, c, err
	}

	q := FromSeiscompml07(ep.Events[0])
	if b, err = proto.Marshal(&q); err != nil {
		return dir, c, err
	}

	if err = os.WriteFile(filepath.Join(dir, "2015p768477.pb"), b, 0644); err != nil {
		return dir, c, err
	}

	if b, err = xml.MarshalIndent(q, "  ", "    "); err != nil {
		return
	}

	if err = os.WriteFile(filepath.Join(dir, "2015p768477.xml"), b, 0644); err != nil {
		return
	}

	if b, err = json.MarshalIndent(q, "  ", "    "); err != nil {
		return dir, c, err
	}

	if err = os.WriteFile(filepath.Join(dir, "2015p768477.json"), b, 0644); err != nil {
		return dir, c, err
	}

	return dir, c, nil
}

func BenchmarkUnmarshalQuakeXML(t *testing.B) {
	dir, cleanup, err := setup()
	defer cleanup()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(filepath.Join(dir, "2015p768477.xml"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	e := &quake.Quake{}

	err = xml.Unmarshal(b, e)
	if err != nil {
		t.Fatal(err)
	}

	if e.PublicID != "2015p768477" {
		t.Fatalf("expected publicID 2015p768477 got %s", e.PublicID)
	}

	for n := 0; n < t.N; n++ {
		// ignore errors
		xml.Unmarshal(b, e)
	}
}

func BenchmarkUnmarshalQuakeJSON(t *testing.B) {
	dir, cleanup, err := setup()
	defer cleanup()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(filepath.Join(dir, "2015p768477.json"))
	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	e := &quake.Quake{}

	err = json.Unmarshal(b, e)
	if err != nil {
		t.Fatal(err)
	}

	if e.PublicID != "2015p768477" {
		t.Fatalf("expected publicID 2015p768477 got %s", e.PublicID)
	}

	for n := 0; n < t.N; n++ {
		// ignore errors
		json.Unmarshal(b, e)
	}
}

func BenchmarkUnmarshalQuakeProtobuf(t *testing.B) {
	dir, cleanup, err := setup()
	defer cleanup()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(filepath.Join(dir, "2015p768477.pb"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return
	}

	e := &quake.Quake{}

	err = proto.Unmarshal(b, e)
	if err != nil {
		t.Fatal(err)
	}

	if e.PublicID != "2015p768477" {
		t.Fatalf("expected publicID 2015p768477 got %s", e.PublicID)
	}

	for n := 0; n < t.N; n++ {
		// ignore errors
		proto.Unmarshal(b, e)
	}
}
