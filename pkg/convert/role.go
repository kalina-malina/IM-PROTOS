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

func StringToTypeRole(r string) (rolepb.Role, error) {
	switch r {
	case client:
		return rolepb.Role_ROLE_CLIENT, nil
	case admin:
		return rolepb.Role_ROLE_ADMIN, nil
	case operator:
		return rolepb.Role_ROLE_OPERATOR, nil
	case marketing:
		return rolepb.Role_ROLE_MARKETING, nil
	case support:
		return rolepb.Role_ROLE_SUPPORT, nil
	case management:
		return rolepb.Role_ROLE_MANAGEMENT, nil
	default:
		return 0, fmt.Errorf("неизвестная роль в StringToTypeRole: %s", r)
	}
}

func RoleToString(r rolepb.Role) (string, error) {
	switch r {
	case rolepb.Role_ROLE_CLIENT:
		return client, nil
	case rolepb.Role_ROLE_ADMIN:
		return admin, nil
	case rolepb.Role_ROLE_OPERATOR:
		return operator, nil
	case rolepb.Role_ROLE_MARKETING:
		return marketing, nil
	case rolepb.Role_ROLE_SUPPORT:
		return support, nil
	case rolepb.Role_ROLE_MANAGEMENT:
		return management, nil
	default:
		return "", fmt.Errorf("неизвестная роль в RoleToString: %s", r)

	}
}
