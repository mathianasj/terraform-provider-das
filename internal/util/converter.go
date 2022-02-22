package util

import "github.com/hashicorp/terraform/helper/schema"

func GetStringRef(d *schema.ResourceData, key string) *string {
	out := GetString(d, key)

	return &out
}

func GetString(d *schema.ResourceData, key string) string {
	return d.Get(key).(string)
}

func GetSet(d *schema.ResourceData, key string) *schema.Set {
	return d.Get(key).(*schema.Set)
}

func ConvertStringRef(value string) *string {
	return &value
}

func ConvertBoolRef(value bool) *bool {
	return &value
}
