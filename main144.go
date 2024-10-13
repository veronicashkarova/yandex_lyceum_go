package main

import "testing"

func TestDeleteVowelsc) {
	// набор тестов
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		value string
		// желаемый результат
		want string
}{
	// тестовые данные № 1
	{
		name: "test №1",
		value: "Literally",
		want: "Ltrlly",
	},
	// тестовые данные № 2
	{
		name: "test №2",
		value: "you",
		want: "y",
	},
}
// перебор всех тестов
for _, tc := range cases {
	tc := tc
	// запуск отдельного теста
	t.Run(tc.name, func(t *testing.T) {
		// тестируем функцию Sum
		got := DeleteVowels(tc.value)
		// проверим полученное значение
		if got != tc.want {
				t.Errorf("Sum(%v) = %v; want %v", tc.value, got, tc.want)
		}
	})
}
}