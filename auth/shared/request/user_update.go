package request

// UserUpdate struct
// Contains field and value to update user info
type UserUpdate []struct {
	Field string `json:"field" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func (u *UserUpdate) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, v := range *u {
		result[v.Field] = v.Value
	}
	return result
}
