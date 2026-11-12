# HTML Static Resources

本目录包含 HTML 模块使用的静态资源源文件。

## 目录说明

- `zcss.js`：前端样式/交互辅助脚本
- `zview.js`：ZView.js 前端脚本
- `skills/`：与本目录相关的 Agent 技能说明

## 资源生成与使用

静态资源源码存放于本目录，通过 `html/gen.go` 转换为 Go 源文件（生成 `html/static_data.go`），由 znet 直接输出，无需运行时读取磁盘。

在 `html` 模块目录执行：

```bash
go generate ./html
```

> 修改 `html/static/*.js` 后请重新执行 `go generate ./html` 同步生成产物。

## 构建控制

- 默认构建会包含 `zcss.js` / `zview.js` 及对应静态路由。
- 使用 `-tags nostatic` 可剔除内置静态资源与路由（适用于外部 CDN 或最小化构建场景）。

## Agent 指引

如果你是处理本目录的 Agent，请先查看并优先使用：

`./skills`

在创建、编辑或审查本目录文件时，优先遵循对应技能约定。