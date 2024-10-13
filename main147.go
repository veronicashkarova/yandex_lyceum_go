package main

import (
	"testing"
)

func TestGetUTFLength (t *testing.T) {
	// набор тестов
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		values []byte
		// желаемый результат
		want int
		err error
}{
	// тестовые данные № 1
	{
		name: "test №1",
		values: []byte{'h','e','l','l','o'},
		want: 5,
		err: nil,
	},
	// тестовые данные № 2
	{
		name: "test №2",
		values: []byte{},
		want: 0,
		err: nil,
	},
	{
		name: "test №3",
		values: []byte{0xFF},
		want: 0,
		err: ErrInvalidUTF8,
	},
}
// перебор всех тестов
for _, tc := range cases {
	tc := tc
	// запуск отдельного теста
	t.Run(tc.name, func(t *testing.T) {
		// тестируем функцию Sum
		got,err := GetUTFLength(tc.values)
		// проверим полученное значение
		if err != tc.err || got != tc.want {
				t.Errorf("Sum(%v) = %v; want %v", tc.values, got, tc.want)
		}
	})
}
}