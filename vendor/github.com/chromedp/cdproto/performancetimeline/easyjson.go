// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package performancetimeline

import (
	json "encoding/json"
	cdp "github.com/chromedp/cdproto/cdp"
	dom "github.com/chromedp/cdproto/dom"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline(in *jlexer.Lexer, out *TimelineEvent) {
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
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "type":
			out.Type = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				if out.Time == nil {
					out.Time = new(cdp.TimeSinceEpoch)
				}
				(*out.Time).UnmarshalEasyJSON(in)
			}
		case "duration":
			out.Duration = float64(in.Float64())
		case "lcpDetails":
			if in.IsNull() {
				in.Skip()
				out.LcpDetails = nil
			} else {
				if out.LcpDetails == nil {
					out.LcpDetails = new(LargestContentfulPaint)
				}
				(*out.LcpDetails).UnmarshalEasyJSON(in)
			}
		case "layoutShiftDetails":
			if in.IsNull() {
				in.Skip()
				out.LayoutShiftDetails = nil
			} else {
				if out.LayoutShiftDetails == nil {
					out.LayoutShiftDetails = new(LayoutShift)
				}
				(*out.LayoutShiftDetails).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline(out *jwriter.Writer, in TimelineEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"frameId\":"
		out.RawString(prefix[1:])
		out.String(string(in.FrameID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix)
		if in.Time == nil {
			out.RawString("null")
		} else {
			(*in.Time).MarshalEasyJSON(out)
		}
	}
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		out.RawString(prefix)
		out.Float64(float64(in.Duration))
	}
	if in.LcpDetails != nil {
		const prefix string = ",\"lcpDetails\":"
		out.RawString(prefix)
		(*in.LcpDetails).MarshalEasyJSON(out)
	}
	if in.LayoutShiftDetails != nil {
		const prefix string = ",\"layoutShiftDetails\":"
		out.RawString(prefix)
		(*in.LayoutShiftDetails).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TimelineEvent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TimelineEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TimelineEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TimelineEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline1(in *jlexer.Lexer, out *LayoutShiftAttribution) {
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
		case "previousRect":
			if in.IsNull() {
				in.Skip()
				out.PreviousRect = nil
			} else {
				if out.PreviousRect == nil {
					out.PreviousRect = new(dom.Rect)
				}
				(*out.PreviousRect).UnmarshalEasyJSON(in)
			}
		case "currentRect":
			if in.IsNull() {
				in.Skip()
				out.CurrentRect = nil
			} else {
				if out.CurrentRect == nil {
					out.CurrentRect = new(dom.Rect)
				}
				(*out.CurrentRect).UnmarshalEasyJSON(in)
			}
		case "nodeId":
			(out.NodeID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline1(out *jwriter.Writer, in LayoutShiftAttribution) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"previousRect\":"
		out.RawString(prefix[1:])
		if in.PreviousRect == nil {
			out.RawString("null")
		} else {
			(*in.PreviousRect).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"currentRect\":"
		out.RawString(prefix)
		if in.CurrentRect == nil {
			out.RawString("null")
		} else {
			(*in.CurrentRect).MarshalEasyJSON(out)
		}
	}
	if in.NodeID != 0 {
		const prefix string = ",\"nodeId\":"
		out.RawString(prefix)
		out.Int64(int64(in.NodeID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LayoutShiftAttribution) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LayoutShiftAttribution) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LayoutShiftAttribution) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LayoutShiftAttribution) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline2(in *jlexer.Lexer, out *LayoutShift) {
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
		case "value":
			out.Value = float64(in.Float64())
		case "hadRecentInput":
			out.HadRecentInput = bool(in.Bool())
		case "lastInputTime":
			if in.IsNull() {
				in.Skip()
				out.LastInputTime = nil
			} else {
				if out.LastInputTime == nil {
					out.LastInputTime = new(cdp.TimeSinceEpoch)
				}
				(*out.LastInputTime).UnmarshalEasyJSON(in)
			}
		case "sources":
			if in.IsNull() {
				in.Skip()
				out.Sources = nil
			} else {
				in.Delim('[')
				if out.Sources == nil {
					if !in.IsDelim(']') {
						out.Sources = make([]*LayoutShiftAttribution, 0, 8)
					} else {
						out.Sources = []*LayoutShiftAttribution{}
					}
				} else {
					out.Sources = (out.Sources)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *LayoutShiftAttribution
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(LayoutShiftAttribution)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Sources = append(out.Sources, v1)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline2(out *jwriter.Writer, in LayoutShift) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Value))
	}
	{
		const prefix string = ",\"hadRecentInput\":"
		out.RawString(prefix)
		out.Bool(bool(in.HadRecentInput))
	}
	{
		const prefix string = ",\"lastInputTime\":"
		out.RawString(prefix)
		if in.LastInputTime == nil {
			out.RawString("null")
		} else {
			(*in.LastInputTime).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"sources\":"
		out.RawString(prefix)
		if in.Sources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Sources {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LayoutShift) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LayoutShift) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LayoutShift) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LayoutShift) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline3(in *jlexer.Lexer, out *LargestContentfulPaint) {
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
		case "renderTime":
			if in.IsNull() {
				in.Skip()
				out.RenderTime = nil
			} else {
				if out.RenderTime == nil {
					out.RenderTime = new(cdp.TimeSinceEpoch)
				}
				(*out.RenderTime).UnmarshalEasyJSON(in)
			}
		case "loadTime":
			if in.IsNull() {
				in.Skip()
				out.LoadTime = nil
			} else {
				if out.LoadTime == nil {
					out.LoadTime = new(cdp.TimeSinceEpoch)
				}
				(*out.LoadTime).UnmarshalEasyJSON(in)
			}
		case "size":
			out.Size = float64(in.Float64())
		case "elementId":
			out.ElementID = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "nodeId":
			(out.NodeID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline3(out *jwriter.Writer, in LargestContentfulPaint) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"renderTime\":"
		out.RawString(prefix[1:])
		if in.RenderTime == nil {
			out.RawString("null")
		} else {
			(*in.RenderTime).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"loadTime\":"
		out.RawString(prefix)
		if in.LoadTime == nil {
			out.RawString("null")
		} else {
			(*in.LoadTime).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.Float64(float64(in.Size))
	}
	if in.ElementID != "" {
		const prefix string = ",\"elementId\":"
		out.RawString(prefix)
		out.String(string(in.ElementID))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	if in.NodeID != 0 {
		const prefix string = ",\"nodeId\":"
		out.RawString(prefix)
		out.Int64(int64(in.NodeID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LargestContentfulPaint) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LargestContentfulPaint) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LargestContentfulPaint) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LargestContentfulPaint) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline4(in *jlexer.Lexer, out *EventTimelineEventAdded) {
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
		case "event":
			if in.IsNull() {
				in.Skip()
				out.Event = nil
			} else {
				if out.Event == nil {
					out.Event = new(TimelineEvent)
				}
				(*out.Event).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline4(out *jwriter.Writer, in EventTimelineEventAdded) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix[1:])
		if in.Event == nil {
			out.RawString("null")
		} else {
			(*in.Event).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventTimelineEventAdded) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventTimelineEventAdded) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventTimelineEventAdded) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventTimelineEventAdded) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline5(in *jlexer.Lexer, out *EnableParams) {
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
		case "eventTypes":
			if in.IsNull() {
				in.Skip()
				out.EventTypes = nil
			} else {
				in.Delim('[')
				if out.EventTypes == nil {
					if !in.IsDelim(']') {
						out.EventTypes = make([]string, 0, 4)
					} else {
						out.EventTypes = []string{}
					}
				} else {
					out.EventTypes = (out.EventTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.EventTypes = append(out.EventTypes, v4)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline5(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventTypes\":"
		out.RawString(prefix[1:])
		if in.EventTypes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.EventTypes {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPerformancetimeline5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPerformancetimeline5(l, v)
}