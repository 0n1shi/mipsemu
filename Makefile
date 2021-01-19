BIN = bin
SRC = src
EXAMPLE_DIR = example

DOCKER = docker run --rm -v $(shell pwd):/workdir multiarch/crossbuild
CROSS_ENV_DIR = /usr/mipsel-linux-gnu/bin

CC = $(DOCKER) $(CROSS_ENV_DIR)/cc
AS = $(DOCKER) $(CROSS_ENV_DIR)/as
LD = $(DOCKER) $(CROSS_ENV_DIR)/ld
COPY = $(DOCKER) $(CROSS_ENV_DIR)/objcopy
DUMP = $(DOCKER) $(CROSS_ENV_DIR)/objdump

.PHONY: build example example_elf example_bin

all: build
	$(shell which go) build -o $(BIN)/emu ./cmd/emu/

build:


example:
	$(CC) -ffreestanding -nostdlib -march=mips1 -c $(EXAMPLE_DIR)/main.c -o $(EXAMPLE_DIR)/main.o
	$(LD) -T $(EXAMPLE_DIR)/linker.ld -o $(BIN)/main.elf $(EXAMPLE_DIR)/main.o
	$(COPY) -O binary --only-section=.text --only-section=.data $(BIN)/main.elf $(BIN)/main

example_as:
	$(AS) -march=mips1 -o $(EXAMPLE_DIR)/main.o $(EXAMPLE_DIR)/main.S
	$(LD) -T $(EXAMPLE_DIR)/linker.ld -o $(BIN)/main.elf $(EXAMPLE_DIR)/main.o
	$(COPY) -O binary --only-section=.text --only-section=.data $(BIN)/main.elf $(BIN)/main

example_elf:
	$(DUMP) -D $(BIN)/main.elf

example_bin:
	$(DUMP) -m mips -b binary --endian=little -D $(BIN)/main
	