package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5
func MD5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
