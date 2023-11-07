// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6615c02eDecodeCoffeeChoosePkgConfig(in *jlexer.Lexer, out *ServerConfig) {
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
		case "appName":
			out.AppName = string(in.String())
		case "appVersion":
			out.AppVersion = string(in.String())
		case "port":
			out.Port = string(in.String())
		case "pprofPort":
			out.PprofPort = string(in.String())
		case "mode":
			out.Mode = string(in.String())
		case "jwtSecretKey":
			out.JwtSecretKey = string(in.String())
		case "cookieName":
			out.CookieName = string(in.String())
		case "readTimeout":
			out.ReadTimeout = time.Duration(in.Int64())
		case "writeTimeout":
			out.WriteTimeout = time.Duration(in.Int64())
		case "ssl":
			out.SSL = bool(in.Bool())
		case "ctxDefaultTimeout":
			out.CtxDefaultTimeout = time.Duration(in.Int64())
		case "csrf":
			out.CSRF = bool(in.Bool())
		case "debug":
			out.Debug = bool(in.Bool())
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
func easyjson6615c02eEncodeCoffeeChoosePkgConfig(out *jwriter.Writer, in ServerConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"appName\":"
		out.RawString(prefix[1:])
		out.String(string(in.AppName))
	}
	{
		const prefix string = ",\"appVersion\":"
		out.RawString(prefix)
		out.String(string(in.AppVersion))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.String(string(in.Port))
	}
	{
		const prefix string = ",\"pprofPort\":"
		out.RawString(prefix)
		out.String(string(in.PprofPort))
	}
	{
		const prefix string = ",\"mode\":"
		out.RawString(prefix)
		out.String(string(in.Mode))
	}
	{
		const prefix string = ",\"jwtSecretKey\":"
		out.RawString(prefix)
		out.String(string(in.JwtSecretKey))
	}
	{
		const prefix string = ",\"cookieName\":"
		out.RawString(prefix)
		out.String(string(in.CookieName))
	}
	{
		const prefix string = ",\"readTimeout\":"
		out.RawString(prefix)
		out.Int64(int64(in.ReadTimeout))
	}
	{
		const prefix string = ",\"writeTimeout\":"
		out.RawString(prefix)
		out.Int64(int64(in.WriteTimeout))
	}
	{
		const prefix string = ",\"ssl\":"
		out.RawString(prefix)
		out.Bool(bool(in.SSL))
	}
	{
		const prefix string = ",\"ctxDefaultTimeout\":"
		out.RawString(prefix)
		out.Int64(int64(in.CtxDefaultTimeout))
	}
	{
		const prefix string = ",\"csrf\":"
		out.RawString(prefix)
		out.Bool(bool(in.CSRF))
	}
	{
		const prefix string = ",\"debug\":"
		out.RawString(prefix)
		out.Bool(bool(in.Debug))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServerConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6615c02eEncodeCoffeeChoosePkgConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServerConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6615c02eEncodeCoffeeChoosePkgConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServerConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6615c02eDecodeCoffeeChoosePkgConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServerConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6615c02eDecodeCoffeeChoosePkgConfig(l, v)
}
func easyjson6615c02eDecodeCoffeeChoosePkgConfig1(in *jlexer.Lexer, out *Config) {
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
		case "server":
			if in.IsNull() {
				in.Skip()
				out.Server = nil
			} else {
				if out.Server == nil {
					out.Server = new(ServerConfig)
				}
				(*out.Server).UnmarshalEasyJSON(in)
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
func easyjson6615c02eEncodeCoffeeChoosePkgConfig1(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"server\":"
		out.RawString(prefix[1:])
		if in.Server == nil {
			out.RawString("null")
		} else {
			(*in.Server).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6615c02eEncodeCoffeeChoosePkgConfig1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6615c02eEncodeCoffeeChoosePkgConfig1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6615c02eDecodeCoffeeChoosePkgConfig1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6615c02eDecodeCoffeeChoosePkgConfig1(l, v)
}
