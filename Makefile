BIN = bin
SRC = src
EXAMPLE_DIR = example

DOCKER = docker run --rm -v $(shell pwd):/workdir multiarch/crossbuild
CROSS_ENV_DIR = /usr/mipsel-linux-gnu/bin

CC = $(DOCKER) $(CROSS_ENV_DIR)/gcc
AS = $(DOCKER) $(CROSS_ENV_DIR)/as
LD = $(DOCKER) $(CROSS_ENV_DIR)/ld
COPY = $(DOCKER) $(CROSS_ENV_DIR)/objcopy
DUMP = $(DOCKER) $(CROSS_ENV_DIR)/objdump

.PHONY: build ex_plus ex_mult ex_as build_ex dump_elf dump_bin

all: build

build:
	$(shell which go) build -o $(BIN)/emu ./cmd/emu/

ex_plus:
	$(CC) -ffreestanding -nostdlib -mips1 -O0  -c $(EXAMPLE_DIR)/plus.c -o $(EXAMPLE_DIR)/main.o
	@make build_ex

ex_mult:
	$(CC) -ffreestanding -nostdlib -mips1 -O0 -c $(EXAMPLE_DIR)/mult.c -o $(EXAMPLE_DIR)/main.o
	@make build_ex

ex_sub:
	$(CC) -ffreestanding -nostdlib -mips1 -O0 -c $(EXAMPLE_DIR)/sub.c -o $(EXAMPLE_DIR)/main.o
	@make build_ex

ex_div:
	$(CC) -ffreestanding -nostdlib -mips1 -O0 -mno-check-zero-division -c $(EXAMPLE_DIR)/div.c -o $(EXAMPLE_DIR)/main.o
	@make build_ex

ex_for:
	$(CC) -ffreestanding -nostdlib -mips1 -O0 -c $(EXAMPLE_DIR)/for.c -o $(EXAMPLE_DIR)/main.o
	@make build_ex

ex_as:
	$(AS) -mips1 -O0 -o $(EXAMPLE_DIR)/main.o $(EXAMPLE_DIR)/main.S
	@make build_ex

build_ex: $(EXAMPLE_DIR)/main.o
	$(LD) -T $(EXAMPLE_DIR)/linker.ld -o $(BIN)/main.elf $^
	$(COPY) -O binary --only-section=.text --only-section=.data $(BIN)/main.elf $(BIN)/main

dump_elf:
	$(DUMP) -D $(BIN)/main.elf

dump_bin:
	$(DUMP) -m mips -b binary --endian=little -D $(BIN)/main
	