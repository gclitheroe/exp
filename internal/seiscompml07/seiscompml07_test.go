package seiscompml07

import (
	"github.com/gclitheroe/exp/internal/quake"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
)

func TestUnmarshal(t *testing.T) {

	f, err := os.Open(filepath.Join("testdata", "2015p768477.xml"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	ep, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}

	if len(ep.Events) != 1 {
		t.Fatal("should have found 1 event.")
	}

	e := ep.Events[0]

	if e.PublicID != "2015p768477" {
		t.Errorf("expected publicID 2015p768477 got %s", e.PublicID)
	}

	if e.Type != "earthquake" {
		t.Errorf("expected type earthquake got %s", e.Type)
	}

	if e.CreationInfo.AgencyID != "WEL(GNS_Primary)" {
		t.Errorf("AgencyID expected WEL(GNS_Primary) got %s", e.CreationInfo.AgencyID)
	}

	if e.CreationInfo.CreationTime.Format(time.RFC3339Nano) != "2015-10-12T08:05:25.610839Z" {
		t.Errorf("CreationTime expected 2015-10-12T08:05:25.610839Z got %s", e.CreationInfo.CreationTime.Format(time.RFC3339Nano))
	}

	if e.CreationInfo.ModificationTime.Format(time.RFC3339Nano) != "2015-10-12T22:46:41.228824Z" {
		t.Errorf("ModificationTime expected 2015-10-12T22:46:41.228824Z got %s", e.CreationInfo.ModificationTime.Format(time.RFC3339Nano))
	}

	if e.PreferredOriginID != "NLL.20151012224503.620592.155845" {
		t.Errorf("expected preferredOriginID NLL.20151012224503.620592.155845 got %s", e.PreferredOriginID)
	}

	if e.PreferredOrigin.PublicID != "NLL.20151012224503.620592.155845" {
		t.Errorf("expected NLL.20151012224503.620592.155845 got %s", e.PreferredOrigin.PublicID)
	}

	if e.PreferredOrigin.Time.Value.Format(time.RFC3339Nano) != "2015-10-12T08:05:01.717692Z" {
		t.Errorf("expected 2015-10-12T08:05:01.717692Z, got %s", e.PreferredOrigin.Time.Value.Format(time.RFC3339Nano))
	}

	if e.PreferredOrigin.Latitude.Value != -40.57806609 {
		t.Errorf("Latitude expected -40.57806609 got %f", e.PreferredOrigin.Latitude.Value)
	}
	if e.PreferredOrigin.Latitude.Uncertainty != 1.922480006 {
		t.Errorf("Latitude uncertainty expected 1.922480006 got %f", e.PreferredOrigin.Latitude.Uncertainty)
	}

	if e.PreferredOrigin.Longitude.Value != 176.3257242 {
		t.Errorf("Longitude expected 176.3257242 got %f", e.PreferredOrigin.Longitude.Value)
	}
	if e.PreferredOrigin.Longitude.Uncertainty != 3.435738791 {
		t.Errorf("Longitude uncertainty expected 3.435738791 got %f", e.PreferredOrigin.Longitude.Uncertainty)
	}

	if e.PreferredOrigin.Depth.Value != 23.28125 {
		t.Errorf("Depth expected 23.28125 got %f", e.PreferredOrigin.Depth.Value)
	}
	if e.PreferredOrigin.Depth.Uncertainty != 3.575079654 {
		t.Errorf("Depth uncertainty expected 3.575079654 got %f", e.PreferredOrigin.Depth.Uncertainty)
	}

	if e.PreferredOrigin.MethodID != "NonLinLoc" {
		t.Errorf("MethodID expected NonLinLoc got %s", e.PreferredOrigin.MethodID)
	}

	if e.PreferredOrigin.EarthModelID != "nz3drx" {
		t.Errorf("EarthModelID expected NonLinLoc got %s", e.PreferredOrigin.EarthModelID)
	}

	if e.PreferredOrigin.Quality.StandardError != 0.5592857863 {
		t.Errorf("StandardError expected 0.5592857863 got %f", e.PreferredOrigin.Quality.StandardError)
	}

	if e.PreferredOrigin.Quality.AzimuthalGap != 166.4674465 {
		t.Errorf("AzimuthalGap expected 166.4674465 got %f", e.PreferredOrigin.Quality.AzimuthalGap)
	}

	if e.PreferredOrigin.Quality.MinimumDistance != 0.1217162272 {
		t.Errorf("MinimumDistance expected 0.1217162272 got %f", e.PreferredOrigin.Quality.MinimumDistance)
	}

	if e.PreferredOrigin.Quality.UsedPhaseCount != 44 {
		t.Errorf("UsedPhaseCount expected 44 got %d", e.PreferredOrigin.Quality.UsedPhaseCount)
	}

	if e.PreferredOrigin.Quality.UsedStationCount != 32 {
		t.Errorf("UsedStationCount expected 32 got %d", e.PreferredOrigin.Quality.UsedStationCount)
	}

	var found bool
	for _, v := range e.PreferredOrigin.Arrivals {
		if v.PickID == "Pick#20151012081200.115203.26387" {
			found = true
			if v.Phase != "P" {
				t.Errorf("expected P got %s", v.Phase)
			}

			if v.Azimuth != 211.917806 {
				t.Errorf("azimuth expected 211.917806 got %f", v.Azimuth)
			}

			if v.Distance != 0.1217162272 {
				t.Errorf("distance expected 0.1217162272 got %f", v.Distance)
			}

			if v.Weight != 1.406866218 {
				t.Errorf("weight expected 1.406866218 got %f", v.Weight)
			}

			if v.TimeResidual != -0.01664948232 {
				t.Errorf("time residual expected -0.01664948232 got %f", v.TimeResidual)
			}

			if v.Pick.WaveformID.NetworkCode != "NZ" {
				t.Errorf("Pick.WaveformID.NetworkCode expected NZ, got %s", v.Pick.WaveformID.NetworkCode)
			}

			if v.Pick.WaveformID.StationCode != "BFZ" {
				t.Errorf("Pick.WaveformID.StationCode expected BFZ, got %s", v.Pick.WaveformID.StationCode)
			}

			if v.Pick.WaveformID.LocationCode != "10" {
				t.Errorf("Pick.WaveformID.LocationCode expected 10, got %s", v.Pick.WaveformID.LocationCode)
			}

			if v.Pick.WaveformID.ChannelCode != "HHN" {
				t.Errorf("Pick.WaveformID.ChannelCode expected HHN, got %s", v.Pick.WaveformID.ChannelCode)
			}

			if v.Pick.EvaluationMode != "manual" {
				t.Errorf("Pick.WaveformID.EvaluationMode expected manual got %s", v.Pick.EvaluationMode)
			}

			if v.Pick.EvaluationStatus != "" {
				t.Errorf("Pick.WaveformID.EvaluationStatus expected empty string got %s", v.Pick.EvaluationStatus)
			}

			if v.Pick.Time.Value.Format(time.RFC3339Nano) != "2015-10-12T08:05:06.792207Z" {
				t.Errorf("Pick.Time expected 2015-10-12T08:05:06.792207Z got %s", v.Pick.Time.Value.Format(time.RFC3339Nano))
			}
		}

	}
	if !found {
		t.Error("didn't find PickID Pick#20151012081200.115203.26387")
	}

	if e.PreferredMagnitude.Type != "M" {
		t.Errorf("e.PreferredMagnitude.Type expected M got %s", e.PreferredMagnitude.Type)
	}
	if e.PreferredMagnitude.Magnitude.Value != 5.691131913 {
		t.Errorf("magnitude expected 5.691131913 got %f", e.PreferredMagnitude.Magnitude.Value)
	}
	if e.PreferredMagnitude.Magnitude.Uncertainty != 0 {
		t.Errorf("uncertainty expected 0 got %f", e.PreferredMagnitude.Magnitude.Uncertainty)
	}
	if e.PreferredMagnitude.StationCount != 171 {
		t.Errorf("e.PreferredMagnitude.StationCount expected 171 got %d", e.PreferredMagnitude.StationCount)
	}
	if e.PreferredMagnitude.MethodID != "weighted average" {
		t.Errorf("MethodID expected weighted average got %s", e.PreferredMagnitude.MethodID)
	}

	if e.PreferredMagnitude.CreationInfo.AgencyID != "WEL(GNS_Primary)" {
		t.Errorf("AgencyID expected WEL(GNS_Primary) got %s", e.PreferredMagnitude.CreationInfo.AgencyID)
	}

	if e.PreferredMagnitude.CreationInfo.CreationTime.Format(time.RFC3339Nano) != "2015-10-12T22:46:41.218145Z" {
		t.Errorf("CreationTime expected 2015-10-12T22:46:41.218145Z got %s", e.PreferredMagnitude.CreationInfo.CreationTime.Format(time.RFC3339Nano))
	}

	found = false

	for _, m := range e.PreferredOrigin.Magnitudes {
		if m.PublicID == "Magnitude#20151012224509.743338.156745" {
			found = true

			if m.Type != "ML" {
				t.Error("m.Type expected ML, got ", m.Type)
			}
			if m.Magnitude.Value != 6.057227661 {
				t.Errorf("magnitude expected 6.057227661 got %f", m.Magnitude.Value)
			}
			if m.Magnitude.Uncertainty != 0.2576927171 {
				t.Errorf("Uncertainty expected 0.2576927171 got %f", m.Magnitude.Uncertainty)
			}
			if m.StationCount != 23 {
				t.Errorf("m.StationCount expected 23 got %d", m.StationCount)
			}
			if m.MethodID != "trimmed mean" {
				t.Errorf("m.MethodID expected trimmed mean got %s", m.MethodID)
			}

			if !(len(m.StationMagnitudeContributions) > 1) {
				t.Error("expected more than 1 StationMagnitudeContribution")
			}

			var foundSM bool

			for _, s := range m.StationMagnitudeContributions {
				if s.StationMagnitudeID == "StationMagnitude#20151012224509.743511.156746" {
					foundSM = true

					if s.Weight != 1.0 {
						t.Errorf("Weight expected 1.0 got %f", s.Weight)
					}

					if s.StationMagnitude.Magnitude.Value != 6.096018735 {
						t.Errorf("StationMagnitude.Magnitude.Value expected 6.096018735 got %f", s.StationMagnitude.Magnitude.Value)
					}

					if s.StationMagnitude.Type != "ML" {
						t.Errorf("StationMagnitude.Type expected ML got %s", s.StationMagnitude.Type)
					}

					if s.StationMagnitude.WaveformID.NetworkCode != "NZ" {
						t.Errorf("Pick.WaveformID.NetworkCode expected NZ, got %s", s.StationMagnitude.WaveformID.NetworkCode)
					}

					if s.StationMagnitude.WaveformID.StationCode != "ANWZ" {
						t.Errorf("Pick.WaveformID.StationCode expected ANWZ, got %s", s.StationMagnitude.WaveformID.StationCode)
					}

					if s.StationMagnitude.WaveformID.LocationCode != "10" {
						t.Errorf("Pick.WaveformID.LocationCode expected 10, got %s", s.StationMagnitude.WaveformID.LocationCode)
					}

					if s.StationMagnitude.WaveformID.ChannelCode != "EH" {
						t.Errorf("Pick.WaveformID.ChannelCode expected EH, got %s", s.StationMagnitude.WaveformID.ChannelCode)
					}

					if s.StationMagnitude.Amplitude.Amplitude.Value != 21899.94892 {
						t.Errorf("Amplitude.Value expected 21899.94892 got %f", s.StationMagnitude.Amplitude.Amplitude.Value)
					}
				}
			}
			if !foundSM {
				t.Error("did not find StationMagnitudeContrib StationMagnitude#20151012224509.743511.156746")
			}
		}
	}

	if !found {
		t.Error("did not find magnitude smi:scs/0.7/Origin#20131202033820.196288.25287#netMag.MLv")
	}

	if e.ModificationTime().Format(time.RFC3339Nano) != "2015-10-12T22:46:41.228824Z" {
		t.Errorf("Modification time expected 2015-10-12T22:46:41.228824Z got %s", e.ModificationTime().Format(time.RFC3339Nano))
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
