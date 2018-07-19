package problem0458

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// 15 分钟中毒身亡，总共 60 分钟的测试时间的话
	// 一只猪可以测试 5 组样品，从 0 分钟的时候喝一组，每过 15分钟没死就继续喝下一组
	// 猪在 60 分钟内死亡的话，就是刚刚喝的那一组有毒，没死的话，就是第 5 组有毒。
	base := minutesToTest/minutesToDie + 1
	// 然后，可以把猪的个数看成是空间的维度
	// 把 buckets 看成是多维空间中正立方体的体积
	// 正立方体的边长就是 base
	power := 1
	res := 0

	for power < buckets {
		power *= base
		res++
	}

	return res
}
