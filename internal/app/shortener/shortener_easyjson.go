// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package shortener

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

func easyjsonC6c4a0a8DecodeGithubComBobopylabepolhkYpshortenerInternalAppShortener(in *jlexer.Lexer, out *ShortenURLResponseDTO) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			out.Result = string(in.String())
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
func easyjsonC6c4a0a8EncodeGithubComBobopylabepolhkYpshortenerInternalAppShortener(out *jwriter.Writer, in ShortenURLResponseDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix[1:])
		out.String(string(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ShortenURLResponseDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC6c4a0a8EncodeGithubComBobopylabepolhkYpshortenerInternalAppShortener(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ShortenURLResponseDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC6c4a0a8EncodeGithubComBobopylabepolhkYpshortenerInternalAppShortener(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ShortenURLResponseDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC6c4a0a8DecodeGithubComBobopylabepolhkYpshortenerInternalAppShortener(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ShortenURLResponseDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC6c4a0a8DecodeGithubComBobopylabepolhkYpshortenerInternalAppShortener(l, v)
}
func easyjsonC6c4a0a8DecodeGithubComBobopylabepolhkYpshortenerInternalAppShortener1(in *jlexer.Lexer, out *ShortenURLRequestDTO) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.URL = string(in.String())
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
func easyjsonC6c4a0a8EncodeGithubComBobopylabepolhkYpshortenerInternalAppShortener1(out *jwriter.Writer, in ShortenURLRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ShortenURLRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC6c4a0a8EncodeGithubComBobopylabepolhkYpshortenerInternalAppShortener1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ShortenURLRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC6c4a0a8EncodeGithubComBobopylabepolhkYpshortenerInternalAppShortener1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ShortenURLRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC6c4a0a8DecodeGithubComBobopylabepolhkYpshortenerInternalAppShortener1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ShortenURLRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC6c4a0a8DecodeGithubComBobopylabepolhkYpshortenerInternalAppShortener1(l, v)
}