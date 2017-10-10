# helper

## 功能

1. 按照程序中设置的模板，生成习题解答文件夹，包括:
    1. README.md
    1. [题目].go
    1. [题目]_test.go
1. 统计已经解答题目，在`LeetCode-in-Go/README.md`中自动生成
    1. 答题进度，并生成图标
    1. LeetCode 排名，并生成图标
    1. 按难度分类的答题进度统计表
    1. 已完成题目的目录
    1. 无法使用 Go 解答的题目列表
1. 检查题目是否能够用 Go 语言解答，不能解答题目的题号，会保存起来。

## 使用方式

1. 在命令行的本目录下，运行`make`，会自动生成程序。
1. 在命令行的`LeetCode-in-Go`目录下，运行`./helper`，可以查看程序的帮助文件。

## 配置方法

1. 在`.gitignore`中，添加一行`*.toml`。
1. 在`LeetCode-in-Go`目录下，添加文本文件`leetcode.toml`。
1. 把以下内容中的`test`分别修改为你的 leetcode `用户名`和`密码`后，复制到`leetcode.toml`中。

```toml
Login="test"
Password="test"
```
