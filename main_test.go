package main

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// структура для тестов maximum и maxChunks
type Tests struct {
	note   string
	slice  []int
	answer int
}

var test = []Tests{
	{
		note:   "empty slice returns wrong answer",
		slice:  []int{},
		answer: 0,
	},
	{
		note:   "one element maxInt in slice returns wrong answer",
		slice:  []int{math.MaxInt},
		answer: math.MaxInt,
	},
	{
		note: "simple slice of 5 elements returned wrong answer",
		slice: []int{
			9, 6, math.MaxInt, 1, 21,
		},
		answer: math.MaxInt,
	},
	{
		note: "wrong answer returns, when slice from identical values",
		slice: []int{
			1, 1, 1, 1,
		},
		answer: 1,
	},
}

// Пишите тесты в этом файле
// Тестируем функцию generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(-5) //реакция программы на отрицательное значение
	require.Equal(t, 0, len(slice), "get invalid sie of slice")
	slice = generateRandomElements(0) //реакция программы на нулевое значение
	require.Equal(t, 0, len(slice), "get invalid sie of slice")
	slice = generateRandomElements(10) //реакция программы на положительное значение
	assert.Equal(t, 10, len(slice), "get invalid sie of slice")
	//проверим кусок слайса(10 элементов) на повторения значений
	sort.Ints(slice) //отсортируем слайс
	rep := 0         //счетчик повторяющихся значений
	for i := 1; i < len(slice); i++ {
		if slice[i-1] == slice[i] {
			rep++
		}
		assert.Greater(t, 2, rep, "more then two repeats in random values") //если генератор повторит более 2 значений-предупредим
	}
	//так как по условию "слайс целых положительных чисел", проверим это условие
	slice = generateRandomElements(1_000_000)
	for i := 1; i < len(slice); i++ {
		assert.Greater(t, slice[i], 1, "0 or negative number was found")
	}
}

func TestMaximum(t *testing.T) {

	for _, r := range test {
		maxi := maximum(r.slice)
		assert.Equal(t, r.answer, maxi, r.note)
	}
}

func TestMaxChunks(t *testing.T) {
	for _, r := range test {
		maxi := maxChunks(r.slice)
		assert.Equal(t, r.answer, maxi, r.note)
	}
	//	также очень интересно пробежать по всем элементам слайса, чтоб понимать, что 100% обрабатывается в горутинах
	// ниже создается слайс заполненный 0, по очереди каждый элемент заполняется 9, проверяется, потом присваевается 1, процесс повторяется со следующим элементом
	slice := make([]int, 10_000)
	for i := 0; i < 10_000; i++ {
		slice[i] = 9
		maxi := maxChunks(slice)
		assert.Equal(t, 9, maxi, "problem appeared during one-by-one test")
		slice[i] = 1
	}
}
