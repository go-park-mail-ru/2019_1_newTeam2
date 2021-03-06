// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9806e1DecodeGithubComUser20191NewTeam2SharedModels(in *jlexer.Lexer, out *DictionaryNote) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "borrowerId":
			out.BorrowerId = int(in.Int())
		case "dictionaryName":
			out.DictionaryName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9806e1EncodeGithubComUser20191NewTeam2SharedModels(out *jwriter.Writer, in DictionaryNote) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BorrowerId != 0 {
		const prefix string = ",\"borrowerId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.BorrowerId))
	}
	if in.DictionaryName != "" {
		const prefix string = ",\"dictionaryName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DictionaryName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DictionaryNote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComUser20191NewTeam2SharedModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DictionaryNote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComUser20191NewTeam2SharedModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DictionaryNote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComUser20191NewTeam2SharedModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DictionaryNote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComUser20191NewTeam2SharedModels(l, v)
}
