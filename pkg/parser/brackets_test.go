package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrackets(t *testing.T) {
	tests := []struct {
		name                    string
		input                   string
		beginsWithRoundBracket  bool
		beginsWithSquareBracket bool
		endsWithRoundBracket    bool
		endsWithSquareBracket   bool
	}{
		{
			name:                    "starts and end with round brackets",
			input:                   "Attendance(venue=Jewel Airport, headcount=1023)",
			beginsWithRoundBracket:  true,
			beginsWithSquareBracket: false,
			endsWithRoundBracket:    true,
			endsWithSquareBracket:   false,
		},
		{
			name:                    "starts with round brackets",
			input:                   "Attendance(venue=",
			beginsWithRoundBracket:  true,
			beginsWithSquareBracket: false,
			endsWithRoundBracket:    false,
			endsWithSquareBracket:   false,
		},
		{
			name:                    "ends with round bracket",
			input:                   "createdAt=2022-03-04T09:08:21.052Z, shipperDescription=null)",
			beginsWithRoundBracket:  false,
			beginsWithSquareBracket: false,
			endsWithRoundBracket:    true,
			endsWithSquareBracket:   false,
		},
		{
			name:                    "ends with round bracket and comma",
			input:                   "createdAt=2022-03-04T09:08:21.052Z, shipperDescription=null),",
			beginsWithRoundBracket:  false,
			beginsWithSquareBracket: false,
			endsWithRoundBracket:    true,
			endsWithSquareBracket:   false,
		},
		{
			name:                    "start with square bracket",
			input:                   "[",
			beginsWithRoundBracket:  false,
			beginsWithSquareBracket: true,
			endsWithRoundBracket:    false,
			endsWithSquareBracket:   false,
		},
		{
			name:                    "end with square bracket",
			input:                   "]",
			beginsWithRoundBracket:  false,
			beginsWithSquareBracket: false,
			endsWithRoundBracket:    false,
			endsWithSquareBracket:   true,
		},
		{
			name:                    "starts and ends with square brackets",
			input:                   "[1, 1, 2]",
			beginsWithRoundBracket:  false,
			beginsWithSquareBracket: true,
			endsWithRoundBracket:    false,
			endsWithSquareBracket:   true,
		},
		{
			name:                    "ending with string",
			input:                   "],sstr",
			beginsWithRoundBracket:  false,
			beginsWithSquareBracket: false,
			endsWithRoundBracket:    false,
			endsWithSquareBracket:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.beginsWithRoundBracket, beginsWithBracket(test.input, openRound))
			assert.Equal(t, test.beginsWithSquareBracket, beginsWithBracket(test.input, openSquare))
			assert.Equal(t, test.endsWithRoundBracket, endsWithRoundBracket(test.input))
			assert.Equal(t, test.endsWithSquareBracket, endsWithSquareBracket(test.input))
		})
	}
}
