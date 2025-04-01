package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type frame struct {
	flen     [3]byte
	ftype    byte
	fflag    byte
	reserved bool
	streamID [4]byte
	data     []byte
	raw      []byte
}

func (f *frame) length() uint32 {
	return binary.BigEndian.Uint32(append([]byte{0}, f.flen[:]...))
}

func (f *frame) streamId() uint32 {
	return binary.BigEndian.Uint32(f.streamID[:])
}

func (f *frame) isEmpty() bool {
	return !(len(f.raw) > 0)
}

func (f *frame) isSameStream(id [4]byte) bool {
	for i := range 4 {
		if id[i] != f.streamID[i] {
			return false
		}
	}
	return true
}

func (f *frame) stype() string {
	switch f.ftype {
	case DATA:
		return "DATA"
	case HEADERS:
		return "HEADERS"
	case PRIORITY:
		return "PRIORITY"
	case RST_STREAM:
		return "RST_STREAM"
	case SETTINGS:
		return "SETTINGS"
	case PUSH_PROMISE:
		return "PUSH_PROMISE"
	case PING:
		return "PING"
	case GOAWAY:
		return "GOAWAY"
	case WINDOW_UPDATE:
		return "WINDOW_UPDATE"
	case CONTINUATION:
		return "CONTINUATION"
	case ALTSVC:
		return "ALTSVC"
	case ORIGIN:
		return "ORIGIN"
	default:
		return "UNKNOWN"
	}
}

func (f *frame) sflag() []string {
	pad := fmt.Sprintf("PADDED: %t", f.fflag&PADDED_FMASK != 0)
	priority := fmt.Sprintf("PRIORITY: %t", f.fflag&PRIORITY_FMASK != 0)
	ends := fmt.Sprintf("END_STREAM: %t", f.fflag&END_STREAM_FMASK != 0)
	endh := fmt.Sprintf("END_HEADERS: %t", f.fflag&END_HEADERS_FMASK != 0)
	ack := fmt.Sprintf("ACK: %t", f.fflag&ACK_FMASK != 0)
	na := "(TYPE DOES NOT DEFINE ANY FLAGS)"
	switch f.ftype {
	case DATA:
		return []string{pad, ends}
	case HEADERS:
		return []string{priority, pad, endh, ends}
	case PRIORITY, RST_STREAM, GOAWAY, WINDOW_UPDATE:
		return []string{na}
	case SETTINGS, PING:
		return []string{ack}
	case PUSH_PROMISE:
		return []string{pad, endh, ends}
	case CONTINUATION:
		return []string{endh}
	default:
		return nil
	}
}

func (f *frame) String() string {
	reserved := 0
	if f.reserved {
		reserved = 1
	}
	flags := ""
	for _, s := range f.sflag() {
		flags += "\n\t" + s
	}
	formats := "Length: %d\nType: %s (%d)\nFlag: %08b (0x%x)%s\nR bit: 0b%b, Stream ID: %d [% x]\nData: [% x]\n"
	return fmt.Sprintf(formats, f.length(), f.stype(), f.ftype, f.fflag, f.fflag, flags, reserved, f.streamId(), f.streamID, f.data)
}

func (f *frame) asMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["len"] = f.length()
	m["type"] = f.stype()
	joined := ""
	flags := f.sflag()
	for i, s := range flags {
		if i == len(flags)-1 {
			joined += s
		} else {
			joined += s + ", "
		}
	}
	m["flag"] = joined
	m["reserved"] = f.reserved
	m["streamId"] = f.streamId()
	m["data"] = string(f.data)
	return m
}

func newFrame(raw []byte) (*frame, int, error) {
	if len(raw) == 0 {
		return nil, -1, errors.New("frame is empty")
	}

	if len(raw) < 9 {
		return nil, -1, errors.New("frame headers do not fit: len(raw) < 9B")
	}

	f := &frame{}

	copy(f.flen[:], raw[:3])
	nextFrameIndex := int(9 + f.length())

	f.raw = raw[:nextFrameIndex]

	f.ftype = f.raw[3]
	f.fflag = f.raw[4]

	f.streamID[0] = f.raw[5] & RESERVED_MASK
	f.reserved = f.raw[5]&(^RESERVED_MASK) != 0
	copy(f.streamID[1:], f.raw[6:9])

	if f.length() > uint32(len(f.raw[9:])) {
		return nil, -1, errors.New("frame data does not fit: flen > len(raw[9:])")
	}

	f.data = f.raw[9:]

	if len(raw) == nextFrameIndex {
		nextFrameIndex = -1
	}

	return f, nextFrameIndex, nil
}

func getFrames(b []byte) map[int]*frame {
	frames := make(map[int]*frame)
	nextFrameIndex := 0
	for {
		f, n, err := newFrame(b[nextFrameIndex:])
		if err != nil {
			fmt.Println("error:", err.Error())
			break
		} else {
			frames[nextFrameIndex] = f
			if n != -1 {
				nextFrameIndex += n
			} else {
				break
			}
		}
	}
	return frames
}
