package carbon

import (
	"bytes"
	"strconv"
)

var (
	nextPart      = []byte(" ")
	nextNode      = []byte(".")
	nextTag       = []byte(";")
	nextTagValue  = []byte("=")
	nextDatapoint = []byte("\n")
)

// Metric is a Graphite metric builder
type Metric struct {
	path *bytes.Buffer
	tags *bytes.Buffer
}

// NewMetric returns new Metric
func NewMetric(name string) *Metric {
	metric := &Metric{
		path: &bytes.Buffer{},
		tags: &bytes.Buffer{},
	}
	metric.path.WriteString(name)
	return metric
}

// AddNode adds a given node to the metric path
// For example NewMetric("first").AddNode("second")
// builds metric with path equal to "first.second"
func (m *Metric) AddNode(node string) *Metric {
	m.path.Write(nextNode)
	m.path.WriteString(node)
	return m
}

// AddTag puts given tag=value pair into metric tags
// For example NewMetric("first").AddNode("second").AddTag("tag1","tag1Value")
// builds metric with path "first.second" and tags ";tag1=tag1Value"
func (m *Metric) AddTag(name, value string) *Metric {
	m.tags.Write(nextTag)
	m.tags.WriteString(name)
	m.tags.Write(nextTagValue)
	m.tags.WriteString(value)
	return m
}

// Bytes returns byte representation of Graphite metric
func (m *Metric) Bytes() []byte {
	if m.tags.Len() == 0 {
		return m.path.Bytes()
	}
	return join(m.path.Bytes(), m.tags.Bytes())
}

// GetName returns full metric name
func (m *Metric) GetName() string {
	return m.path.String()
}

// NewDatapoint returns Graphite Plaintext record form of datapoint
// Datapoint is a value stored at a timestamp bucket
// If no value is recorded at a particular timestamp bucket in a series,
// the value will be None (null)
// Doc: https://graphite.readthedocs.io/en/latest/terminology.html
func NewDatapoint(metric *Metric, value float64, timestamp int64) []byte {
	var (
		valueStr     = strconv.FormatFloat(value, 'f', -1, 64)
		timestampStr = strconv.FormatInt(timestamp, 10)
	)
	return join(metric.Bytes(), nextPart, []byte(valueStr), nextPart, []byte(timestampStr), nextDatapoint)
}

func join(s ...[]byte) []byte {
	n := 0
	for _, v := range s {
		n += len(v)
	}
	b, i := make([]byte, n), 0
	for _, v := range s {
		i += copy(b[i:], v)
	}
	return b
}
