package postgres

import (
	"reflect"
)

var types map[string]reflect.Type

func init() {
	types = make(map[string]reflect.Type)
	RegType(Arecept{}, "recept")
	RegType(Aregion{}, "region")
}
func RegType(value interface{}, name string) {
	t := reflect.TypeOf(value)
	//name := t.PkgPath() + "." + t.Name()
	//types["ddd"] = t
	//log.Println(types)
	types[name] = t
}

// func GetType(name string) reflect.Value {
// 	return reflect.New(types[name])
// }

func GetType(name string) reflect.Type {
	return types[name]
}
