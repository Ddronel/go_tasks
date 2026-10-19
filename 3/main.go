package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func main() {
	var numbers StringIntMap
	numbers.Add("go", 10)
	numbers.Add("zero", 0)

	fmt.Println("После добавления:", numbers.data)

	numbers.Add("go", 20)
	fmt.Println("После обновления go:", numbers.data)

	fmt.Printf("Exists(\"go\") = %t\n", numbers.Exists("go"))
	fmt.Printf("Exists(\"missing\") = %t\n", numbers.Exists("missing"))

	for _, key := range []string{"go", "zero", "missing"} {
		value, ok := numbers.Get(key)
		fmt.Printf("Get(%q): value=%d, ok=%t\n", key, value, ok)
	}

	copied := numbers.Copy()
	copied["go"] = 99
	fmt.Println("Копия после изменения:", copied)
	fmt.Println("Оригинал:", numbers.data)

	numbers.Remove("go")
	fmt.Println("Оригинал после удаления go:", numbers.data)
	fmt.Printf("Exists(\"go\") после удаления = %t\n", numbers.Exists("go"))
	fmt.Println("Копия после удаления из оригинала:", copied)

}

func (m *StringIntMap) Add(k string, v int) {
	if m.data == nil {
		m.data = make(map[string]int)
	}

	m.data[k] = v
}

func (m *StringIntMap) Remove(k string) {
	delete(m.data, k)
}

func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.data))

	for k, v := range m.data {
		newMap[k] = v
	}

	return newMap
}

func (m *StringIntMap) Exists(k string) bool {
	_, exists := m.data[k]
	return exists
}

func (m *StringIntMap) Get(k string) (int, bool) {
	v, ok := m.data[k]
	return v, ok
}
