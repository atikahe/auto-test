package promptgen

import (
	"errors"
	"io/ioutil"
	"path"
)

type InputOps string

const (
	Add InputOps = "add"
	Edit
)

type Customized struct {
	Text      string
	Overwrite bool
}

// var functionKeywordMap map[string]string = map[string]string{
// 	".go":   "func ",
// 	".py":   "def test_",
// 	".js":   "test",
// 	".java": "@Test",
// 	".cs":   "[TestMethod]",
// }

var prefixKeywordMap map[string]string = map[string]string{
	".go":   "package ",
	".py":   "import unittest",
	".js":   "const ",
	".java": "import org.junit.Test;",
	".cs":   "[TestMethod]",
}

var Default string = "\n// Here's the unit test of the code.\n// Each function will have its own test function.\n"

// The test function will implement table-based test covering various test cases and assertions.

func Build(filepath string, customPrompt ...Customized) (string, error) {
	// Validate
	ext := path.Ext(filepath)
	if _, ok := prefixKeywordMap[ext]; !ok {
		return "", errors.New("file type not supported")
	}

	// Read contents
	codeText, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", errors.New(err.Error())
	}

	// Build prompt
	prompt := string(codeText)

	if len(customPrompt) > 0 && customPrompt[0].Text != "" {
		if customPrompt[0].Overwrite {
			prompt += string(customPrompt[0].Text)
		} else {
			prompt = prompt + Default + "\n// " + string(customPrompt[0].Text) + prefixKeywordMap[ext]
		}
	}

	if len(customPrompt) == 0 {
		prompt += Default + prefixKeywordMap[ext]
	}

	return prompt, nil
}

func AddPrefix(ext string, text string) string {
	return prefixKeywordMap[ext] + text
}
