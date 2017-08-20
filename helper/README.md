# helper

## 功能
1. 按照程序中设置的模板，生成习题解答文件夹，包括:
    1. README.md
    1. [题目].go
    1. [题目]_test.go
1. 统计已经解答题目，在`LeetCode-in-Golang/README.md`中生成
    1. 统计报表
    1. 已完成题目的目录

## 使用方式 
1. 按照下方的说明，配置helper。
1. 在命令行的本目录下，运行`make`
    
## 配置方法
1. 在`.gitignore`中，添加一行`*.toml`。
1. 在`LeetCode-in-Golang`目录下，添加文件`leetcode.toml`。
1. 把以下内容中的`test`分别修改为你的leetcode用户名和密码后，复制到`leetcode.toml`中。
```toml
Login="test"
Password="test"
```
