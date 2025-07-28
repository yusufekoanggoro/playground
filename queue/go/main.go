package main

import "fmt"

// Queue struct dengan slice
type Queue struct {
	items []int
}

// Enqueue menambahkan item ke queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue mengambil item dari depan queue
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // Queue kosong
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty mengecek apakah queue kosong
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size mengembalikan jumlah item di queue
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	// q := Queue{}

	// // Enqueue
	// q.Enqueue(10)
	// q.Enqueue(20)
	// q.Enqueue(30)

	// fmt.Println("Queue sekarang:", q.items)

	// // Dequeue
	// for !q.IsEmpty() {
	// 	item, _ := q.Dequeue()
	// 	fmt.Println("Dequeued:", item)
	// }

	// fmt.Println("Queue setelah dikosongkan:", q.items)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice[1:])
}
