# makefile

# 项目名称
PROJECT_NAME := EchoFrame Auto

# Go 命令
GO := go

# 帮助命令
.PHONY: help
help:
	$(GO) run main.go help

# 生成代码
.PHONY: code
code:
	$(GO) run main.go code

# 生成handler代码
.PHONY: h
h:
	$(GO) run main.go handler

# 生成service代码
.PHONY: s
s:
	$(GO) run main.go service

# 生成dto代码
.PHONY: d
d:
	$(GO) run main.go dto