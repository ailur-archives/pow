package library

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"crypto/rand"
	"encoding/binary"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
)

func PoW(difficulty uint64, resource string) (string, error) {
	for {
		initialTime := time.Now().Unix()
		var timestamp [8]byte
		binary.LittleEndian.PutUint64(timestamp[:], uint64(initialTime))

		var nonce [16]byte
		_, err := rand.Read(nonce[:])
		if err != nil {
			return "", err
		}

		output := hex.EncodeToString(argon2.IDKey(nonce[:], bytes.Join([][]byte{timestamp[:], []byte(resource)}, []byte{}), 1, 64*1024, 4, 32))
		var difficultyString strings.Builder
		for range difficulty {
			difficultyString.WriteString("0")
		}
		if strings.HasPrefix(output, difficultyString.String()) {
			return strconv.FormatUint(difficulty, 10) + ":" + strconv.FormatInt(initialTime, 10) + ":" + hex.EncodeToString(nonce[:]) + ":" + output, nil
		}
	}
}
