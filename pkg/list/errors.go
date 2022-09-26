package list

import "errors"

var (
	ErrEmptyList  = errors.New("список пуст")
	ErrOutOfRange = errors.New("выход за границу списка")
	ErrWrongArg   = errors.New("передавать в функцию можно либо указатель на узел, либо индекс")
	ErrNoNode     = errors.New("нет указанного узла")
	ErrNilPtr     = errors.New("нулевой указатель")
	ErrSize       = errors.New("размер списка не может быть отрицательным")
)
