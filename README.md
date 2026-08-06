# go-quote

随机输出编程名言，内置 10 条，标准库实现，零依赖。

```powershell
go run .            # 随机一条
go run . -n 3       # 随机 3 条（不重复）
go run . -list      # 列出全部
go run . -seed 42   # 指定随机种子，结果可复现
```

参数：

- `-n` 随机输出条数（默认 1，超过总数则取总数）
- `-list` 列出全部名言
- `-seed` 随机种子（默认 0 表示用当前时间）

```powershell
go build ./...
```
