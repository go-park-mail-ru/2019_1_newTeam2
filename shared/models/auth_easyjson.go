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

func easyjson4a0f95aaDecodeGithubComUser20191NewTeam2SharedModels(in *jlexer.Lexer, out *UserAuth) {
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
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson4a0f95aaEncodeGithubComUser20191NewTeam2SharedModels(out *jwriter.Writer, in UserAuth) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserAuth) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a0f95aaEncodeGithubComUser20191NewTeam2SharedModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserAuth) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a0f95aaEncodeGithubComUser20191NewTeam2SharedModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserAuth) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a0f95aaDecodeGithubComUser20191NewTeam2SharedModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserAuth) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a0f95aaDecodeGithubComUser20191NewTeam2SharedModels(l, v)
}