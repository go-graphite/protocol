package carbon

import "strconv"

const (
	nextPart      = " "
	nextNode      = "."
	nextTag       = ";"
	nextTagValue  = "="
	nextDatapoint = "\n"
)

// Datapoint is a value stored at a timestamp bucket.
// If no value is recorded at a particular timestamp bucket in a series,
// the value will be None (null).
// Doc: https://graphite.readthedocs.io/en/latest/terminology.html
type Datapoint struct {
	path      string
	value     float64
	timestamp int64
}

// NewDatapoint is a datapoint building function
func NewDatapoint(name string, value float64, timestamp int64) *Datapoint {
	return &Datapoint{
		path:      name,
		value:     value,
		timestamp: timestamp,
	}
}

// WithPrefix adds provided prefix to datapoint
// Every function call extends datapoint path from start
// For example WithPrefix("second").WithPrefix("first") is
// equal to WithPrefix("first.second")
func (d *Datapoint) WithPrefix(prefix string) *Datapoint {
	d.path = prefix + nextNode + d.path
	return d
}

// WithTag adds provided tag=value pair to datapoint
// Every function call extends datapoint path
// For example WithTag("tag1","tag1Value").WithTag("tag2","tag2Value") adds ";tag1=tag1Value;tag2=tag2Value"
// While WithTag("tag2","tag2Value").WithTag("tag1","tag1Value") adds ";tag2=tag2Value;tag1=tag1Value"
func (d *Datapoint) WithTag(name, value string) *Datapoint {
	d.path += nextTag + name + nextTagValue + value
	return d
}

// Plaintext returns Graphite Plaintext record form of datapoint
func (d *Datapoint) Plaintext() string {
	var (
		value     = strconv.FormatFloat(d.value, 'f', -1, 64)
		timestamp = strconv.FormatInt(d.timestamp, 10)
	)
	return d.path + nextPart + value + nextPart + timestamp + nextDatapoint
}
