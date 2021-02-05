package array_list

var (
	items []string //数组
	n     = 0      //数组大小
	head  = 0      //队头下标
	tail  = 0      //队尾下标
)

func CircularQueue(capacity int) {
	n = capacity

}

// 入队
func enqueue(item string) bool {
	// 队列满了
	if (tail+1)%n == head {
		return false
	}
	items[tail] = item
	tail = (tail + 1) % n
	return true
}

func main() {

}
