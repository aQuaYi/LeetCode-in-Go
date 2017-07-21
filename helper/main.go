package main

const (
	// USER 是你在leetcode.com的用户名
	// 登录后，在页面右上角可以看到
	USER = "aQuaYi"
)

func main() {
	categories := []string{
		"algorithms",
		"shell",
		"draft",
		"database",
		"system-design",
		"operating-system",
	}

	j := getJSON(url(categories[1]))
	c := getCategory(j)
	c.run()
}
