package converter

import (
	"fmt"
	rolepb "github.com/kalina-malina/IM-PROTOS/generated/users/v1/role"
)

const (
	client     = "client"
	admin      = "admin"
	operator   = "operator"
	marketing  = "marketing"
	support    = "support"
	management = "management"
)

func StringToTypeRole(r string) (rolepb.UserRole, error) {
	switch r {
	case client:
		return rolepb.UserRole_USER_ROLE_CLIENT, nil
	case admin:
		return rolepb.UserRole_USER_ROLE_ADMIN, nil
	case operator:
		return rolepb.UserRole_USER_ROLE_OPERATOR, nil
	case marketing:
		return rolepb.UserRole_USER_ROLE_MARKETING, nil
	case support:
		return rolepb.UserRole_USER_ROLE_SUPPORT, nil
	case management:
		return rolepb.UserRole_USER_ROLE_MANAGEMENT, nil
	default:
		return 0, fmt.Errorf("неизвестная роль в StringToTypeRole: %s", r)
	}
}

func RoleToString(r rolepb.UserRole) (string, error) {
	switch r {
	case rolepb.UserRole_USER_ROLE_CLIENT:
		return client, nil
	case rolepb.UserRole_USER_ROLE_ADMIN:
		return admin, nil
	case rolepb.UserRole_USER_ROLE_OPERATOR:
		return operator, nil
	case rolepb.UserRole_USER_ROLE_MARKETING:
		return marketing, nil
	case rolepb.UserRole_USER_ROLE_SUPPORT:
		return support, nil
	case rolepb.UserRole_USER_ROLE_MANAGEMENT:
		return management, nil
	default:
		return "", fmt.Errorf("неизвестная роль в RoleToString: %s", r)

	}
}
