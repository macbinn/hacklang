package builtin

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/macbinn/hacklang/value"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func BcryptGenerate(args...value.Object) value.Object {
	s := args[0].(*String).S
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	if err != nil {
		return nil
	}
	return NewString(string(bs))
}

func BcryptCompare(args...value.Object) value.Object {
	hash := args[0].(*String).S
	pass := args[1].(*String).S
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return BoolTrue
	}
	return BoolFalse
}

func hmacGenrate(key, text string) string {
	h := hmac.New(sha1.New, []byte(key))
	bs := h.Sum([]byte(text))
	return base64.StdEncoding.EncodeToString(bs)
}

func HmacGenrate(args...value.Object) value.Object {
	key := args[0].(*String).S
	text := args[1].(*String).S
	s := hmacGenrate(key, text)
	return NewString(s)
}

func hmacCompare(hash, key, text string) bool {
	hashBytes, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	h := hmac.New(sha1.New, []byte(key))
	bs := h.Sum([]byte(text))
	return hmac.Equal(bs, hashBytes)
}

func HmacCompare(args...value.Object) value.Object {
	hash := args[0].(*String).S
	key := args[1].(*String).S
	text := args[2].(*String).S
	if hmacCompare(hash, key, text) {
		return BoolTrue
	}
	return BoolFalse
}

func TicketGenrate(args...value.Object) value.Object {
	key := args[0].(*String).S
	data := args[1].(*String).S
	expire := int64(args[2].(*Number).Int)
	t := time.Now().Unix() + expire
	text := fmt.Sprintf("%s%d", data, t)
	hash := hmacGenrate(key, text)
	s := fmt.Sprintf("%s:%d:%s", data, t, hash)
	return NewString(s)
}

func init() {
	GlobalScope.Register("hash", NewMap(map[string]value.Object{
		"bcrypt": NewMap(map[string]value.Object{
			"generate": NewFunction("bcrypt.generate", BcryptGenerate),
			"compare": NewFunction("bcrypt.compare", BcryptCompare),
		}),
		"hmac": NewMap(map[string]value.Object{
			"generate": NewFunction("hmac.generate", HmacGenrate),
			"compare": NewFunction("hmac.compare", HmacCompare),
		}),
		"ticket": NewMap(map[string]value.Object{
			"generate": NewFunction("ticket.generate", TicketGenrate),
		}),
	}))
}