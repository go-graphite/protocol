package carbon

import (
	"bytes"
	"testing"
)

func Test_NewMetric(t *testing.T) {
	testCases := []struct {
		actual            *Metric
		expectedName      string
		expectedBytes     []byte
		expectedDatapoint []byte
	}{
		{
			actual:            NewMetric("metric1"),
			expectedName:      "metric1",
			expectedBytes:     []byte("metric1"),
			expectedDatapoint: []byte("metric1 3.14159265359 1552503600\n"),
		},
		{
			actual:            NewMetric("metric1").AddNode("foo"),
			expectedName:      "metric1.foo",
			expectedBytes:     []byte("metric1.foo"),
			expectedDatapoint: []byte("metric1.foo 3.14159265359 1552503600\n"),
		},
		{
			actual:            NewMetric("metric1").AddNode("foo").AddNode("bar"),
			expectedName:      "metric1.foo.bar",
			expectedBytes:     []byte("metric1.foo.bar"),
			expectedDatapoint: []byte("metric1.foo.bar 3.14159265359 1552503600\n"),
		},
		{
			actual:            NewMetric("metric1").AddNode("foo").AddNode("bar").AddTag("cpu", "cpu1"),
			expectedName:      "metric1.foo.bar",
			expectedBytes:     []byte("metric1.foo.bar;cpu=cpu1"),
			expectedDatapoint: []byte("metric1.foo.bar;cpu=cpu1 3.14159265359 1552503600\n"),
		},
		{
			actual:            NewMetric("metric1").AddTag("cpu", "cpu1").AddNode("foo").AddNode("bar"),
			expectedName:      "metric1.foo.bar",
			expectedBytes:     []byte("metric1.foo.bar;cpu=cpu1"),
			expectedDatapoint: []byte("metric1.foo.bar;cpu=cpu1 3.14159265359 1552503600\n"),
		},
		{
			actual:            NewMetric("metric1").AddNode("foo").AddNode("bar").AddTag("cpu", "cpu1").AddTag("dc", "dc1"),
			expectedName:      "metric1.foo.bar",
			expectedBytes:     []byte("metric1.foo.bar;cpu=cpu1;dc=dc1"),
			expectedDatapoint: []byte("metric1.foo.bar;cpu=cpu1;dc=dc1 3.14159265359 1552503600\n"),
		},
		{
			actual:            NewMetric("metric1").AddTag("cpu", "cpu1").AddTag("dc", "dc1").AddNode("foo").AddNode("bar"),
			expectedName:      "metric1.foo.bar",
			expectedBytes:     []byte("metric1.foo.bar;cpu=cpu1;dc=dc1"),
			expectedDatapoint: []byte("metric1.foo.bar;cpu=cpu1;dc=dc1 3.14159265359 1552503600\n"),
		},
	}

	for _, testCase := range testCases {
		const (
			value     = 3.14159265359
			timestamp = 1552503600
		)
		var (
			actualName      = testCase.actual.GetName()
			actualBytes     = testCase.actual.Bytes()
			actualDatapoint = NewDatapoint(testCase.actual, value, timestamp)
		)
		if actualName != testCase.expectedName {
			t.Fatalf("expected name: %s\nactual name: %s", testCase.expectedName, actualName)
		}
		if bytes.Compare(actualBytes, testCase.expectedBytes) != 0 {
			t.Fatalf("expected bytes: %s\nactual bytes: %s", testCase.expectedBytes, actualBytes)
		}
		if bytes.Compare(actualDatapoint, testCase.expectedDatapoint) != 0 {
			t.Fatalf("expected datapoint: %s\nactual datapoint: %s", testCase.expectedDatapoint, actualDatapoint)
		}
	}
}

func Benchmark_NewDatapointPreCreated(b *testing.B) {
	metric := NewMetric("metric1").
		AddNode("foo").
		AddNode("bar").
		AddTag("cpu", "cpu1").
		AddTag("dc", "dc1").
		AddTag("env", "stage").
		AddTag("service", "exporter").
		AddTag("version", "v0.0.1").
		AddTag("git-commit", "0b1a09c")
	for i := 0; i < b.N; i++ {
		NewDatapoint(metric, 3.14159265359, 1552503600)
	}
}

func Benchmark_NewDatapointOnTheFly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		metric := NewMetric("metric1").
			AddNode("foo").
			AddNode("bar").
			AddTag("cpu", "cpu1").
			AddTag("dc", "dc1").
			AddTag("env", "stage").
			AddTag("service", "exporter").
			AddTag("version", "v0.0.1").
			AddTag("git-commit", "0b1a09c")
		NewDatapoint(metric, 3.14159265359, 1552503600)
	}
}
