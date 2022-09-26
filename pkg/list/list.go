package list

import (
	"github.com/google/uuid"
)

type (
	// Структура узла
	// uuid в списке и узле используется для функций которые принимают в качестве параметра указатель на узел,
	// чтобы определить принадлежит ли данный узел данному списку
	Node struct {
		Data interface{}
		next *Node
		prev *Node
		id   uuid.UUID
	}

	// Структура списка
	List struct {
		head  *Node
		tail  *Node
		count int
		id    uuid.UUID
	}
)

// Конструктор списка
// Возможна инициализация списка
func NewList(args ...interface{}) *List {

	if args == nil {
		return &List{nil, nil, 0, uuid.New()}
	}

	l := List{}
	l.id = uuid.New()
	l.head = &Node{args[0], nil, nil, l.id}
	l.tail = l.head
	l.count++

	ptr := l.head

	for _, val := range args[1:] {
		ptr.next = &Node{val, nil, ptr, l.id}
		l.tail = ptr.next
		ptr = ptr.next
		l.count++
	}

	return &l
}

// Возвращает указатель на голову списка
// Возвращает нулевой указатель и ошибку если список пуст
func (l *List) Begin() (*Node, error) {
	if l.count == 0 {
		return nil, ErrEmptyList
	} else {
		return l.head, nil
	}
}

// Возвращает указатель на хвост списка
// Возвращает нулевой указатель и ошибку если список пуст
func (l *List) End() (*Node, error) {
	if l.count == 0 {
		return nil, ErrEmptyList
	} else {
		return l.tail, nil
	}
}

// Сдвигает указатель на указанное целочисленное значение
// Возвращает ошибку если указатель вышел за пределы списка,
// если был передан нулевой указатель
func Advance(ptr **Node, num int) error {
	// Проверка указателя
	if ptr == nil || *ptr == nil {
		return ErrNilPtr
	}

	// Выбор в какую сторону перемещать указатель
	if num > 0 {
		for i := 0; i < num; i++ {
			*ptr = (*ptr).next
			if *ptr == nil {
				return ErrOutOfRange
			}
		}
	} else {
		for i := 0; i > num; i-- {
			*ptr = (*ptr).prev
			if *ptr == nil {
				return ErrOutOfRange
			}
		}
	}
	return nil
}

// Возвращает указатель на узел по указанному индексу
// Возвращает ошибку если список пуст,
// если индекс вышел за пределы списка
func (l *List) NewIterator(num int) (*Node, error) {
	// Проверка на пустоту
	if l.count == 0 {
		return nil, ErrEmptyList
	}

	// Проверка диапазона
	if num > l.count || num < 1 {
		return nil, ErrOutOfRange
	}

	// Выбор с какой стороны двигаться
	var ptr *Node
	if num < l.count/2 {
		ptr = l.head

		for i := 1; i < num; i++ {
			ptr = ptr.next
		}

		return ptr, nil
	} else {
		ptr = l.tail

		for i := 0; i < l.count-num; i++ {
			ptr = ptr.prev
		}

		return ptr, nil
	}
}

// Добавляет элемент в начало
func (l *List) Push_front(data interface{}) {

	// Обработка случая с пустым списком
	if l.count == 0 {
		l.head = &Node{data, nil, nil, l.id}
		l.tail = l.head
		l.count++
		return
	}

	l.head = &Node{data, l.head, nil, l.id}
	l.head.next.prev = l.head
	l.count++
}

// Добавляет элемент в конец
func (l *List) Push_back(data interface{}) {

	// Обработка случая с пустым списком
	if l.count == 0 {
		l.head = &Node{data, nil, nil, l.id}
		l.tail = l.head
		l.count++
		return
	}

	l.tail = &Node{data, nil, l.tail, l.id}
	l.tail.prev.next = l.tail
	l.count++
}

// Используется в функции Insert
func (l *List) insertIndex(index int, data interface{}) (err error) {
	// Проверка индекса на принадлежность к диапазону списка
	if index > l.count+1 || index < 1 {
		return ErrOutOfRange
	}

	// Обработка случая когда функция Insert эквивалентна функции Push_front
	if index == 1 && l.count == 0 {
		l.Push_front(data)
		return nil
	}

	// Обработка случая когда функция insert эквивалентна функции Push_back
	if index == l.count+1 {
		l.Push_back(data)
		return nil
	}

	// Выбор откуда ближе двигаться
	if index < l.count/2 {
		ptr := l.head
		for i := 1; i < index; i++ {
			ptr = ptr.next
		}

		ptr.prev.next = &Node{data, ptr, ptr.prev, l.id}
		ptr.prev = ptr.prev.next
		l.count++
		return nil
	} else {
		ptr := l.tail
		for i := 0; i < l.count-index; i++ {
			ptr = ptr.prev
		}

		ptr.prev.next = &Node{data, ptr, ptr.prev, l.id}
		ptr.prev = ptr.prev.next
		l.count++
		return nil
	}

}

// Используется в функции Insert
func (l *List) insertNode(ptr *Node, data interface{}) (err error) {
	if l.id != ptr.id {
		return ErrNoNode
	}

	// Обработка случая когда функция Insert эквивалентна функции Push_front
	if ptr == l.head {
		l.Push_front(data)
		return nil
	}

	ptr.prev.next = &Node{data, ptr, ptr.prev, l.id}
	ptr.prev = ptr.prev.next
	l.count++
	return nil
}

// Вставляет элемент в список перед указанным узлом (по индексу либо по указателю)
// Возвращает ошибку если индекс находится вне диапазона списка,
// если указанного узла нет,
// если передан не индекс и не указатель на узел
func (l *List) Insert(ptr interface{}, data interface{}) (err error) {
	switch p := ptr.(type) {
	case int:
		err = l.insertIndex(p, data)
		return err
	case *Node:
		err = l.insertNode(p, data)
		return err
	default:
		return ErrWrongArg
	}
}

// Возвращает данные из головы списка
// Возвращает ошибку если список пуст
func (l *List) Front() (interface{}, error) {
	if l.count == 0 {
		return nil, ErrEmptyList
	}

	return l.head.Data, nil
}

// Возвращает данные из хвоста списка
// Возвращает ошибку если список пуст
func (l *List) Back() (interface{}, error) {
	if l.count == 0 {
		return nil, ErrEmptyList
	}

	return l.tail.Data, nil
}

// Удаляет элемент из головы списка
// Возвращает ошибку если список пуст
func (l *List) Pop_front() error {
	if l.count == 0 {
		return ErrEmptyList
	}

	if l.head == l.tail {
		l.count--
		l.head = nil
		l.tail = nil
		return nil
	}

	l.head = l.head.next
	l.count--

	return nil
}

// Удаляет элемент из хвоста списка
// Возвращает ошибку если список пуст
func (l *List) Pop_back() error {
	if l.count == 0 {
		return ErrEmptyList
	}

	if l.head == l.tail {
		l.count--
		l.head = nil
		l.tail = nil
		return nil
	}

	l.tail.prev.next = nil
	l.tail = l.tail.prev
	l.count--

	return nil
}

// Удаляет указанный элемент (по индексу либо по указателю)
// Возвращает указатель на следующий элемент
// Возвращает ошибку если список пуст,
// если индекс находится вне диапазона списка,
// если указанного узла нет,
// если передан не индекс и не указатель на узел
func (l *List) Erase(value interface{}) (*Node, error) {
	switch val := value.(type) {
	case int:
		p, err := l.eraseIndex(val)
		return p, err
	case *Node:
		p, err := l.eraseNode(val)
		return p, err
	default:
		return nil, ErrWrongArg
	}
}

// Используется в функции Erase
func (l *List) eraseIndex(index int) (*Node, error) {
	// Проверка диапазона
	if index < 1 || index > l.count {
		return nil, ErrOutOfRange
	}

	// Обработка случая когда функция Erase эквивалентна функции Pop_front
	if index == 1 {
		l.Pop_front()
		p, err := l.Begin()
		return p, err
	}

	// Обработка случая когда функция Erase эквивалентна функции Pop_back
	if index == l.count {
		err := l.Pop_back()

		return nil, err
	}

	// Выбор с какой стороны ближе двигаться
	if index < l.count/2 {
		ptr := l.head
		for i := 1; i < index; i++ {
			ptr = ptr.next
		}
		ptr.prev.next = ptr.next
		ptr.next.prev = ptr.prev
		l.count--
		return ptr.next, nil
	} else {
		ptr := l.tail
		for i := 0; i < l.count-index; i++ {
			ptr = ptr.prev
		}
		ptr.prev.next = ptr.next
		ptr.next.prev = ptr.prev
		l.count--
		return ptr.next, nil
	}
}

// Используется в функции Erase
func (l *List) eraseNode(ptr *Node) (*Node, error) {
	// Проверка id
	if ptr.id != l.id {
		return nil, ErrNoNode
	}

	// Обработка случая когда функция Erase эквивалентна функции Pop_front
	if ptr == l.head {
		l.Pop_back()
		p, err := l.Begin()
		return p, err
	}

	// Обработка случая когда функция Erase эквивалентна функции Pop_back
	if ptr == l.tail {
		err := l.Pop_back()
		return nil, err
	}

	ptr.prev.next = ptr.next
	ptr.next.prev = ptr.prev
	l.count--
	return ptr.next, nil
}

// Проверка на пустоту
// Возвращает true если список пуст, иначе возвращает false
func (l *List) Empty() bool {
	if l.count == 0 {
		return true
	} else {
		return false
	}
}

// Удаляет все элементы из списка с указанным значением
// Возвращает ошибку если список пуст
func (l *List) Remove(data interface{}) error {
	// Проверка на пустоту
	if l.count == 0 {
		return ErrEmptyList
	}

	ptr := l.head
	for ptr != nil {
		if ptr.Data == data {
			ptr, _ = l.eraseNode(ptr)
			continue
		}
		ptr = ptr.next
	}
	return nil
}

// Заменяет содержимое списка содержимым указанного среза
func (l *List) Assign(slice []interface{}) {
	if l.count != 0 {
		l.head = nil
		l.tail = nil
		l.count = 0
	}

	for _, val := range slice {
		l.Push_back(val)
	}
}

// Удаляет все элементы из списка
// Возвращает ошибку если список уже был пуст
func (l *List) Clear() error {
	// Проверка на пустоту
	if l.count == 0 {
		return ErrEmptyList
	} else {
		l.head = nil
		l.tail = nil
		l.count = 0
		return nil
	}
}

// Меняет местами содержимое двух списков
func (l *List) Swap(obj *List) {
	tempH := obj.head
	tempT := obj.tail
	tempC := obj.count
	tempID := obj.id

	obj.head = l.head
	obj.tail = l.tail
	obj.count = l.count
	obj.id = l.id

	l.head = tempH
	l.tail = tempT
	l.count = tempC
	l.id = tempID
}

// Меняет размер списка на указанный
// Если размер меняется на бОльший, то добавляются узлы без данных
// Возвращает ошибку если был передан отрицательный размер
func (l *List) Resize(size int) error {
	// Проверка переданного размера на соответствие
	if size < 0 {
		return ErrSize
	}

	// Обработка ситуации когда функция Resize эквивалентна функции Clear
	if size == 0 {
		l.Clear()
		return nil
	}

	// Обработка ситуации когда переданный размер меньше размера списка
	if size < l.count {
		for i := 0; i < l.count-size; {
			l.Pop_back()
		}
		return nil
	}

	// Обработка ситуации когда переданный размер больше размера списка
	for i := 0; i < size-l.count; {
		l.Push_back(nil)
	}
	return nil

}

// Меняет порядок следования элементов списка на противоположный
// Возвращает ошибку если контейнер пуст
func (l *List) Reverse() error {
	// Проверка на пустоту
	if l.count == 0 {
		return ErrEmptyList
	}

	temp := l.head
	l.head = l.tail
	l.tail = temp

	ptr := l.head
	for i := 0; i < l.count; i++ {
		ptr.next = temp
		ptr.next = ptr.prev
		ptr.prev = temp
		ptr = ptr.next
	}

	return nil
}

// Возвращает размер односвязного списка
func (l *List) Size() (count int) {
	return l.count
}

// Вставляет указанный список перед указанным узлом (по номеру либо по указателю)
// После выполнения указанный переданный список становится пустым
// Возвращает ошибку если указанного узла нет,
// если был передан не указатель на узел и не число
// если число находится за пределами списка
// если список пуст и индекс не равен 1
func (l *List) Splice(val interface{}, obj *List) error {
	switch i := val.(type) {
	case int:
		err := l.spliceInt(i, obj)
		return err
	case *Node:
		err := l.spliceNode(i, obj)
		return err
	default:
		return ErrWrongArg
	}
}

// Используется в функции Splice
// Меняет id узлов списка на id списка в который они добавляются
func (l *List) changeID(id uuid.UUID) {
	ptr := l.head
	for i := 0; i < l.count; i++ {
		ptr.id = id
		ptr = ptr.next
	}
}

// Используется в функции Splice
func (l *List) spliceInt(index int, obj *List) error {
	// Проверка диапазона
	if index < 0 || index > l.count+1 {
		return ErrOutOfRange
	}

	// Обрабатываем ситуацию когда указанный список пуст
	if obj.count == 0 {
		return ErrEmptyList
	}

	// Обрабатываем ситуацию когда список в который добавляем пуст и номер равен единице
	if l.count == 0 && index == 1 {
		l.head = obj.head
		l.tail = obj.tail
		l.count = obj.count
		l.id = obj.id

		obj.Clear()

		return nil
	}

	// Меняем id узлов указанного списка на id списка в который добавляем
	obj.changeID(l.id)

	// Обрабатываем ситуацию когда указанный список добавляется в конец списка
	if index == l.count+1 {
		l.tail.next = obj.head
		obj.head = l.tail
		l.count += obj.count
		obj.Clear()
		return nil
	}

	// Обрабатываем ситуацию когда указанный список добавляет в середину списка
	ptr := l.head
	Advance(&ptr, index-1)

	ptr.prev.next = obj.head
	obj.head.prev = ptr.prev
	obj.tail.next = ptr
	ptr.prev = obj.tail

	l.count += obj.count

	obj.Clear()

	return nil
}

func (l *List) spliceNode(ptr *Node, obj *List) error {
	// Обрабатываем ситуацию когда указанный узел отсутствует в списке
	if l.id != ptr.id {
		return ErrNoNode
	}

	// Обрабатываем ситуацию когда указанный список пуст
	if obj.count == 0 {
		return nil
	}

	// Обрабатываем ситуацию когда указанный узел является головой списка
	if ptr == l.head {
		obj.tail.next = l.head
		l.head.prev = obj.tail
		l.count += obj.count

		obj.Clear()
		return nil
	}

	// Обрабатывем случай когда указанный список добавляется в середину списка
	ptr.prev.next = obj.head
	obj.head.prev = ptr.prev

	obj.tail.next = ptr
	ptr.prev = obj.tail

	l.count += obj.count

	obj.Clear()
	return nil
}

// Удаляет из списка все дубликаты
// Возвращает ошибку если список пуст
func (l *List) Unique() error {
	if l.count == 0 {
		return ErrEmptyList
	}

	lMap := make(map[interface{}]interface{})
	for !l.Empty() {
		lMap[l.head.Data] = l.head.Data
		l.Pop_front()
	}
	for _, val := range lMap {
		l.Push_front(val)
	}
	return nil
}

// Возвращает данные указанного узла (по номеру либо по указателю)
// Возвращает ошибку если список пуст,
// если номер выходит за границы списка,
// если узла по такому указателю в списке нет,
// если передано не число и не указатель на узел
func (l *List) GetData(val interface{}) (interface{}, error) {
	switch i := val.(type) {
	case int:
		ptr, err := l.NewIterator(i)
		return ptr.Data, err
	case *Node:
		// Обрабатываем ситуацию когда указанный узел отсутствует в списке
		if l.id != i.id {
			return nil, ErrNoNode
		}

		return i.Data, nil
	default:
		return nil, ErrWrongArg
	}
}
