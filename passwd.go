package gfUtils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"math/rand"
	"strconv"
)

// PBKDF2加密
// n 为 递归加密的测试
// keyFunc func返回string格式的key
// plain 为加密的密钥
func PBKDF2Encode(plain string, n int, keyFunc func() string) (string, error) {
	// 1. 生成随机盐
	salt, err := generateSalt(32, keyFunc)
	if err != nil {
		return "", err
	}

	// 2. 随机盐与密码加密n次
	cipherText := recursionSha1([]byte(plain+salt), 0, n)

	return salt + cipherText, nil
}

// PBKDF2解密
func PBKDF2Decode(plain, cipher string, n int, keyFunc func() string) bool {
	// 1. 截取随机盐
	saltHex := cipher[:32*2]

	// 2. 随机盐解密
	salt, _ := hex.DecodeString(saltHex)
	if salt == nil {
		return false
	}

	// 3. 铭文与salt 一起sha1 n 次
	cipherText := recursionSha1([]byte(plain+saltHex), 0, n)

	return (saltHex + cipherText) == cipher
}

// 生成n位随机盐
func generateSalt(n int, keyFunc func() string) (string, error) {
	if n < 1 {
		return "", errors.New(fmt.Sprintf("numBytes argument must be a positive integer(1 or larger): %d", n))
	}

	r := rand.New(rand.NewSource(gtime.Now().UnixNano()))

	s := strconv.Itoa(r.Intn(10000)) + keyFunc()

	h := sha256.New()

	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil)), nil
}

// 递归sha1加密
func recursionSha1(b []byte, n, max int) string {
	if n >= max {
		return hex.EncodeToString(b)
	}

	h := sha1.New()
	h.Write(b)
	return recursionSha1(h.Sum(nil), n+1, max)
}


