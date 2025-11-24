package converter

import (
	"fmt"
	userspb "github.com/kalina-malina/IM-PROTOS/generated/users/v1"
)

const (
	client     = "client"
	admin      = "admin"
	operator   = "operator"
	marketing  = "marketing"
	support    = "support"
	management = "management"
)

func StringToTypeRole(r string) (userspb.Role, error) {
	switch r {
	case client:
		return userspb.Role_ROLE_CLIENT, nil
	case admin:
		return userspb.Role_ROLE_ADMIN, nil
	case operator:
		return userspb.Role_ROLE_OPERATOR, nil
	case marketing:
		return userspb.Role_ROLE_MARKETING, nil
	case support:
		return userspb.Role_ROLE_SUPPORT, nil
	case management:
		return userspb.Role_ROLE_MANAGEMENT, nil
	default:
		return 0, fmt.Errorf("неизвестная роль: %s", r)
	}
}

func RoleToString(r userspb.Role) (string, error) {
	switch r {
	case userspb.Role_ROLE_CLIENT:
		return client, nil
	case userspb.Role_ROLE_ADMIN:
		return admin, nil
	case userspb.Role_ROLE_OPERATOR:
		return operator, nil
	case userspb.Role_ROLE_MARKETING:
		return marketing, nil
	case userspb.Role_ROLE_SUPPORT:
		return support, nil
	case userspb.Role_ROLE_MANAGEMENT:
		return management, nil
	default:
		return "", fmt.Errorf("неизвестная роль: %s", r)

	}
}
