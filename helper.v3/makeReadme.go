package main

// func makeREADME(c Categories, s Solveds) {
// 	file := "README.md"
// 	os.Remove(file)

// 	template := `%s

// ## 答题进度
// %s

// ## 参考解答
// %s

// %s
// `
// 	headFormat := read("README_HEAD.md")
// 	head := fmt.Sprintf(headFormat, username, username, ranking, username)
// 	count := c.String()
// 	solved := s.String()
// 	tail := read("README_TAIL.md")
// 	content := fmt.Sprintf(template, head, count, solved, tail)

// 	err := ioutil.WriteFile(file, []byte(content), 0755)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
