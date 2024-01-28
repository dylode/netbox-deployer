package netbox

import "reflect"

func getUniqueModelTags(v any, uniqueTags map[string]struct{}) {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("model")

		if tag != "" {
			uniqueTags[tag] = struct{}{}
		}

		fieldValue := val.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			getUniqueModelTags(fieldValue.Interface(), uniqueTags)
		case reflect.Slice:
			for j := 0; j < fieldValue.Len(); j++ {
				getUniqueModelTags(fieldValue.Index(j).Interface(), uniqueTags)
			}
		}
	}
}

func AllModelNames() []ModelName {
	uniqueTags := make(map[string]struct{})
	getUniqueModelTags(VirtualMachine{}, uniqueTags)

	result := make([]ModelName, 0, len(uniqueTags))
	for tag := range uniqueTags {
		result = append(result, ModelName(tag))
	}
	return result
}

func HasComponent(vm VirtualMachine, modelName ModelName, modelID ModelID) bool {
	return false
}
