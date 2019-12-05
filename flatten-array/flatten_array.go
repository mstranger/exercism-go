package flatten

import "reflect"

// Flatten takes a nested list and returns a single flattened list
// with all values except nil.
func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}

	switch reflect.TypeOf(nested).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(nested)

		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == nil {
				continue
			}
			t := s.Index(i).Interface().(interface{})
			result = append(result, Flatten(t)...)
		}
	default:
		result = append(result, nested)
	}

	return result
}
