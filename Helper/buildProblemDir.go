package main

// func buildProblemDir(num string) {
// 	log.Printf("~~ 开始生成第 %s 题的文件夹 ~~\n", num)
// //
// 	var err error
// 	var problemNum int
// //
// 	// 获取问题编号
// 	problemNum, err = strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		log.Fatalln("无法获取问题编号：", err)
// 	}
// 	if problemNum <= 0 {
// 		log.Fatalln("题目编号应该 >0")
// 	}
// //
// 	// 需要创建答题文件夹
// 	lc, err := readLeetCodeRecord()
// 	if err != nil {
// 		log.Fatalln("读取 LeetCode 记录失败: ", err)
// 	}
// //
// 	makeProblemDir(lc.Problems, problemNum)
// //
// 	log.Printf("~~ 第 %s 题的文件夹，已经生成 ~~\n", num)
// }
