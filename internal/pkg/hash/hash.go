package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5ByFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr), nil
}
