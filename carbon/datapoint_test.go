package carbon

import "testing"

func Test_NewDatapoint(t *testing.T) {
	const (
		testName             = "name"
		testPrefixParentNode = "parent"
		testPrefixChildNode  = "child"
		testFirstTagName     = "firstTag"
		testFirstTagValue    = "firstTagValue"
		testSecondTagName    = "secondTag"
		testSecondTagValue   = "secondTagValue"
		testValue            = 3.14159265359
		testTimestamp        = 1552503600
	)

	t.Run("simple", func(t *testing.T) {
		var (
			actual   = NewDatapoint(testName, testValue, testTimestamp).Plaintext()
			expected = "name 3.14159265359 1552503600\n"
		)
		if actual != expected {
			t.Fatalf("expected: %s\nactual: %s", expected, actual)
		}
	})

	t.Run("with prefix", func(t *testing.T) {
		var (
			actual = NewDatapoint(testName, testValue, testTimestamp).
				WithPrefix(testPrefixParentNode).
				Plaintext()
			expected = "parent.name 3.14159265359 1552503600\n"
		)
		if actual != expected {
			t.Fatalf("expected: %s\nactual: %s", expected, actual)
		}
	})

	t.Run("with nested prefix", func(t *testing.T) {
		var (
			actual = NewDatapoint(testName, testValue, testTimestamp).
				WithPrefix(testPrefixChildNode).
				WithPrefix(testPrefixParentNode).
				Plaintext()
			expected = "parent.child.name 3.14159265359 1552503600\n"
		)
		if actual != expected {
			t.Fatalf("expected: %s\nactual: %s", expected, actual)
		}
	})

	t.Run("with nested prefix and tag", func(t *testing.T) {
		var (
			recordForms = []string{
				NewDatapoint(testName, testValue, testTimestamp).
					WithPrefix(testPrefixChildNode).
					WithPrefix(testPrefixParentNode).
					WithTag(testFirstTagName, testFirstTagValue).
					Plaintext(),
				NewDatapoint(testName, testValue, testTimestamp).
					WithTag(testFirstTagName, testFirstTagValue).
					WithPrefix(testPrefixChildNode).
					WithPrefix(testPrefixParentNode).
					Plaintext(),
			}
			expected = "parent.child.name;firstTag=firstTagValue 3.14159265359 1552503600\n"
		)
		for _, actual := range recordForms {
			if actual != expected {
				t.Fatalf("expected: %s\nactual: %s", expected, actual)
			}
		}
	})

	t.Run("with nested prefix and tags", func(t *testing.T) {
		var (
			recordForms = []string{
				NewDatapoint(testName, testValue, testTimestamp).
					WithPrefix(testPrefixChildNode).
					WithPrefix(testPrefixParentNode).
					WithTag(testFirstTagName, testFirstTagValue).
					WithTag(testSecondTagName, testSecondTagValue).
					Plaintext(),
				NewDatapoint(testName, testValue, testTimestamp).
					WithTag(testFirstTagName, testFirstTagValue).
					WithTag(testSecondTagName, testSecondTagValue).
					WithPrefix(testPrefixChildNode).
					WithPrefix(testPrefixParentNode).
					Plaintext(),
			}
			expected = "parent.child.name;firstTag=firstTagValue;secondTag=secondTagValue 3.14159265359 1552503600\n"
		)
		for _, actual := range recordForms {
			if actual != expected {
				t.Fatalf("expected: %s\nactual: %s", expected, actual)
			}
		}
	})
}
