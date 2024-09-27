package handshake

import (
	"fmt"
	"io"
)

type Handshake struct {
	Pstr     string
	InfoHash [20]byte
	PeerId   [20]byte
}

func New(infoHash, peerId [20]byte) *Handshake {
	return &Handshake{
		Pstr:     "BitTorrent protocol",
		InfoHash: infoHash,
		PeerId:   peerId,
	}
}

func (h *Handshake) Serialize() []byte {
	buf := make([]byte, len(h.Pstr)+49)
	buf[0] = byte(len(h.Pstr))
	pos := 1
	pos += copy(buf[pos:], []byte(h.Pstr))
	pos += copy(buf[pos:], make([]byte, 8))
	pos += copy(buf[pos:], h.InfoHash[:])
	pos += copy(buf[pos:], h.PeerId[:])

	return buf
}

func Read(r io.Reader) (*Handshake, error) {
	lengthBuf := make([]byte, 1)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}

	pstrLen := int(lengthBuf[0])
	if pstrLen == 0 {
		err := fmt.Errorf("pstrLen cannot be 0")
		return nil, err
	}

	handshakeBuf := make([]byte, 48+pstrLen)
	_, err = io.ReadFull(r, handshakeBuf)
	if err != nil {
		return nil, err
	}

	var infoHash, peerId [20]byte

	copy(infoHash[:], handshakeBuf[pstrLen+8:pstrLen+8+20])
	copy(peerId[:], handshakeBuf[pstrLen+8+20:])

	h := Handshake{
		Pstr:     string(handshakeBuf[0:pstrLen]),
		InfoHash: infoHash,
		PeerId:   peerId,
	}

	return &h, nil
}
