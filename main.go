package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 { //если запрошен невозможный размер слайса, возращаем пустой
		return []int{}
	}
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int() //упростил генерацию в соответствии с замечанием
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) <= 0 { //если слайс отрицательный или пустой- возвращаем 0
		return 0
	}
	if len(data) == 1 { //если указан SIZE из одного значения- оно и будет максимальным
		return data[0]
	}

	maxi := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxi {
			maxi = data[i]
		}
	}
	return maxi
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	// выполняем те же первичные проверки, что и в maximum
	if len(data) <= 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	slicik := make([]int, CHUNKS) // слайс с максимальными значениями кусков
	piece := len(data) / CHUNKS   // размер куска, для обработки горутиной
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		begin := i * piece //начальное значение куска
		//считаем конечное значения куска, ибо оно может быть больше остальных
		var end int
		end = begin + piece
		if i == CHUNKS-1 { //else убрал
			end = len(data) // последний слайс считаем до последнего элемента
		}
		wg.Add(1)
		//в горутинах вызываем функцию maximum
		go func() {
			slicik[i] = maximum(data[begin:end])
			wg.Done()
		}()
	}
	wg.Wait()
	return maximum(slicik) //также найдем максимум через функцию maximum
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)

	fmt.Print("Ищем максимальное значение в один поток ")
	// ваш код здесь
	start := time.Now()
	maks := maximum(slice)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maks, elapsed.Microseconds()) //исправил ms на us, т.к. в задании вывод в микросекундах)

	fmt.Printf("Ищем максимальное значение в %d потоков ", CHUNKS)
	// ваш код здесь
	start = time.Now()
	maks = maxChunks(slice)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maks, elapsed.Microseconds())

}
