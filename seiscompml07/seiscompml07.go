/*
Package seiscompml07 is for parsing SeisComPML version 0.7
e.g., http://geofon.gfz-potsdam.de/schema/0.7/sc3ml_0.7.xsd

The schema allows many elements to be 0..1 or 0..* The elements
mapped here are assumed to be present and will have zero values
if not.
*/
package seiscompml07

import (
	"encoding/xml"
	"sort"
	"time"
)

// Seiscomp is for parsing SeisComPML
type Seiscomp struct {
	EventParameters EventParameters `xml:"EventParameters"`
}

// EventParameters is for parsing SeisComPML
type EventParameters struct {
	Events     []Event     `xml:"event"`
	Picks      []Pick      `xml:"pick"`
	Amplitudes []Amplitude `xml:"amplitude"`
	Origins    []Origin    `xml:"origin"`
}

// Event is for parsing SeisComPML
type Event struct {
	PublicID             string       `xml:"publicID,attr"`
	PreferredOriginID    string       `xml:"preferredOriginID"`
	PreferredMagnitudeID string       `xml:"preferredMagnitudeID"`
	Type                 string       `xml:"type"`
	CreationInfo         CreationInfo `xml:"creationInfo"`
	PreferredOrigin      Origin
	PreferredMagnitude   Magnitude
}

// CreationInfo is for parsing SeisComPML
type CreationInfo struct {
	AgencyID         string    `xml:"agencyID"`
	CreationTime     time.Time `xml:"creationTime"`
	ModificationTime time.Time `xml:"modificationTime"`
}

// Origin is for parsing SeisComPML
type Origin struct {
	PublicID          string             `xml:"publicID,attr"`
	Time              TimeValue          `xml:"time"`
	Latitude          RealQuantity       `xml:"latitude"`
	Longitude         RealQuantity       `xml:"longitude"`
	Depth             RealQuantity       `xml:"depth"`
	DepthType         string             `xml:"depthType"`
	MethodID          string             `xml:"methodID"`
	EarthModelID      string             `xml:"earthModelID"`
	Quality           Quality            `xml:"quality"`
	EvaluationMode    string             `xml:"evaluationMode"`
	EvaluationStatus  string             `xml:"evaluationStatus"`
	Arrivals          []Arrival          `xml:"arrival"`
	StationMagnitudes []StationMagnitude `xml:"stationMagnitude"`
	Magnitudes        []Magnitude        `xml:"magnitude"`
	CreationInfo      CreationInfo       `xml:"creationInfo"`
}

// Quality is for parsing SeisComPML
type Quality struct {
	UsedPhaseCount   int64   `xml:"usedPhaseCount"`
	UsedStationCount int64   `xml:"usedStationCount"`
	StandardError    float64 `xml:"standardError"`
	AzimuthalGap     float64 `xml:"azimuthalGap"`
	MinimumDistance  float64 `xml:"minimumDistance"`
}

// Arrival is for parsing SeisComPML
type Arrival struct {
	PickID       string  `xml:"pickID"`
	Phase        string  `xml:"phase"`
	Azimuth      float64 `xml:"azimuth"`
	Distance     float64 `xml:"distance"`
	TimeResidual float64 `xml:"timeResidual"`
	Weight       float64 `xml:"weight"`
	Pick         Pick
}

// Pick is for parsing SeisComPML
type Pick struct {
	PublicID         string     `xml:"publicID,attr"`
	Time             TimeValue  `xml:"time"`
	WaveformID       WaveformID `xml:"waveformID"`
	EvaluationMode   string     `xml:"evaluationMode"`
	EvaluationStatus string     `xml:"evaluationStatus"`
}

// WaveformID is for parsing SeisComPML
type WaveformID struct {
	NetworkCode  string `xml:"networkCode,attr"`
	StationCode  string `xml:"stationCode,attr"`
	LocationCode string `xml:"locationCode,attr"`
	ChannelCode  string `xml:"channelCode,attr"`
}

// RealQuantity is for parsing SeisComPML
type RealQuantity struct {
	Value       float64 `xml:"value"`
	Uncertainty float64 `xml:"uncertainty"`
}

// TimeValue is for parsing SeisComPML
type TimeValue struct {
	Value time.Time `xml:"value"`
}

// Magnitude is for parsing SeisComPML
type Magnitude struct {
	PublicID                      string                         `xml:"publicID,attr"`
	Magnitude                     RealQuantity                   `xml:"magnitude"`
	Type                          string                         `xml:"type"`
	MethodID                      string                         `xml:"methodID"`
	StationCount                  int64                          `xml:"stationCount"`
	StationMagnitudeContributions []StationMagnitudeContribution `xml:"stationMagnitudeContribution"`
	CreationInfo                  CreationInfo                   `xml:"creationInfo"`
}

// StationMagnitudeContribution is for parsing SeisComPML
type StationMagnitudeContribution struct {
	StationMagnitudeID string  `xml:"stationMagnitudeID"`
	Weight             float64 `xml:"weight"`
	StationMagnitude   StationMagnitude
}

// StationMagnitude is for parsing SeisComPML
type StationMagnitude struct {
	PublicID    string       `xml:"publicID,attr"`
	Magnitude   RealQuantity `xml:"magnitude"`
	Type        string       `xml:"type"`
	AmplitudeID string       `xml:"amplitudeID"`
	WaveformID  WaveformID   `xml:"waveformID"`
	Amplitude   Amplitude
}

// Amplitude is for parsing SeisComPML
type Amplitude struct {
	PublicID  string       `xml:"publicID,attr"`
	Amplitude RealQuantity `xml:"amplitude"`
}

// Unmarshal unmarshals the SeisComPML in b and initialises all
// the objects referenced by ID in the SeisComPML e.g., PreferredOrigin,
// PreferredMagnitude etc.
func Unmarshal(b []byte) (EventParameters, error) {
	var q Seiscomp

	if err := xml.Unmarshal(b, &q); err != nil {
		return q.EventParameters, err
	}

	var picks = make(map[string]Pick)
	for k, v := range q.EventParameters.Picks {
		picks[v.PublicID] = q.EventParameters.Picks[k]
	}

	var amplitudes = make(map[string]Amplitude)
	for k, v := range q.EventParameters.Amplitudes {
		amplitudes[v.PublicID] = q.EventParameters.Amplitudes[k]
	}

	for i := range q.EventParameters.Origins {
		for k, v := range q.EventParameters.Origins[i].Arrivals {
			q.EventParameters.Origins[i].Arrivals[k].Pick = picks[v.PickID]
		}

		var stationMagnitudes = make(map[string]StationMagnitude)

		for k, v := range q.EventParameters.Origins[i].StationMagnitudes {
			q.EventParameters.Origins[i].StationMagnitudes[k].Amplitude = amplitudes[v.AmplitudeID]
			stationMagnitudes[v.PublicID] = q.EventParameters.Origins[i].StationMagnitudes[k]
		}

		for j := range q.EventParameters.Origins[i].Magnitudes {
			for k, v := range q.EventParameters.Origins[i].Magnitudes[j].StationMagnitudeContributions {
				q.EventParameters.Origins[i].Magnitudes[j].StationMagnitudeContributions[k].StationMagnitude = stationMagnitudes[v.StationMagnitudeID]
			}
		}
	}

	for i := range q.EventParameters.Events {

		for k, v := range q.EventParameters.Origins {
			if v.PublicID == q.EventParameters.Events[i].PreferredOriginID {
				q.EventParameters.Events[i].PreferredOrigin = q.EventParameters.Origins[k]

				for l, m := range q.EventParameters.Events[i].PreferredOrigin.Magnitudes {
					if m.PublicID == q.EventParameters.Events[i].PreferredMagnitudeID {
						q.EventParameters.Events[i].PreferredMagnitude = q.EventParameters.Events[i].PreferredOrigin.Magnitudes[l]
					}
				}
			}
		}

	}

	return q.EventParameters, nil
}

// ModificationTime returns the most recent creation or modification time
// for the Event, PreferredOrigin, or PreferredMagnitude.
func (e *Event) ModificationTime() time.Time {
	var t []string

	t = append(t, e.CreationInfo.CreationTime.Format(time.RFC3339Nano))
	t = append(t, e.CreationInfo.ModificationTime.Format(time.RFC3339Nano))
	t = append(t, e.PreferredOrigin.CreationInfo.CreationTime.Format(time.RFC3339Nano))
	t = append(t, e.PreferredOrigin.CreationInfo.ModificationTime.Format(time.RFC3339Nano))
	t = append(t, e.PreferredMagnitude.CreationInfo.CreationTime.Format(time.RFC3339Nano))
	t = append(t, e.PreferredMagnitude.CreationInfo.ModificationTime.Format(time.RFC3339Nano))

	sort.Sort(sort.Reverse(sort.StringSlice(t)))

	tm, _ := time.Parse(time.RFC3339Nano, t[0])
	return tm
}
