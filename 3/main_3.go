package main

// Задание 3
// Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число". Ваша структура должна поддерживать следующие операции:

// Добавление элемента: Метод Add(key string, value int),
// который добавляет новую пару "ключ-значение" в карту.

// Удаление элемента: Метод Remove(key string),
// который удаляет элемент по ключу из карты.

// Копирование карты: Метод Copy() map[string]int,
// который возвращает новую карту, содержащую все элементы текущей карты.

// Проверка наличия ключа: Метод Exists(key string) bool,
//  который проверяет, существует ли ключ в карте.

// Получение значения: Метод Get(key string) (int, bool),
//  который возвращает значение по ключу и булевый флаг,
//  указывающий на успешность операции.

// Напишите unit тесты к созданным функциям

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

type StringIntMap struct {
	data map[string]int
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

func main() {
	sim := NewStringIntMap()
	sim.Add("one", 1)
	sim.Add("two", 2)

	value, exists := sim.Get("one")
	if exists {
		println("Value for 'one':", value)
	} else {
		println("'one' does not exist")
	}
	sim.Remove("one")
	exists = sim.Exists("one")
	if exists {
		println("'one' still exists")
	} else {
		println("'one' has been removed")
	}

	copyMap := sim.Copy()
	for k, v := range copyMap {
		println("Copy - Key:", k, "Value:", v)
	}
}
