package convert

import (
	userspb "github.com/kalina-malina/IM-PROTOS/generated/users/v1"
)
enum ABTestGroup {
  AB_TEST_GROUP_A = 0;
  AB_TEST_GROUP_B = 1;
  AB_TEST_GROUP_C = 2;
  AB_TEST_GROUP_D = 3;
}

const (
	ab_test_group_a = "group_a"
	ab_test_group_b = "group_b"
	ab_test_group_c = "group_c"
	ab_test_group_d = "group_d"
)

func StringToTypeABTestGroup(r string) (userspb.ABTestGroup, error) {
	switch r {
	case ab_test_group_a:
		return userspb.ABTestGroup_AB_TEST_GROUP_A, nil
	case ab_test_group_b:
		return userspb.ABTestGroup_AB_TEST_GROUP_B, nil
	case ab_test_group_c:
		return userspb.ABTestGroup_AB_TEST_GROUP_C, nil
	case ab_test_group_d:
		return userspb.ABTestGroup_AB_TEST_GROUP_D, nil
	default:
		return 0, fmt.Errorf("неизвестная группа AB теста в StringToTypeABTestGroup: %s", r)
	}
}
func ABTestGroupToString(r userspb.ABTestGroup) (string, error) {
	switch r {
	case userspb.ABTestGroup_AB_TEST_GROUP_A:
		return ab_test_group_a, nil
	case userspb.ABTestGroup_AB_TEST_GROUP_B:
		return ab_test_group_b, nil
	case userspb.ABTestGroup_AB_TEST_GROUP_C:
		return ab_test_group_c, nil
	case userspb.ABTestGroup_AB_TEST_GROUP_D:
		return ab_test_group_d, nil
	default:
		return "", fmt.Errorf("неизвестная группа AB теста в ABTestGroupToString: %d", r)
	}
}