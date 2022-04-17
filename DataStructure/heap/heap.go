package heap

/*
	Heap堆(最大堆):
		堆序性:大根堆的左右节点都比根小,也就是头节点是最大值(还有小根堆)
		节点下标为i时，左子节点下标为2i+1,右子节点下标为2i+2,父节点是i/2 (3/2 = 2/2 = 1,整除)
		所有的叶子应处于第h/h-1层(h是树的高度),也就是说应是一棵完全二插树
*/
type ElementType int
type Heap []ElementType

// Create 创建一个Heap堆
func Create() Heap {
	return make(Heap, 0)
}

// Insert 插入一个节点在Heap的尾部,然后Shift Up(上滤)最后一个节点,保证堆序性
func (heap *Heap) Insert(data ElementType) {
	*heap = append(*heap, data)
	heap.shiftUp(len(*heap) - 1)
}

// GetMax 弹出最大值,也就是返回0节点,用最后一个节点代替它然后下滤
func (heap *Heap) GetMax() ElementType {
	max := (*heap)[0]
	// 交换最大和最小
	heap.swap(0, len(*heap)-1)
	// 删除最后一个元素(也就是最大值,被swap到最后了)
	*heap = (*heap)[:len(*heap)-1]
	// 下滤第一个节点,保证堆序性
	heap.shiftDown(0)
	return max
}

// shiftUp 上滤,将一个(最后一个)节点向上移动到正确的位置(最大堆)
func (heap *Heap) shiftUp(index int) {
	for index > 0 {
		// 父节点
		fatherIndex := index / 2
		if (*heap)[fatherIndex] < (*heap)[index] {
			// 父节点比当前节点小,把当前节点和父节点交换
			heap.swap(fatherIndex, index)
			index /= 2
		} else {
			// 父节点比当前节点大 退出
			break
		}
	}
}

// shiftDown 下滤,将一个(第一个)节点向下移动到正确的位置(最大堆)
func (heap *Heap) shiftDown(index int) {
	for {
		// 最大下标
		maxIndex := len(*heap) - 1
		// 左右儿子节点的下标
		leftChild := index*2 + 1
		rightChild := index*2 + 2
		if leftChild <= maxIndex && rightChild <= maxIndex {
			// 有两个儿子节点,判断是否都大于当前节点
			if (*heap)[leftChild] > (*heap)[index] && (*heap)[rightChild] > (*heap)[index] {
				// 都大于当前节点,找一个最大的交换
				if (*heap)[leftChild] > (*heap)[rightChild] {
					// 左节点较大
					heap.swap(leftChild, index)
					index = leftChild
				} else {
					// 右节点较大
					heap.swap(rightChild, index)
					index = rightChild
				}
			} else {
				// 都小于当前节点,直接退出
				break
			}
		} else if leftChild <= maxIndex {
			// 只有左儿子节点
			if (*heap)[leftChild] > (*heap)[index] {
				// 左儿子比当前节点大,交换
				heap.swap(leftChild, index)
				index = leftChild
			} else {
				// 左儿子比当前节点小,退出
				break
			}
		} else {
			// 没有儿子节点,直接退出(因为是完全二插树所以不存在有右儿子但是没有左儿子的情况)
			break
		}
	}
}

// swap 交换两个元素
func (heap *Heap) swap(index1, index2 int) {
	(*heap)[index1], (*heap)[index2] = (*heap)[index2], (*heap)[index1]
}
