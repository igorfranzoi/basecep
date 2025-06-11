package utils

import (
	"reflect"

	"github.com/google/uuid"
)

func ReturnUUID[T any](obj T) {
	val := reflect.ValueOf(obj)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return
	}

	idField := val.Elem().FieldByName("ID")
	if idField.IsValid() && idField.CanSet() && idField.Type() == reflect.TypeOf(uuid.UUID{}) {
		id := idField.Interface().(uuid.UUID)
		if id == uuid.Nil {
			idField.Set(reflect.ValueOf(uuid.New()))
		}
	}
}
