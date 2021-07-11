package carbon

import (
	"bytes"
	"strconv"
)

const (
	nextPart      = " "
	nextNode      = "."
	nextTag       = ";"
	nextTagValue  = "="
	nextDatapoint = "\n"
)

// Datapoint is a value stored at a timestamp bucket.
// If no value is recorded at a particular timestamp bucket in a series,
// the value will be None (null)
// Doc: https://graphite.readthedocs.io/en/latest/terminology.html
type Datapoint struct {
	path      *bytes.Buffer
	value     float64
	timestamp int64
}

// NewDatapoint is a datapoint building function
func NewDatapoint(name string, value float64, timestamp int64) *Datapoint {
	datapoint := &Datapoint{
		path:      &bytes.Buffer{},
		value:     value,
		timestamp: timestamp,
	}
	datapoint.path.WriteString(name)
	return datapoint
}

// WithPrefix adds provided prefix to datapoint
// Every function call extends datapoint path from start
// For example WithPrefix("second").WithPrefix("first") is
// equal to WithPrefix("first.second")
func (d *Datapoint) WithPrefix(prefix string) *Datapoint {
	path := d.path.String()
	d.path.Reset()
	d.path.WriteString(prefix)
	d.path.WriteString(nextNode)
	d.path.WriteString(path)
	return d
}

// WithTag adds provided tag=value pair to datapoint
// Every function call extends datapoint path
// For example WithTag("tag1","tag1Value").WithTag("tag2","tag2Value") adds ";tag1=tag1Value;tag2=tag2Value"
// While WithTag("tag2","tag2Value").WithTag("tag1","tag1Value") adds ";tag2=tag2Value;tag1=tag1Value"
func (d *Datapoint) WithTag(name, value string) *Datapoint {
	d.path.WriteString(nextTag)
	d.path.WriteString(name)
	d.path.WriteString(nextTagValue)
	d.path.WriteString(value)
	return d
}

// String returns Graphite Plaintext record form of datapoint
func (d *Datapoint) String() string {
	var (
		value     = strconv.FormatFloat(d.value, 'f', -1, 64)
		timestamp = strconv.FormatInt(d.timestamp, 10)
	)
	d.path.WriteString(nextPart)
	d.path.WriteString(value)
	d.path.WriteString(nextPart)
	d.path.WriteString(timestamp)
	d.path.WriteString(nextDatapoint)
	return d.path.String()
}
