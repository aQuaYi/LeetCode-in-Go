# helper

## 功能
1. 按照程序中设置的模板，生成习题解答文件夹，包括:
    1. README.md
    1. [题目].go
    1. [题目]_test.go
1. 统计已经解答题目，在`LeetCode-in-Golang/README.md`中生成
    1. 统计报表
    1. 已完成的题目的目录
## 使用方式 
1. 按照下方的说明，配置helper。
1. 在命令行的本目录下，运行`make`
    
## 配置方法
1. 设置main.go中的USER常量为你的leetcode.com的用户名。
1. 把你的leetcode.com的cookie，保存到LeetCode-in-Golang目录下，文件名为`leetcode.cookie`。
    1. 安装Chrome浏览器，及其扩展程序EditThisCookie
    1. 设置EditThisCookie，选项 → 选择cookies的导出格式 → "Netscape HTTP Cookie File"
    1. 登录LeetCode.com，点击EditThisCookie扩展，点击"导出Cookies"按钮。
    1. 把剪切板中内容，保存到文件。