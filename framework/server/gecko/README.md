# Gecko

Gecko 是一个手撸 HTTP 开发框架的实验。它是著名 web 框架 Gin 和 Echo 的简易版本。

Gecko 是指守宫，一种蜥蜴，你也可以把它当作是 Gin + Echo 的发音组合。

Gecko is an experiment of writing an HTTP Server framework. It is a simple version of Gin and Echo framework.

The name of Gecko means a lizard. It can also be considered as a combination of Gin and Echo : )

## Feature

1. 使用了 `Trie` 树作为路由查找表。
2. 支持两种 `wildcard` 路由。使用 `:` 表示一个路由参数，在路径末尾使用 `*` 表示一个任意匹配。
