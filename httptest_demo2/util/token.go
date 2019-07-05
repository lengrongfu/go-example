package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func Token() string {
	hash := md5.New()
	_, _ = io.WriteString(hash, strconv.FormatInt(time.Now().UnixNano()+rand.Int63(), 10))
	token := fmt.Sprintf("%x", hash.Sum(nil))
	return token
}
