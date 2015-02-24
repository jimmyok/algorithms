/*Quick union use a set of trees (a forest) to quickly find the root element and hence connectedness of two nodes.
* The runtime of this is O(N)
*
 */
package main

import "fmt"

type QU struct {
	id []int
}

func (q *QU) Init(size int) {
	for i := 0; i < size; i++ {
		q.id = append(q.id, i)
	}
}

func (q *QU) Root(i int) int {
	for i != q.id[i] {
		i = q.id[i]
	}
	return i
}

func (q *QU) Union(a int, b int) {
	i := q.Root(a)
	j := q.Root(b)
	q.id[i] = j
}

func (q *QU) Connected(a int, b int) bool {
	return q.Root(a) == q.Root(b)
}

func main() {
	qu := QU{}
	qu.Init(10)
	fmt.Printf("Array initialized. %v", qu.id)
	qu.Union(3, 4)
	qu.Union(4, 8)
	fmt.Printf("\n\n%v is Connected to %v %v", 3, 8, qu.Connected(3, 8))
}
