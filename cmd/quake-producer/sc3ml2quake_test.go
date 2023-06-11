package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gclitheroe/exp/internal/quake"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
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

	var b []byte

	q, err := fromSC3ML(filepath.Join("testdata", "2015p768477.xml"))
	if err != nil {
		return dir, c, err
	}

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

func TestQuakeProto(t *testing.T) {
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
	if err = proto.Unmarshal(b, e); err != nil {
		t.Error(err)
	}

	if e.PublicID != "2015p768477" {
		t.Errorf("expected publicID 2015p768477 got %s", e.PublicID)
	}

	if e.Type != "earthquake" {
		t.Errorf("expected type earthquake got %s", e.Type)
	}

	if e.Agency != "WEL(GNS_Primary)" {
		t.Errorf("Agency expected WEL(GNS_Primary) got %s", e.Agency)
	}

	if e.Time == nil {
		t.Fatal("nil time")
	}

	if e.ModificationTime == nil {
		t.Fatal("nil modification time")
	}

	if time.Unix(e.Time.Seconds, e.Time.Nanos).UTC().Format(time.RFC3339Nano) != "2015-10-12T08:05:01.717692Z" {
		t.Errorf("time expected 2015-10-12T08:05:01.717692Z got %s", time.Unix(e.Time.Seconds, e.Time.Nanos).UTC().Format(time.RFC3339Nano))
	}

	if e.Latitude != -40.57806609 {
		t.Errorf("Latitude expected -40.57806609 got %f", e.Latitude)
	}
	if e.LatitudeUncertainty != 1.922480006 {
		t.Errorf("Latitude uncertainty expected 1.922480006 got %f", e.LatitudeUncertainty)
	}

	if e.Longitude != 176.3257242 {
		t.Errorf("Longitude expected 176.3257242 got %f", e.Longitude)
	}
	if e.LongitudeUncertainty != 3.435738791 {
		t.Errorf("Longitude uncertainty expected 3.435738791 got %f", e.LongitudeUncertainty)
	}

	if e.Depth != 23.28125 {
		t.Errorf("Depth expected 23.28125 got %f", e.Depth)
	}
	if e.DepthUncertainty != 3.575079654 {
		t.Errorf("Depth uncertainty expected 3.575079654 got %f", e.DepthUncertainty)
	}

	if e.Method != "NonLinLoc" {
		t.Errorf("Method expected NonLinLoc got %s", e.Method)
	}

	if e.EarthModel != "nz3drx" {
		t.Errorf("EarthModel expected NonLinLoc got %s", e.EarthModel)
	}

	if e.StandardError != 0.5592857863 {
		t.Errorf("StandardError expected 0.5592857863 got %f", e.StandardError)
	}

	if e.AzimuthalGap != 166.4674465 {
		t.Errorf("AzimuthalGap expected 166.4674465 got %f", e.AzimuthalGap)
	}

	if e.MinimumDistance != 0.1217162272 {
		t.Errorf("MinimumDistance expected 0.1217162272 got %f", e.MinimumDistance)
	}

	if e.UsedPhaseCount != 44 {
		t.Errorf("UsedPhaseCount expected 44 got %d", e.UsedPhaseCount)
	}

	if e.UsedStationCount != 32 {
		t.Errorf("UsedStationCount expected 32 got %d", e.UsedStationCount)
	}

	if e.MagnitudeType != "M" {
		t.Errorf("e.MagnitudeType expected M got %s", e.MagnitudeType)
	}

	if e.Magnitude != 5.691131913 {
		t.Errorf("magnitude expected 5.691131913 got %f", e.Magnitude)
	}

	if e.MagnitudeUncertainty != 0 {
		t.Errorf("uncertainty expected 0 got %f", e.MagnitudeUncertainty)
	}

	if e.MagnitudeStationCount != 171 {
		t.Errorf("e.MagnitudeStationCount expected 171 got %d", e.MagnitudeStationCount)
	}

	if len(e.Magnitudes) != 3 {
		t.Errorf("expected 3 magnitudes got %d", len(e.Magnitudes))
	}

	var found bool
	for _, v := range e.Magnitudes {
		if v.Type == "ML" {
			found = true

			if v.Magnitude != 6.057227661 {
				t.Errorf("magnitude expected 6.057227661 got %f", v.Magnitude)
			}
			if v.MagnitudeUncertainty != 0.2576927171 {
				t.Errorf("Uncertainty expected 0.2576927171 got %f", v.MagnitudeUncertainty)
			}
			if v.StationCount != 23 {
				t.Errorf("v.StationCount expected 23 got %d", v.StationCount)
			}
			if v.Method != "trimmed mean" {
				t.Errorf("v.Method expected trimmed mean got %s", v.Method)
			}

			if len(v.StationMagnitude) != 23 {
				t.Errorf("station magnitudes expected 23 got %d", len(v.StationMagnitude))
			}
		}
	}

	if !found {
		t.Error("did not find magnitude ML")
	}

	if len(e.Phases) != 44 {
		t.Errorf("phases expected 44 got %d", len(e.Phases))
	}
}
