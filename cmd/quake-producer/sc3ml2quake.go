package main

import (
	"github.com/gclitheroe/exp/internal/quake"
	"github.com/gclitheroe/exp/internal/sc3ml"
)

/*
fromSC3ML converts sc3ml.Event to a Quake.
Only Arrivals and StationMagnitudeContribution that have contributed
to Origins or Magnitudes (Weight > 0) are included in the Quake.
*/
func fromSC3ML(e sc3ml.Event) quake.Quake {
	mt := e.ModificationTime()

	q := quake.Quake{
		PublicID: e.PublicID,
		Type:     e.Type,
		Agency:   e.CreationInfo.AgencyID,
		ModificationTime: &quake.Timestamp{
			Seconds: mt.Unix(),
			Nanos:   int64(mt.Nanosecond()),
		},
		Time: &quake.Timestamp{
			Seconds: e.PreferredOrigin.Time.Value.Unix(),
			Nanos:   int64(e.PreferredOrigin.Time.Value.Nanosecond()),
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
					Seconds: v.Pick.Time.Value.Unix(),
					Nanos:   int64(v.Pick.Time.Value.Nanosecond()),
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
			Type:                 m.Type,
			Method:               m.MethodID,
			StationCount:         m.StationCount,
		}

		for _, v := range m.StationMagnitudeContributions {
			if v.Weight > 0.0 {
				s := &quake.StationMagnitude{
					Weight:       v.Weight,
					NetworkCode:  v.StationMagnitude.WaveformID.NetworkCode,
					StationCode:  v.StationMagnitude.WaveformID.StationCode,
					LocationCode: v.StationMagnitude.WaveformID.LocationCode,
					ChannelCode:  v.StationMagnitude.WaveformID.ChannelCode,
					Magnitude:    v.StationMagnitude.Magnitude.Value,
					Type:         v.StationMagnitude.Type,
					Amplitude:    v.StationMagnitude.Amplitude.Amplitude.Value,
				}

				mag.StationMagnitude = append(mag.StationMagnitude, s)
			}
		}
		q.Magnitudes = append(q.Magnitudes, mag)
	}

	return q
}
