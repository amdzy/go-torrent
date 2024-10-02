package utils

import (
	"crypto/rand"
)

func GeneratePeerId() ([20]byte, error) {
	peerId := make([]byte, 20)
	copy(peerId[:8], []byte("-GT0000-"))
	_, err := rand.Read(peerId[8:])
	if err != nil {
		return [20]byte{}, err
	}

	var peerID [20]byte
	copy(peerID[:], peerId)

	return peerID, nil
}
