package converter

import (
	"fmt"
	userspb "github.com/kalina-malina/IM-PROTOS/generated/users/v1"
)

const (
	//не понравился продукт
	reason_not_liked_product = "liked_product"
	//не понравился сервис
	reason_not_liked_service = "liked_service"
	//не понравился качество и обслуживание
	liked_quality_and_service = "liked_quality_and_service"
	//не понравился персонал
	liked_personal = "liked_personal"
	//не понравился местоположение
	liked_location = "liked_location"
	//не понравился время работы
	liked_time = "liked_time"
	//долгая доставка
	liked_delivery = "liked_delivery"
	//не понравился цена
	liked_price = "liked_price"
	//другое
	liked_other = "liked_other"
	//другое
)

const (
	//не пристойное поведение
	ban_reason_unproper_behavior = "ban_reason_unproper_behavior"
	//спам
	ban_reason_spam = "ban_reason_spam"
	//нарушение правил
	ban_reason_violation_of_rules = "ban_reason_violation_of_rules"
	//другое
	ban_reason_other = "ban_reason_other"
)

func StringToTypeReasonType(r string) (userspb.ReasonTypeEnum, error) {
	switch r {
	case reason_not_liked_product:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRODUCT, nil
	case reason_not_liked_service:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_SERVICE, nil
	case liked_delivery:
		return userspb.ReasonTypeEnum_REASON_LONG_DELIVERY, nil
	case liked_quality_and_service:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_QUALITY_AND_SERVICE, nil
	case liked_personal:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_PERSONAL, nil
	case liked_location:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_LOCATION, nil
	case liked_time:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_TIME, nil
	case liked_price:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRICE, nil
	case liked_other:
		return userspb.ReasonTypeEnum_REASON_OTHER, nil
	default:
		return 0, fmt.Errorf("неизвестная причина при удалении пользователя: %s", r)
	}
}

func ReasonTypeToString(r userspb.ReasonTypeEnum) (string, error) {
	switch r {
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRODUCT:
		return reason_not_liked_product, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_SERVICE:
		return reason_not_liked_service, nil
	case userspb.ReasonTypeEnum_REASON_LONG_DELIVERY:
		return liked_delivery, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_QUALITY_AND_SERVICE:
		return liked_quality_and_service, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_PERSONAL:
		return liked_personal, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_LOCATION:
		return liked_location, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_TIME:
		return liked_time, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRICE:
		return liked_price, nil
	case userspb.ReasonTypeEnum_REASON_OTHER:
		return liked_other, nil
	default:
		return "", fmt.Errorf("неизвестная причина при удалении пользователя: %s", r)

	}
}

func StringToTypeBanReasonType(r string) (userspb.BanReasonTypeEnum, error) {
	switch r {
	case ban_reason_unproper_behavior:
		return userspb.BanReasonTypeEnum_BAN_REASON_UNPROPER_BEHAVIOR, nil
	case ban_reason_spam:
		return userspb.BanReasonTypeEnum_BAN_REASON_SPAM, nil
	case ban_reason_violation_of_rules:
		return userspb.BanReasonTypeEnum_BAN_REASON_VIOLATION_OF_RULES, nil
	case ban_reason_other:
		return userspb.BanReasonTypeEnum_BAN_REASON_OTHER, nil
	default:
		return 0, fmt.Errorf("неизвестная причина при бане пользователя: %s", r)
	}
}

func BanReasonTypeToString(r userspb.BanReasonTypeEnum) (string, error) {
	switch r {
	case userspb.BanReasonTypeEnum_BAN_REASON_UNPROPER_BEHAVIOR:
		return ban_reason_unproper_behavior, nil
	case userspb.BanReasonTypeEnum_BAN_REASON_SPAM:
		return ban_reason_spam, nil
	case userspb.BanReasonTypeEnum_BAN_REASON_VIOLATION_OF_RULES:
		return ban_reason_violation_of_rules, nil
	case userspb.BanReasonTypeEnum_BAN_REASON_OTHER:
		return ban_reason_other, nil
	default:
		return "", fmt.Errorf("неизвестная причина при бане пользователя: %s", r)
	}
}
