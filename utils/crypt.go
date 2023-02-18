package utils

import (
	"strings"

	"github.com/google/uuid"
)

// ユニークなtokenを取得する
func CreateToken() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
