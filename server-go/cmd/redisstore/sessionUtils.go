package redisstore

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func checkIdSession(id string) (string, error) {
	query, err2 := url.QueryUnescape(id) // возможные символы escape
	if err2 != nil {
		log.Println("Ошибка ID Escape из Cookie", err2)
	}
	spl := strings.Split(query, ":")
	if len(spl) > 1 && spl[0] == "s" {
		// это ID выставлен Express Session Js
		spl = strings.Split(spl[1], ".") // разделим на код и подпись
		err := checkSignedCookie(spl[0], spl[1])
		if err != nil {
			log.Println("Нарушена подпись ключа в Cookie")
			return "BadSigned", errors.New("нарушена подпись ключа")
		}
		return spl[0], nil // сам код для Redis
	} else {
		return "BadCode", errors.New("это не наш код") // это не наш код.
	}
}

// Проверка подписи Id полученного из Cookie
func checkSignedCookie(sesionId string, hashCompare string) error {

	hsh := hmac.New(sha256.New, []byte(os.Getenv("SESSION_SECRET"))) // secret
	_, err := hsh.Write([]byte(sesionId))                            // есть в куке
	if err != nil {
		return errors.New("не удалось получить хеш ключа Cookie")
	}
	b64 := base64.StdEncoding.EncodeToString(hsh.Sum(nil)) // выдаст хеш  совпадает с JS - https://github.com/tj/node-cookie-signature/blob/master/index.js
	sampleRegexp := regexp.MustCompile(`\=+$`)             // просто с конца цдаляем
	result := sampleRegexp.ReplaceAllString(b64, "")       // расчет совпадает с JS !
	if result == hashCompare {
		return nil
	}
	return errors.New("хеш ключа Cookie не совпадает")
}

// Подписать ID сессии и заменить Escape неразрешенные символы
// т.е. возвращает код для передачи в Cookie
func getSignedKey(sesionId string) (string, error) {

	hsh := hmac.New(sha256.New, []byte(os.Getenv("SESSION_SECRET"))) // secret
	_, err := hsh.Write([]byte(sesionId))                            // есть в куке
	if err != nil {
		return "", errors.New("не удалось получить хеш ключа Cookie")
	}
	b64 := base64.StdEncoding.EncodeToString(hsh.Sum(nil)) // выдаст хеш  совпадает с JS - https://github.com/tj/node-cookie-signature/blob/master/index.js
	sampleRegexp := regexp.MustCompile(`\=+$`)             // подготовка regexp
	result := sampleRegexp.ReplaceAllString(b64, "")       // просто с конца удаляем, расчет совпадает с JS !
	result = url.QueryEscape("s:" + sesionId + "." + result)
	return result, nil
}

// generate RandomKey returns a new random key
func generateRandomKey() (string, error) {
	k := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return "", err
	}
	return strings.TrimRight(base32.StdEncoding.EncodeToString(k), "="), nil
}
