package main

// 如果需要保护代码片段，可将代码块重构成匿名函数，如此可确保后续代码被执行。

func test(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()

		z = x / y
		return
	}()

	println("x / y =", z)
}

func main() {
	test(4, 2)
	test(4, 0)
}
