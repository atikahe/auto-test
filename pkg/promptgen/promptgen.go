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

var prefixKeywordMap map[string]string = map[string]string{
	".go":   "func Test",
	".py":   "def test_",
	".js":   "test(",
	".java": "@Test",
	".cs":   "[TestMethod]",
}

var Default string = "\n// Generate unit test for each functions.\n"

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
	useCustomPrompt := len(customPrompt) > 0 && customPrompt[0].Text != ""

	if useCustomPrompt && customPrompt[0].Overwrite {
		prompt += string(customPrompt[0].Text)
	} else if useCustomPrompt {
		prompt = prompt + Default + "// " + string(customPrompt[0].Text) + "\n\n" + prefixKeywordMap[ext]
	} else {
		prompt += Default + "\n" + prefixKeywordMap[ext]
	}

	if len(customPrompt) == 0 {
		prompt += Default + prefixKeywordMap[ext]
	}

	return prompt, nil
}

func AddPrefix(ext string, text string) string {
	return prefixKeywordMap[ext] + text
}
