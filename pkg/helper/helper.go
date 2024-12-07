package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	goshared "github.com/willy182/goshare/shared"
)

var TimeLocal *time.Location

func init() {
	TimeLocal, _ = time.LoadLocation(TimeZoneAsia)
}

// ToAsiaJakartaTime convert time with UTC location to AsiaJakarta time
func ToAsiaJakartaTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(),
		t.Nanosecond(), TimeLocal)
}

// RecoverPanic helper
func RecoverPanic(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%v", r)
	}
	return
}

// GenerateTokenByString function for generate token by string
func GenerateTokenByString(str string) string {
	var token string

	// generate random string
	randomString := goshared.RandomString(30)

	mix := fmt.Sprintf("%s-%s", str, randomString)
	md := md5.New()
	io.WriteString(md, mix)
	sumMd5 := md.Sum(nil)
	hash := hex.EncodeToString(sumMd5[:])

	sh := sha1.New()
	io.WriteString(sh, hash)
	sumSha := sh.Sum(nil)
	token = hex.EncodeToString(sumSha[:])

	return token
}
