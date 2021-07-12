package carbon

import "testing"

func Test_NewDatapoint(t *testing.T) {
	testCases := []struct {
		actual   string
		expected string
	}{
		{
			actual:   NewDatapoint("metric1", 3.14159265359, 1552503600).String(),
			expected: "metric1 3.14159265359 1552503600\n",
		},
		{
			actual: NewDatapoint("metric1", 3.14159265359, 1552503600).
				WithPrefix("foo").
				String(),
			expected: "foo.metric1 3.14159265359 1552503600\n",
		},
		{
			actual: NewDatapoint("metric1", 3.14159265359, 1552503600).
				WithPrefix("foo").
				WithPrefix("bar").
				String(),
			expected: "bar.foo.metric1 3.14159265359 1552503600\n",
		},
		{
			actual: NewDatapoint("metric1", 3.14159265359, 1552503600).
				WithPrefix("foo").
				WithPrefix("bar").
				WithTag("cpu", "cpu1").
				String(),
			expected: "bar.foo.metric1;cpu=cpu1 3.14159265359 1552503600\n",
		},
		{
			actual: NewDatapoint("metric1", 3.14159265359, 1552503600).
				WithTag("cpu", "cpu1").
				WithPrefix("foo").
				WithPrefix("bar").
				String(),
			expected: "bar.foo.metric1;cpu=cpu1 3.14159265359 1552503600\n",
		},
		{
			actual: NewDatapoint("metric1", 3.14159265359, 1552503600).
				WithPrefix("foo").
				WithPrefix("bar").
				WithTag("cpu", "cpu1").
				WithTag("dc", "dc1").
				String(),
			expected: "bar.foo.metric1;cpu=cpu1;dc=dc1 3.14159265359 1552503600\n",
		},
		{
			actual: NewDatapoint("metric1", 3.14159265359, 1552503600).
				WithTag("cpu", "cpu1").
				WithTag("dc", "dc1").
				WithPrefix("foo").
				WithPrefix("bar").
				String(),
			expected: "bar.foo.metric1;cpu=cpu1;dc=dc1 3.14159265359 1552503600\n",
		},
	}

	for _, testCase := range testCases {
		if testCase.actual != testCase.expected {
			t.Fatalf("expected: %s\nactual: %s", testCase.expected, testCase.actual)
		}
	}
}

func BenchmarkDatapoint_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDatapoint("metric1", 3.14159265359, 1552503600).
			WithPrefix("foo").
			WithPrefix("bar").
			WithTag("cpu", "cpu1").
			WithTag("dc", "dc1").
			WithTag("env", "stage").
			WithTag("service", "exporter").
			WithTag("version", "v0.0.1").
			WithTag("git-commit", "0b1a09c").
			String()
	}
}

func BenchmarkDatapoint_StringByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDatapoint(
			"bar.foo.metric1;cpu=cpu1;dc=dc1;env=stage;service=exporter;version=v0.0.1;git-commit=0b1a09c",
			3.14159265359,
			1552503600).
			String()
	}
}
