package output

import "github.com/fatih/color"

func PrintError(value any) { // interface{} === any

	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Не известный тип ошибки")
	}
}

//func PrintError(value any) { // interface{} === any
//	intVal, ok := value.(int)
//	if ok {
//		color.Red("Код ошибки: %d", intVal)
//		return
//	}
//
//	strVal, ok := value.(string)
//	if ok {
//		color.Red(strVal)
//		return
//	}
//
//	errVal, ok := value.(error)
//	if ok {
//		color.Red(errVal.Error())
//		return
//	}
//
//	color.Red("Не известный тип ошибки")
//}

func sum[T int | float64 | float32 | string](a, b T) T {
	return a + b
}

//type List[T any] struct {
//	elements []T
//}
