# Makefile
all: hooks

.PHONY: hooks
hooks:
	@git config --local core.hooksPath .githooks/