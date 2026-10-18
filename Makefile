# SPDX-FileCopyrightText: NONE
# SPDX-License-Identifier: Unlicense

CC := cc
CFLAGS := -Wall -Wextra -std=c99 -ldiscord -lcurl -lpthread

SRC_DIR := src
OUT := out

.PHONY: all
all:
	mkdir -p out
	$(CC) -O2 $(CFLAGS) -o $(OUT)/monica $(SRC_DIR)/*.c

.PHONY: debug
debug:
	mkdir -p out
	$(CC) -O0 -g $(CFLAGS) -o $(OUT)/monica $(SRC_DIR)/*.c -fsanitize=address,undefined -fno-omit-frame-pointer
