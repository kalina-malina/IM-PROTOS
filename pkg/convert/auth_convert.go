package converter

import (
	"fmt"
	authpb "github.com/kalina-malina/IM-PROTOS/generated/auth/v1"
)

const (
	sms    = "sms"
	yandex = "yandex"
)

func StringToTypeLogin(s string) (authpb.TypeLogin, error) {
	switch s {
	case sms:
		return authpb.TypeLogin_TYPE_LOGIN_SMS, nil
	case yandex:
		return authpb.TypeLogin_TYPE_LOGIN_YANDEX, nil
	default:
		return 0, fmt.Errorf("неизвестный тип авторизации: %s, должен быть %s или %s", s, sms, yandex)
	}
}

func TypeLoginToString(t authpb.TypeLogin) (string, error) {
	switch t {
	case authpb.TypeLogin_TYPE_LOGIN_SMS:
		return sms, nil
	case authpb.TypeLogin_TYPE_LOGIN_YANDEX:
		return yandex, nil
	default:
		return "", fmt.Errorf("неизвестный тип авторизации: %d, должен быть %s или %s", t, sms, yandex)
	}
}
