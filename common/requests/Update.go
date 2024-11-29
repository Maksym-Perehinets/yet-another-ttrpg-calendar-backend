package requests

type Update []struct {
	Field string `json:"field" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func (u *Update) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, v := range *u {
		result[v.Field] = v.Value
	}
	return result
}
