# 🤖 auto-test
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Automatically generate test file for your code from the CLI so you don't need to start from scratch ✨ Built with OpenAI's Codex.

## Installation
Installation guide will go here.

## Usage
- Get access to OpenAI Beta and get your API Key [here](https://beta.openai.com/account/api-keys).
```bash
export OPENAI_API_KEY=YOUR_API_KEY
```
- Go to the directory of code that you want to test. 
- Run auto-test
```bash
auto-test -f file_name.py
```
- See generated test file on the same folder, make some modification as the code won't be complete.
- To add custom prompt, use `-p` flag followed by a string of prompt.
```
auto-test -f file_name.py -p "This is additional prompt"
```
- To overwrite default prompt, use `-o` flag.
- Use `auto-test --help` for more info.

## Disclaimer
This package is not intended to replace the writing of unit-test by AI entirely. You will still need to make adjustments and mock some dependencies manually.

## Resources
- [OpenAI Codex](https://openai.com/blog/openai-codex/)