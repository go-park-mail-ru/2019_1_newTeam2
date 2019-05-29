// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

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

func easyjson6615c02eDecodeGithubComUser20191NewTeam2SharedPkgConfig(in *jlexer.Lexer, out *Config) {
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
		case "secret":
			out.Secret = string(in.String())
		case "port":
			out.Port = string(in.String())
		case "authHost":
			out.AuthHost = string(in.String())
		case "authPort":
			out.AuthPort = string(in.String())
		case "commonDbServPort":
			out.ScorePort = string(in.String())
		case "uploadPath":
			out.UploadPath = string(in.String())
		case "avatarsPath":
			out.AvatarsPath = string(in.String())
		case "AllowedHosts":
			if in.IsNull() {
				in.Skip()
				out.AllowedHosts = nil
			} else {
				in.Delim('[')
				if out.AllowedHosts == nil {
					if !in.IsDelim(']') {
						out.AllowedHosts = make([]string, 0, 4)
					} else {
						out.AllowedHosts = []string{}
					}
				} else {
					out.AllowedHosts = (out.AllowedHosts)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.AllowedHosts = append(out.AllowedHosts, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "scoreHost":
			out.ScoreHost = string(in.String())
		case "dbName":
			out.DBName = string(in.String())
		case "dbUser":
			out.DBUser = string(in.String())
		case "dbPassUser":
			out.DBPassUser = string(in.String())
		case "dbHost":
			out.DBHost = string(in.String())
		case "roomSize":
			out.RoomSize = int(in.Int())
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
func easyjson6615c02eEncodeGithubComUser20191NewTeam2SharedPkgConfig(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Secret != "" {
		const prefix string = ",\"secret\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Secret))
	}
	{
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Port))
	}
	{
		const prefix string = ",\"authHost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AuthHost))
	}
	{
		const prefix string = ",\"authPort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AuthPort))
	}
	if in.ScorePort != "" {
		const prefix string = ",\"commonDbServPort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ScorePort))
	}
	if in.UploadPath != "" {
		const prefix string = ",\"uploadPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UploadPath))
	}
	if in.AvatarsPath != "" {
		const prefix string = ",\"avatarsPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarsPath))
	}
	if len(in.AllowedHosts) != 0 {
		const prefix string = ",\"AllowedHosts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.AllowedHosts {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"scoreHost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ScoreHost))
	}
	if in.DBName != "" {
		const prefix string = ",\"dbName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DBName))
	}
	if in.DBUser != "" {
		const prefix string = ",\"dbUser\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DBUser))
	}
	if in.DBPassUser != "" {
		const prefix string = ",\"dbPassUser\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DBPassUser))
	}
	if in.DBHost != "" {
		const prefix string = ",\"dbHost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DBHost))
	}
	if in.RoomSize != 0 {
		const prefix string = ",\"roomSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.RoomSize))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6615c02eEncodeGithubComUser20191NewTeam2SharedPkgConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6615c02eEncodeGithubComUser20191NewTeam2SharedPkgConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6615c02eDecodeGithubComUser20191NewTeam2SharedPkgConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6615c02eDecodeGithubComUser20191NewTeam2SharedPkgConfig(l, v)
}