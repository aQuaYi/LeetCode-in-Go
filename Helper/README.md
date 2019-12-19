# helper

## 配置方法

1. 在`.gitignore`中，添加一行`*.toml`
1. 在`LeetCode-in-Go`目录下，添加文本文件`config.toml`。
1. 把以下内容复制到`config.toml`中。
1. 把`config.toml`中的`test`分别修改为你的 leetcode `用户名`和`密码`。
1. 去 leetcode 登录后，把网站 cookie 复制 (有洁癖者只要复制 LEETCODE_SESSION 就够了) 并替换 `config.toml`中的`Cookie`。

```toml
Username="test"
Password="test"
Cookie="LEETCODE_SESSION=XXXXXXXXX;"
```