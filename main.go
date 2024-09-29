package ailur_pow

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

func PoW(difficulty uint64, resource string, wait int64) (string, error) {
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
			return strconv.FormatUint(difficulty, 10) + ":" + strconv.FormatInt(initialTime, 10) + ":" + hex.EncodeToString(nonce[:]) + ":" + resource + ":", nil
		}

		if wait > 0 {
			// Wait for a while before trying again
			time.Sleep(time.Duration(wait) * time.Millisecond)
		}
	}
}

func VerifyPoW(pow string) bool {
	powSplit := strings.Split(pow, ":")
	difficulty, err := strconv.ParseUint(powSplit[0], 10, 64)
	if err != nil {
		return false
	}
	timestamp, err := strconv.ParseInt(powSplit[1], 10, 64)
	if err != nil {
		return false
	}
	timestampBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(timestampBytes, uint64(timestamp))
	nonce, err := hex.DecodeString(powSplit[2])
	if err != nil {
		return false
	}
	resource := powSplit[3]
	output := hex.EncodeToString(argon2.IDKey(nonce, bytes.Join([][]byte{timestampBytes, []byte(resource)}, []byte{}), 1, 64*1024, 4, 32))
	var difficultyString strings.Builder
	for range difficulty {
		difficultyString.WriteString("0")
	}
	if strings.HasPrefix(output, difficultyString.String()) {
		return true
	} else {
		return false
	}
}
