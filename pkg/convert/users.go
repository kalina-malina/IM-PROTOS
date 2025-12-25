package converter

import (
	"fmt"
	getUsers "github.com/kalina-malina/IM-PROTOS/generated/users/v1"
	abGroups "github.com/kalina-malina/IM-PROTOS/generated/users/v1/abGroups"
)

const (
	ab_test_group_a = "group_a"
	ab_test_group_b = "group_b"
	ab_test_group_c = "group_c"
	ab_test_group_d = "group_d"
)

func StringToTypeABTestGroup(r string) (abGroups.ABTestGroup, error) {
	switch r {
	case ab_test_group_a:
		return abGroups.ABTestGroup_AB_TEST_GROUP_A, nil
	case ab_test_group_b:
		return abGroups.ABTestGroup_AB_TEST_GROUP_B, nil
	case ab_test_group_c:
		return abGroups.ABTestGroup_AB_TEST_GROUP_C, nil
	case ab_test_group_d:
		return abGroups.ABTestGroup_AB_TEST_GROUP_D, nil
	default:
		return 0, fmt.Errorf("неизвестная группа AB теста в StringToTypeABTestGroup: %s", r)
	}
}
func ABTestGroupToString(r abGroups.ABTestGroup) (string, error) {
	switch r {
	case abGroups.ABTestGroup_AB_TEST_GROUP_A:
		return ab_test_group_a, nil
	case abGroups.ABTestGroup_AB_TEST_GROUP_B:
		return ab_test_group_b, nil
	case abGroups.ABTestGroup_AB_TEST_GROUP_C:
		return ab_test_group_c, nil
	case abGroups.ABTestGroup_AB_TEST_GROUP_D:
		return ab_test_group_d, nil
	default:
		return "", fmt.Errorf("неизвестная группа AB теста в ABTestGroupToString: %d", r)
	}
}

const (
	sort_by_guid           = "guid"
	sort_by_card_number    = "card_number"
	sort_by_email          = "email"
	sort_by_first_name     = "first_name"
	sort_by_last_name      = "last_name"
	sort_by_role           = "role"
	sort_by_ab_test_group  = "ab_test_group"
	sort_by_verified_email = "verified_email"
	sort_by_phone          = "phone"
	// sort_by_created_at      = "created_at"
	// sort_by_updated_at      = "updated_at"
	// sort_by_deleted_at      = "deleted_at"
	// sort_by_status          = "status"
	// sort_by_type            = "type"
)

func StringToTypeSortBy(r string) (getUsers.SortBy, error) {
	switch r {
	case sort_by_guid:
		return getUsers.SortBy_SORT_BY_GUID, nil
	case sort_by_first_name:
		return getUsers.SortBy_SORT_BY_FIRST_NAME, nil
	case sort_by_last_name:
		return getUsers.SortBy_SORT_BY_LAST_NAME, nil
	case sort_by_card_number:
		return getUsers.SortBy_SORT_BY_CARD_NUMBER, nil
	case sort_by_email:
		return getUsers.SortBy_SORT_BY_EMAIL, nil
	case sort_by_role:
		return getUsers.SortBy_SORT_BY_ROLE, nil
	case sort_by_ab_test_group:
		return getUsers.SortBy_SORT_BY_AB_TEST_GROUP, nil
	case sort_by_verified_email:
		return getUsers.SortBy_SORT_BY_VERIFIED_EMAIL, nil
	case sort_by_phone:
		return getUsers.SortBy_SORT_BY_PHONE, nil
	default:
		return 0, fmt.Errorf("неизвестный тип сортировки в  должны быть отправлено поле тогоже столбца таблицы: %s", r)
	}
}

func SortByToString(r getUsers.SortBy) (string, error) {
	switch r {
	case getUsers.SortBy_SORT_BY_GUID:
		return sort_by_guid, nil
	case getUsers.SortBy_SORT_BY_FIRST_NAME:
		return sort_by_first_name, nil
	case getUsers.SortBy_SORT_BY_LAST_NAME:
		return sort_by_last_name, nil
	case getUsers.SortBy_SORT_BY_CARD_NUMBER:
		return sort_by_card_number, nil
	case getUsers.SortBy_SORT_BY_EMAIL:
		return sort_by_email, nil
	case getUsers.SortBy_SORT_BY_ROLE:
		return sort_by_role, nil
	case getUsers.SortBy_SORT_BY_AB_TEST_GROUP:
		return sort_by_ab_test_group, nil
	case getUsers.SortBy_SORT_BY_VERIFIED_EMAIL:
		return sort_by_verified_email, nil
	case getUsers.SortBy_SORT_BY_PHONE:
		return sort_by_phone, nil
	default:
		return "", fmt.Errorf("неизвестный тип сортировки в должны быть отправлено поле тогоже столбца таблицы: %d", r)
	}
}

func StringToTypeSortOrder(r string) (getUsers.SortOrder, error) {
	switch r {
	case "asc":
		return getUsers.SortOrder_SORT_ORDER_ASC, nil
	case "desc":
		return getUsers.SortOrder_SORT_ORDER_DESC, nil
	default:
		return 0, fmt.Errorf("неизвестный тип сортировки в должны быть отправлено asc или desc: %s", r)
	}
}

func SortOrderToString(r getUsers.SortOrder) (string, error) {
	switch r {
	case getUsers.SortOrder_SORT_ORDER_ASC:
		return "asc", nil
	case getUsers.SortOrder_SORT_ORDER_DESC:
		return "desc", nil
	default:
		return "", fmt.Errorf("неизвестный тип сортировки в должны быть отправлено asc или desc: %d", r)
	}
}
