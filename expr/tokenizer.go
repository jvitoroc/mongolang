package expr

import (
	"log"
	"math"
	"regexp"
)

const (
	EQUAL = "EQUAL"

	GREATER_THAN       = "GREATER_THAN"
	GREATER_EQUAL_THAN = "GREATER_EQUAL_THAN"

	LESS_THAN       = "LESS_THAN"
	LESS_EQUAL_THAN = "LESS_EQUAL_THAN"

	AND = "AND"
	OR  = "OR"

	LEFT_PARENTHESIS  = "LEFT_PARENTHESIS"
	RIGHT_PARENTHESIS = "RIGHT_PARENTHESIS"

	STRING = "STRING"

	NUMBER = "NUMBER"

	FIELD = "FIELD"

	PLACEHOLDER = "PLACEHOLDER"

	INVALID = "INVALID"
)

const (
	COMPARISON_OPERATOR = "COMPARISON_OPERATOR"
	LOGICAL_OPERATOR    = "LOGICAL_OPERATOR"
	PARATHENSIS         = "PARENTHESIS"
	LITERAL             = "LITERAL"
	VARIABLE            = "VARIABLE"
)

type TokenMatcher struct {
	Kind    string
	Name    string
	Pattern string
	Regexp  *regexp.Regexp
}

type Token struct {
	Kind  string
	Name  string
	Value string
}

type Tokenizer struct {
	tokens []*TokenMatcher
}

func NewTokenizer() *Tokenizer {
	t := &Tokenizer{}

	t.addToken(STRING, LITERAL, `("[^"]*"|'[^']*')`)
	t.addToken(NUMBER, LITERAL, `\b\d+(\.\d+)?\b`)
	t.addToken(PLACEHOLDER, VARIABLE, `\b\?\b`)
	t.addToken(AND, LOGICAL_OPERATOR, `\band\b`)
	t.addToken(OR, LOGICAL_OPERATOR, `\bor\b`)
	t.addToken(GREATER_EQUAL_THAN, COMPARISON_OPERATOR, `>=`)
	t.addToken(LESS_EQUAL_THAN, COMPARISON_OPERATOR, `<=`)
	t.addToken(GREATER_THAN, COMPARISON_OPERATOR, `>`)
	t.addToken(LESS_THAN, COMPARISON_OPERATOR, `<`)
	t.addToken(EQUAL, COMPARISON_OPERATOR, `=`)
	t.addToken(LEFT_PARENTHESIS, PARATHENSIS, `\(`)
	t.addToken(RIGHT_PARENTHESIS, PARATHENSIS, `\)`)
	t.addToken(FIELD, VARIABLE, `\b\w+(\.\w+|\[(\*|\d+)\])*\b`)
	t.addToken(INVALID, "", `[^\s]+`)

	t.compileTokens()

	return t
}

func (t *Tokenizer) Tokenize(value string) []*Token {
	result := []*Token{}
	index := 0
	asRunes := []rune(value)
	runeCount := len(asRunes)

	for index < runeCount {
		match, token := t.getTokenAt(asRunes[index:], index)
		if match != "" {
			result = append(result, &Token{Kind: token.Kind, Name: token.Name, Value: match})
		}
		offset := math.Max(float64(len(match)), 1)
		index += int(offset)
	}

	return result
}

func (t *Tokenizer) addToken(name string, kind string, pattern string) {
	t.tokens = append(t.tokens, &TokenMatcher{Name: name, Kind: kind, Pattern: pattern})
}

func (t *Tokenizer) compileTokens() {
	for _, token := range t.tokens {
		token.Regexp = regexp.MustCompile("^" + token.Pattern)
	}
}

func (t *Tokenizer) getTokenAt(value []rune, index int) (string, *TokenMatcher) {
	if match, token := t.isKnownToken(value); match != "" {
		if token.Name == "INVALID" {
			log.Fatalf("invalid token \"%s\" at %d", match, index)
		}

		return match, token
	}

	return "", nil
}

func (t *Tokenizer) isKnownToken(value []rune) (string, *TokenMatcher) {
	for _, token := range t.tokens {
		if matched, match := t.isToken(value, token); matched {
			return match, token
		}
	}

	return "", nil
}

func (t *Tokenizer) isToken(value []rune, token *TokenMatcher) (bool, string) {
	target := string(value)
	if loc := token.Regexp.FindStringIndex(target); loc != nil {
		match := target[loc[0]:loc[1]]
		return true, match
	}

	return false, ""
}
