package testfile

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		filePath string
		expected string
	}{
		{"/path/to/file.java", "/path/to/fileTest.java"},
		{"/path/to/file.js", "/path/to/file.test.js"},
		{"/path/to/file.py", "/path/to/test_file.py"},
		{"/path/to/file.cs", "/path/to/fileTests.cs"},
		{"/path/to/file.go", "/path/to/file_test.go"},
	}

	for _, test := range tests {
		actual, err := Generate(test.filePath)
		if err != nil {
			t.Errorf("Generate(%s) returns error %s", test.filePath, err)
		}
		if actual != test.expected {
			t.Errorf("Generate(%s) = %s; want %s", test.filePath, actual, test.expected)
		}
	}
}
