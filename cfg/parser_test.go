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
		want    *Assignment
		wantErr bool
	}{
		{
			name: "Assignment 1",
			fields: fields{
				lexer: NewLexer(`s32_sprite_frame_width                     = 48`),
			},
			want: &Assignment{
				Line:       1,
				Identifier: "s32_sprite_frame_width",
				Type:       STRING_TYPE,
				Value:      "48",
			},
			wantErr: false,
		},
		{
			name: "Assignment 2",
			fields: fields{
				lexer: NewLexer(`@u8 gib_frame                          = 4; 5; 6; 7;`),
			},
			want: &Assignment{
				Line:       1,
				Identifier: "gib_frame",
				Type:       INT_ARRAY_TYPE,
				Value:      []int{4, 5, 6, 7},
			},
			wantErr: false,
		},
		{
			name: "Assignment 3",
			fields: fields{
				lexer: NewLexer(`@f32 verticesXY                            =  0.0; 0.0;  
				40.0; 0.0; 
			 32.0; 10.0; 
			 8.0; 10.0;	`),
			},
			want: &Assignment{
				Line:       1,
				Identifier: "verticesXY",
				Type:       FLOAT_ARRAY_TYPE,
				Value:      []float64{0.0, 0.0, 40.0, 0.0, 32.0, 10.0, 8.0, 10.0},
			},
			wantErr: false,
		},
		{
			name: "Assignment 4",
			fields: fields{
				lexer: NewLexer(`@$attachment_points                        =  FLYER;   	   0;  -2;  0; 1; 7;
				VEHICLE; 		15;  -4;  0; 0; 0;		
				CARGO; 		0; 14;  0; 0; 0;		`),
			},
			want: &Assignment{
				Line:       1,
				Identifier: "attachment_points",
				Type:       STRING_ARRAY_TYPE,
				Value: []string{"FLYER", "0", "-2", "0", "1", "7",
					"VEHICLE", "15", "-4", "0", "0", "0",
					"CARGO", "0", "14", "0", "0", "0"},
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
			if !got.Equal(tt.want) {
				t.Errorf("Parser.parseAssignment() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestParser_parseComment(t *testing.T) {
	type fields struct {
		lexer     *Lexer
		curToken  Token
		peekToken Token
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Comment
		wantErr bool
	}{
		{
			name: "Comment 1",
			fields: fields{
				lexer: NewLexer(`# Boat config file`),
			},
			want: &Comment{
				Line:  1,
				Value: "# Boat config file",
			},
			wantErr: false,
		},
		{
			name: "Comment 2",
			fields: fields{
				lexer: NewLexer(`# $ string #$@=;`),
			},
			want: &Comment{
				Line:  1,
				Value: "# $ string #$@=;",
			},
			wantErr: false,
		},
		{
			name: "Comment 3",
			fields: fields{
				lexer: NewLexer(`# sprite

				$sprite_factory                            = generic_sprite`),
			},
			want: &Comment{
				Line:  1,
				Value: "# sprite",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				lexer:     tt.fields.lexer,
				curToken:  tt.fields.curToken,
				peekToken: tt.fields.peekToken,
			}
			got, err := p.parseComment()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
