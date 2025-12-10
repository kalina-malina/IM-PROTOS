package converter

import (
	"fmt"
	userspb "github.com/kalina-malina/IM-PROTOS/generated/users/v1"
)

const (
	//не понравился продукт
	not_liked_product = "not_liked_product"
	//не понравился сервис
	not_liked_service = "not_liked_service"
	//не понравился качество и обслуживание
	liked_quality_and_service = "not_liked_quality_and_service"
	//не понравился персонал
	not_liked_personal = "not_liked_personal"
	//не понравился местоположение
	not_liked_location = "not_liked_location"
	//не понравился время работы
	not_liked_time = "not_liked_time"
	//долгая доставка
	not_liked_delivery = "not_liked_delivery"
	//не понравился цена
	not_liked_price = "not_liked_price"
	//другое
	not_liked_other = "not_liked_other"
)

const (
	//не пристойное поведение
	unproper_behavior = "unproper_behavior"
	//спам
	spam = "spam"
	//нарушение правил
	violation_of_rules = "violation_of_rules"
	//другое
	other = "other"
)

func StringToTypeReasonType(r string) (userspb.ReasonTypeEnum, error) {
	switch r {
	case not_liked_product:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRODUCT, nil
	case not_liked_service:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_SERVICE, nil
	case not_liked_delivery:
		return userspb.ReasonTypeEnum_REASON_LONG_DELIVERY, nil
	case liked_quality_and_service:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_QUALITY_AND_SERVICE, nil
	case not_liked_personal:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_PERSONAL, nil
	case not_liked_location:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_LOCATION, nil
	case not_liked_time:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_TIME, nil
	case not_liked_price:
		return userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRICE, nil
	case not_liked_other:
		return userspb.ReasonTypeEnum_REASON_OTHER, nil
	default:
		return 0, fmt.Errorf("неизвестная причина при удалении пользователя: %s", r)
	}
}

func ReasonTypeToString(r userspb.ReasonTypeEnum) (string, error) {
	switch r {
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRODUCT:
		return not_liked_product, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_SERVICE:
		return not_liked_service, nil
	case userspb.ReasonTypeEnum_REASON_LONG_DELIVERY:
		return not_liked_delivery, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_QUALITY_AND_SERVICE:
		return liked_quality_and_service, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_PERSONAL:
		return not_liked_personal, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_LOCATION:
		return not_liked_location, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_TIME:
		return not_liked_time, nil
	case userspb.ReasonTypeEnum_REASON_NOT_LIKED_PRICE:
		return not_liked_price, nil
	case userspb.ReasonTypeEnum_REASON_OTHER:
		return not_liked_other, nil
	default:
		return "", fmt.Errorf("неизвестная причина при удалении пользователя: %s", r)

	}
}

func StringToTypeBanReasonType(r string) (userspb.BanReasonTypeEnum, error) {
	switch r {
	case unproper_behavior:
		return userspb.BanReasonTypeEnum_BAN_REASON_UNPROPER_BEHAVIOR, nil
	case spam:
		return userspb.BanReasonTypeEnum_BAN_REASON_SPAM, nil
	case violation_of_rules:
		return userspb.BanReasonTypeEnum_BAN_REASON_VIOLATION_OF_RULES, nil
	case other:
		return userspb.BanReasonTypeEnum_BAN_REASON_OTHER, nil
	default:
		return 0, fmt.Errorf("неизвестная причина при бане пользователя: %s", r)
	}
}

func BanReasonTypeToString(r userspb.BanReasonTypeEnum) (string, error) {
	switch r {
	case userspb.BanReasonTypeEnum_BAN_REASON_UNPROPER_BEHAVIOR:
		return unproper_behavior, nil
	case userspb.BanReasonTypeEnum_BAN_REASON_SPAM:
		return spam, nil
	case userspb.BanReasonTypeEnum_BAN_REASON_VIOLATION_OF_RULES:
		return violation_of_rules, nil
	case userspb.BanReasonTypeEnum_BAN_REASON_OTHER:
		return other, nil
	default:
		return "", fmt.Errorf("неизвестная причина при бане пользователя: %s", r)
	}
}
