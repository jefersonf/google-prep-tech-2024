package module3

type RLEIterator struct {
	encoding     []int
	currentIndex int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding, currentIndex: 0}
}

func (this *RLEIterator) Next(n int) int {
	for this.currentIndex < len(this.encoding) && this.encoding[this.currentIndex] == 0 {
		if this.currentIndex+2 >= len(this.encoding) {
			break
		}
		this.currentIndex += 2
	}

	var lastVal int = -1
	for this.currentIndex < len(this.encoding) && n > this.encoding[this.currentIndex] {
		n -= this.encoding[this.currentIndex]
		this.encoding[this.currentIndex] = 0
		lastVal = this.encoding[this.currentIndex+1]
		if this.currentIndex+2 >= len(this.encoding) {
			break
		}
		this.currentIndex += 2
	}

	if n <= this.encoding[this.currentIndex] {
		this.encoding[this.currentIndex] -= n
		n = 0
		lastVal = this.encoding[this.currentIndex+1]
	}

	if n != 0 {
		return -1
	}

	return lastVal
}
