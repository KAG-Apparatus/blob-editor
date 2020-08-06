package cfg

import (
	"reflect"
	"testing"
)

func TestParser_parseAssignment(t *testing.T) {
	type fields struct {
		lexer     *Lexer
		curToken  Token
		peekToken Token
	}
	tests := []struct {
		name    string
		fields  fields
		want    Element
		wantErr bool
	}{
		{
			name: "Assignment 1",
			fields: fields{
				lexer: NewLexer(`s32_sprite_frame_width                     = 48`),
			},
			want: &Assignment{
				LeftHand: []Token{
					{Type: TEXT, Literal: "s32_sprite_frame_width"},
				},
				RightHand: []Token{
					{Type: TEXT, Literal: "48"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.fields.lexer)
			got, err := p.parseAssignment()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseAssignment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
