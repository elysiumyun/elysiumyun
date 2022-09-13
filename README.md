# Layout
[golang-standards/project-layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

# Usage

- build

```bash
make
```

- run

```bash
./bin/elysium
```
# Workflow

## githook
[pre-commit](https://pre-commit.com/)

> go install golang.org/x/tools/cmd/goimports@latest

## github actions
[codeql-action](https://github.com/github/codeql-action)

[https://golangci-lint.run/](https://golangci-lint.run/)

[https://github.com/golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action)

## Git 贡献提交规范

- 参考 [vue](https://github.com/vuejs/vue/blob/dev/.github/COMMIT_CONVENTION.md) 规范 ([Angular](https://github.com/conventional-changelog/conventional-changelog/tree/master/packages/conventional-changelog-angular))

  - `feat` 增加新功能
  - `fix` 修复问题/BUG
  - `style` 代码风格相关无影响运行结果的
  - `perf` 优化/性能提升
  - `refactor` 重构
  - `revert` 撤销修改
  - `test` 测试相关
  - `docs` 文档/注释
  - `chore` 依赖更新/脚手架配置修改等
  - `workflow` 工作流改进
  - `ci` 持续集成
  - `types` 类型定义文件更改
  - `wip` 开发中

## TimeZone Support
待定: [elysiumyun/timezone](https://github.com/elysiumyun/timezone.git)
