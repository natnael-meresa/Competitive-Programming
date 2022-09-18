package main

type maxHeap struct {
	array []int
}

func (m *maxHeap) buildHeap() {
	startIndex := (len(m.array) / 2) - 1

	for i := startIndex; i >= 0; i-- {
		m.downHeapify(i)
	}

}

func (m *maxHeap) downHeapify(i int) {
	highest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < len(m.array) && m.array[left] > m.array[highest] {
		highest = left
	}

	if right < len(m.array) && m.array[right] > m.array[highest] {
		highest = right
	}

	if highest != i {
		m.swap(highest, i)
		m.downHeapify(highest)
	}
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func (m *maxHeap) swap(pI, i int) {
	m.array[pI], m.array[i] = m.array[i], m.array[pI]
}

func (m *maxHeap) heapifyUP(i int) {
	for m.array[i] > m.array[parentIndex(i)] {
		m.swap(parentIndex(i), i)
		i = parentIndex(i)
	}
}

func (m *maxHeap) pop() int {
	l := len(m.array) - 1
	extracted := m.array[0]
	m.array[0] = m.array[l]
	m.array = m.array[:l]

	m.downHeapify(0)

	return extracted
}

func lastStoneWeight(stones []int) int {
	mxHeap := maxHeap{array: stones}
	mxHeap.buildHeap()

	for len(mxHeap.array) > 1 {
		max1 := mxHeap.pop()
		max2 := mxHeap.pop()
		diff := max1 - max2
		if diff > 0 {
			mxHeap.array = append(mxHeap.array, diff)
			mxHeap.heapifyUP(len(mxHeap.array) - 1)

		}

	}
	if len(mxHeap.array) > 0 {
		return mxHeap.array[0]
	}
	return 0
}

// func main() {
// 	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
// }
