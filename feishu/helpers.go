package feishu

import "time"

// CreateTextField 创建文本字段
func CreateTextField(value string) interface{} {
	return value
}

// CreateNumberField 创建数字字段
func CreateNumberField(value float64) interface{} {
	return value
}

// CreateDateTimeField 创建日期时间字段（Unix秒转为毫秒）
func CreateDateTimeField(timestamp int64) interface{} {
	// 飞书多维表格使用毫秒时间戳
	return timestamp * 1000
}

// CreateDateTimeFieldFromTime 从 time.Time 创建日期时间字段
func CreateDateTimeFieldFromTime(t time.Time) interface{} {
	return t.Unix() * 1000
}

// CreateURLField 创建链接字段
func CreateURLField(url, text string) interface{} {
	return map[string]string{
		"link": url,
		"text": text,
	}
}

// CreateSingleSelectField 创建单选字段
func CreateSingleSelectField(optionName string) interface{} {
	return optionName
}

// CreateMultiSelectField 创建多选字段
func CreateMultiSelectField(optionNames []string) interface{} {
	return optionNames
}

// CreateCheckboxField 创建复选框字段
func CreateCheckboxField(checked bool) interface{} {
	return checked
}

// CreateUserField 创建人员字段
func CreateUserField(userIDs []string) interface{} {
	users := make([]map[string]string, len(userIDs))
	for i, userID := range userIDs {
		users[i] = map[string]string{
			"id": userID,
		}
	}
	return users
}

// CreatePhoneField 创建电话字段
func CreatePhoneField(phone string) interface{} {
	return phone
}

// CreateLocationField 创建地理位置字段
func CreateLocationField(location string) interface{} {
	return location
}
