package builtin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/macbinn/hacklang/value"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
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

func hmacSum(key, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func hmacGenrate(key, text string) string {
	bs := hmacSum([]byte(key), []byte(text))
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
	bs := hmacSum([]byte(key), []byte(text))
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
	text := fmt.Sprintf("%d%s", t, data)
	hash := hmacGenrate(key, text)
	s := fmt.Sprintf("%s:%d:%s", data, t, hash)
	return NewString(s)
}

func TicketGetData(args...value.Object) value.Object {
	ticket := args[0].(*String).S
	key := args[1].(*String).S
	parts := strings.Split(ticket, ":")
	if len(parts) != 3 {
		return nil
	}
	data := parts[0]
	t, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil
	}
	if time.Now().Unix() > int64(t) {
		return nil
	}
	hash := parts[2]
	if hmacCompare(hash, key, parts[1] + data) {
		return NewString(data)
	}
	return nil
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
			"getData": NewFunction("ticket.getData", TicketGetData),
		}),
	}))
}