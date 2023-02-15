package validator

// 验证对象是否合法

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// PermittedValue 如果列表中存在特定的值，则返回 true
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

// Unique 切片中所有值都是唯一，返回 true
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]struct{})
	for _, v := range values {
		uniqueValues[v] = struct{}{}
	}
	return len(values) == len(uniqueValues)
}
