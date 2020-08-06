package cfg

import (
	"reflect"
	"testing"
)

func TestNextToken(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []Token
	}{
		{
			name: "Simple tokens",
			args: args{
				input: `#$@=;`,
			},
			want: []Token{
				{Type: HASH, Literal: "#"},
				{Type: DOLLAR, Literal: "$"},
				{Type: AT, Literal: "@"},
				{Type: ASSIGN, Literal: "="},
				{Type: SEMICOLON, Literal: ";"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name: "Snippet 1",
			args: args{
				input: `f32 sprite_offset_x                        = 0
				f32 sprite_offset_y                        = 8
				
					$sprite_gibs_start                     = *start*
				
					$gib_type                              = predefined
					$gib_style                             = wooden`,
			},
			want: []Token{
				{Type: TEXT, Literal: "f32"},
				{Type: TEXT, Literal: "sprite_offset_x"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "0"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: TEXT, Literal: "f32"},
				{Type: TEXT, Literal: "sprite_offset_y"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "8"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: DOLLAR, Literal: "$"},
				{Type: TEXT, Literal: "sprite_gibs_start"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "*start*"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: DOLLAR, Literal: "$"},
				{Type: TEXT, Literal: "gib_type"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "predefined"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: DOLLAR, Literal: "$"},
				{Type: TEXT, Literal: "gib_style"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "wooden"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name: "Snippet 2",
			args: args{
				input: `@$sprite_scripts                           = SeatsGUI.as;
				VehicleGUI.as;
				Wooden.as;`,
			},
			want: []Token{
				{Type: AT, Literal: "@"},
				{Type: DOLLAR, Literal: "$"},
				{Type: TEXT, Literal: "sprite_scripts"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "SeatsGUI.as"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: TEXT, Literal: "VehicleGUI.as"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: NEWLINE, Literal: "\n"},
				{Type: TEXT, Literal: "Wooden.as"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name: "Snippet 3",
			args: args{
				input: `@u8 gib_frame                          = 4; 5; 6; 7;`,
			},
			want: []Token{
				{Type: AT, Literal: "@"},
				{Type: TEXT, Literal: "u8"},
				{Type: TEXT, Literal: "gib_frame"},
				{Type: ASSIGN, Literal: "="},
				{Type: TEXT, Literal: "4"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: TEXT, Literal: "5"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: TEXT, Literal: "6"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: TEXT, Literal: "7"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: EOF, Literal: ""},
			},
		},
	}
	for i, tt := range tests {
		got := make([]Token, 0)
		l := NewLexer(tt.args.input)
		for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
			got = append(got, tok)
		}
		got = append(got, Token{Type: EOF, Literal: ""})

		if len(got) != len(tt.want) {
			t.Fatalf("tests[%d - %q] - wrong token count.\nWant %d tokens:\t%s,\nGot %d tokens\t%s", i,
				tt.name, len(tt.want), tt.want, len(got), got)
		}
		for j := range got {
			if !reflect.DeepEqual(got[j], tt.want[j]) {
				t.Fatalf("tests[%d - %q] wrong. want[%d]=%q, got[%d]=%q", i,
					tt.name, j, tt.want[j], j, got[j])
			}
		}
	}
}
