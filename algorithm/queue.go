// Trong bài queue này thì sẽ implement
// enqueue: thêm phần tử và cuối (0)
// dequeue: lấy phần tử ở đầu (-1)
// front: đọc phần tử ở đầu của queue
// back: đọc phần tử ở cuối của queue
// length: số phần tử của queue
// clear
// print

package main

import "fmt"

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(element interface{}) {
	fmt.Println("The data in Queue: ", q.data)
	// append in go: first element is slice, and the second one is value
	// append into slice
	// Vì trong hàm enqueue ta muốn append vào data ở vị trí 0 và đẩy phần tử còn
	// lại lùi về sau --> append vào 1 slice mới với elem 0 = element rồi mới append
	// toàn bộ dữ liệu của q vào
	q.data = append([]interface{}{element}, q.data...)
	// q.data = append(q.data, element)
}

func (q *Queue) Dequeue() interface{} {
	leng := len(q.data)
	fmt.Println("Number of element in Queue is ", leng)
	if leng > 0 {
		elem := q.data[leng-1]
		q.data = q.data[:leng-1]
		return elem
	}
	return nil
}

func main() {
	q := Queue{}
	q.Enqueue("Do Tieu Thien")
	q.Enqueue("Do Tieu Khoi")
	q.Enqueue("Le Tran Nhat Linh")
	fmt.Println(q)

	elem := q.Dequeue()
	fmt.Println("The value of Dequeue is ", elem)

}