package utils

import "crypto/rand"

func GeneratePeerId() ([]byte, error) {
	peerId := make([]byte, 20)
	copy(peerId[:8], []byte("-GT0000-"))
	_, err := rand.Read(peerId[8:])
	if err != nil {
		return nil, err
	}

	return peerId, nil
}
