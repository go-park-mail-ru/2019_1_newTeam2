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

func easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels(in *jlexer.Lexer, out *ParametersId) {
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
		case "id":
			out.ID = int(in.Int())
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
func easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels(out *jwriter.Writer, in ParametersId) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ParametersId) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ParametersId) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ParametersId) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ParametersId) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels(l, v)
}
func easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels1(in *jlexer.Lexer, out *DictionaryInfoPrivilege) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "userId":
			out.UserId = int(in.Int())
		case "privilege":
			out.Privilege = bool(in.Bool())
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
func easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels1(out *jwriter.Writer, in DictionaryInfoPrivilege) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.UserId))
	}
	if in.Privilege {
		const prefix string = ",\"privilege\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Privilege))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DictionaryInfoPrivilege) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DictionaryInfoPrivilege) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DictionaryInfoPrivilege) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DictionaryInfoPrivilege) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels1(l, v)
}
func easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels2(in *jlexer.Lexer, out *DictionaryInfo) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "userId":
			out.UserId = int(in.Int())
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
func easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels2(out *jwriter.Writer, in DictionaryInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.UserId != 0 {
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DictionaryInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DictionaryInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DictionaryInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DictionaryInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels2(l, v)
}
func easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels3(in *jlexer.Lexer, out *DictInfos) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(DictInfos, 0, 1)
			} else {
				*out = DictInfos{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 DictionaryInfo
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels3(out *jwriter.Writer, in DictInfos) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v DictInfos) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DictInfos) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DictInfos) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DictInfos) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels3(l, v)
}
func easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels4(in *jlexer.Lexer, out *CreateDictionary) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]Card, 0, 2)
					} else {
						out.Cards = []Card{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Card
					(v4).UnmarshalEasyJSON(in)
					out.Cards = append(out.Cards, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels4(out *jwriter.Writer, in CreateDictionary) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if len(in.Cards) != 0 {
		const prefix string = ",\"cards\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Cards {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateDictionary) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateDictionary) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson58209a50EncodeGithubComUser20191NewTeam2SharedModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateDictionary) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateDictionary) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson58209a50DecodeGithubComUser20191NewTeam2SharedModels4(l, v)
}
