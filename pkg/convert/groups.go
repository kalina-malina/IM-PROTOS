package converter

import (
	"fmt"
	groups "github.com/kalina-malina/IM-PROTOS/generated/product-service/v1/groups"
)

const (
	group_level1 = "group_level1"
	group_level2 = "group_level2"
	group_level3 = "group_level3"
)

func StringToGroupType(s string) (groups.GroupType, error) {
	switch s {
	case group_level1:
		return groups.GroupType_GROUP_TYPE_LEVEL1, nil
	case group_level2:
		return groups.GroupType_GROUP_TYPE_LEVEL2, nil
	case group_level3:
		return groups.GroupType_GROUP_TYPE_LEVEL3, nil
	default:
		return 0, fmt.Errorf("ошибка при конвертации строки в тип группы: %s", s)
	}
}

func GroupTypeToString(t groups.GroupType) (string, error) {
	switch t {
	case groups.GroupType_GROUP_TYPE_LEVEL1:
		return group_level1, nil
	case groups.GroupType_GROUP_TYPE_LEVEL2:
		return group_level2, nil
	case groups.GroupType_GROUP_TYPE_LEVEL3:
		return group_level3, nil
	default:
		return "", fmt.Errorf("ошибка при конвертации типа группы в строку: %d", t)
	}
}
