package lexer

import (
	"reflect"
	"testing"

	"github.com/KAG-Apparatus/blob-editor/token"
)

func TestNextToken(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []token.Token
	}{
		{
			name: "Simple tokens",
			args: args{
				input: `#$@=;`,
			},
			want: []token.Token{
				{Type: token.HASH, Literal: "#"},
				{Type: token.DOLLAR, Literal: "$"},
				{Type: token.AT, Literal: "@"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
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
			want: []token.Token{
				{Type: token.TEXT, Literal: "f32"},
				{Type: token.TEXT, Literal: "sprite_offset_x"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "0"},
				{Type: token.TEXT, Literal: "f32"},
				{Type: token.TEXT, Literal: "sprite_offset_y"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "8"},
				{Type: token.DOLLAR, Literal: "$"},
				{Type: token.TEXT, Literal: "sprite_gibs_start"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "*start*"},
				{Type: token.DOLLAR, Literal: "$"},
				{Type: token.TEXT, Literal: "gib_type"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "predefined"},
				{Type: token.DOLLAR, Literal: "$"},
				{Type: token.TEXT, Literal: "gib_style"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "wooden"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "Snippet 2",
			args: args{
				input: `@$sprite_scripts                           = SeatsGUI.as;
				VehicleGUI.as;
				Wooden.as;`,
			},
			want: []token.Token{
				{Type: token.AT, Literal: "@"},
				{Type: token.DOLLAR, Literal: "$"},
				{Type: token.TEXT, Literal: "sprite_scripts"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "SeatsGUI.as"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.TEXT, Literal: "VehicleGUI.as"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.TEXT, Literal: "Wooden.as"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "Snippet 3",
			args: args{
				input: `@u8 gib_frame                          = 4; 5; 6; 7;`,
			},
			want: []token.Token{
				{Type: token.AT, Literal: "@"},
				{Type: token.TEXT, Literal: "u8"},
				{Type: token.TEXT, Literal: "gib_frame"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.TEXT, Literal: "4"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.TEXT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.TEXT, Literal: "6"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.TEXT, Literal: "7"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}
	for i, tt := range tests {
		got := make([]token.Token, 0)
		l := New(tt.args.input)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			got = append(got, tok)
		}
		got = append(got, token.Token{Type: token.EOF, Literal: ""})

		if len(got) != len(tt.want) {
			t.Fatalf("tests[%d - %q] - tokentype wrong.\nWant\t%s,\nGot\t%s", i,
				tt.name, tt.want, got)
		}
		for j := range got {
			if !reflect.DeepEqual(got[j], tt.want[j]) {
				t.Fatalf("tests[%d - %q] wrong. want[%d]=%q, got[%d]=%q", i,
					tt.name, j, tt.want[j], j, got[j])
			}
		}
	}
}
