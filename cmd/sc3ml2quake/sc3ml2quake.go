// Converts SeisCompML (XML) to quake protobufs.
// Usage: sc3ml2quake input-dir output-dir
// Usage: sc3ml2quake input-dir output-dir
// Source data s3://seiscompml07
package main

import (
	"flag"
	"github.com/gclitheroe/exp/internal/quake"
	"github.com/gclitheroe/exp/internal/sc3ml"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
)

type files struct {
	in, out string
}

func main() {
	var inDir, outDir string

	flag.StringVar(&inDir, "input-dir", "/work/sc3ml", "directory with input sc3ml files")
	flag.StringVar(&outDir, "output-dir", "/work/quake", "directory for quake protobuf file output")

	flag.Parse()

	d, err := os.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	// use a chan to fan out the work to as many processor functions as there are cores.
	sc3ml := make(chan files)

	go func() {
		defer close(sc3ml)

		var o string
		var xml bool

		for _, f := range d {
			if f.Type().IsRegular() {
				o, xml = strings.CutSuffix(f.Name(), ".xml")
				if !xml {
					continue
				}

				sc3ml <- files{in: inDir + string(os.PathSeparator) + f.Name(), out: outDir + string(os.PathSeparator) + o + ".pb"}
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			procSC3ML(sc3ml)
			wg.Done()
		}()
	}
	wg.Wait()
}

func procSC3ML(sc3ml <-chan files) {
	var b []byte
	var err error
	var q quake.Quake

	for f := range sc3ml {
		log.Println(f.in)

		q, err = fromSC3ML(f.in)
		if err != nil {
			log.Println(err)
			continue
		}

		b, err = proto.Marshal(&q)
		if err != nil {
			log.Println(err)
			continue
		}

		err = os.WriteFile(f.out, b, 0644)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

/*
fromSC3ML converts sc3ml.Event to a Quake.
Only Arrivals and StationMagnitudeContribution that have contributed
to Origins or Magnitudes (Weight > 0) are included in the Quake.
*/
func fromSC3ML(file string) (quake.Quake, error) {
	var b []byte
	var q quake.Quake

	f, err := os.Open(file)
	if err != nil {
		return q, err
	}

	defer f.Close()

	b, err = io.ReadAll(f)
	if err != nil {
		return q, err
	}

	ep, err := sc3ml.Unmarshal(b)
	if err != nil {
		return q, err
	}

	var e sc3ml.Event = ep.Events[0]

	mt := e.ModificationTime()

	q = quake.Quake{
		PublicID:  e.PublicID,
		QuakeType: e.Type,
		Agency:    e.CreationInfo.AgencyID,
		ModificationTime: &quake.Timestamp{
			Secs:  mt.Unix(),
			Nanos: int64(mt.Nanosecond()),
		},
		Time: &quake.Timestamp{
			Secs:  e.PreferredOrigin.Time.Value.Unix(),
			Nanos: int64(e.PreferredOrigin.Time.Value.Nanosecond()),
		},
		Latitude:              e.PreferredOrigin.Latitude.Value,
		LatitudeUncertainty:   e.PreferredOrigin.Latitude.Uncertainty,
		Longitude:             e.PreferredOrigin.Longitude.Value,
		LongitudeUncertainty:  e.PreferredOrigin.Longitude.Uncertainty,
		Depth:                 e.PreferredOrigin.Depth.Value,
		DepthUncertainty:      e.PreferredOrigin.Depth.Uncertainty,
		Method:                e.PreferredOrigin.MethodID,
		EarthModel:            e.PreferredOrigin.EarthModelID,
		EvaluationMode:        e.PreferredOrigin.EvaluationMode,
		EvaluationStatus:      e.PreferredOrigin.EvaluationStatus,
		UsedPhaseCount:        e.PreferredOrigin.Quality.UsedPhaseCount,
		UsedStationCount:      e.PreferredOrigin.Quality.UsedStationCount,
		StandardError:         e.PreferredOrigin.Quality.StandardError,
		AzimuthalGap:          e.PreferredOrigin.Quality.AzimuthalGap,
		MinimumDistance:       e.PreferredOrigin.Quality.MinimumDistance,
		Magnitude:             e.PreferredMagnitude.Magnitude.Value,
		MagnitudeUncertainty:  e.PreferredMagnitude.Magnitude.Uncertainty,
		MagnitudeType:         e.PreferredMagnitude.Type,
		MagnitudeStationCount: e.PreferredMagnitude.StationCount,
	}

	for _, v := range e.PreferredOrigin.Arrivals {
		if v.Weight > 0.0 {
			p := &quake.Phase{
				Time: &quake.Timestamp{
					Secs:  v.Pick.Time.Value.Unix(),
					Nanos: int64(v.Pick.Time.Value.Nanosecond()),
				},
				Phase:            v.Phase,
				Residual:         v.TimeResidual,
				Weight:           v.Weight,
				Azimuth:          v.Azimuth,
				Distance:         v.Distance,
				NetworkCode:      v.Pick.WaveformID.NetworkCode,
				StationCode:      v.Pick.WaveformID.StationCode,
				LocationCode:     v.Pick.WaveformID.LocationCode,
				ChannelCode:      v.Pick.WaveformID.ChannelCode,
				EvaluationMode:   v.Pick.EvaluationMode,
				EvaluationStatus: v.Pick.EvaluationStatus,
			}

			q.Phases = append(q.Phases, p)
		}
	}

	for _, m := range e.PreferredOrigin.Magnitudes {
		mag := &quake.Magnitude{
			Magnitude:            m.Magnitude.Value,
			MagnitudeUncertainty: m.Magnitude.Uncertainty,
			MagnitudeType:        m.Type,
			MagnitudeMethod:      m.MethodID,
			StationCount:         m.StationCount,
		}

		for _, v := range m.StationMagnitudeContributions {
			if v.Weight > 0.0 {
				s := &quake.StationMagnitude{
					Weight:        v.Weight,
					NetworkCode:   v.StationMagnitude.WaveformID.NetworkCode,
					StationCode:   v.StationMagnitude.WaveformID.StationCode,
					LocationCode:  v.StationMagnitude.WaveformID.LocationCode,
					ChannelCode:   v.StationMagnitude.WaveformID.ChannelCode,
					Magnitude:     v.StationMagnitude.Magnitude.Value,
					MagnitudeType: v.StationMagnitude.Type,
					Amplitude:     v.StationMagnitude.Amplitude.Amplitude.Value,
				}

				mag.StationMagnitude = append(mag.StationMagnitude, s)
			}
		}
		q.Magnitudes = append(q.Magnitudes, mag)
	}

	return q, nil
}
