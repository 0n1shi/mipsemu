BIN = bin
SRC = src
TEST_DIR = $(SRC)/test

DOCKER = docker run --rm -v $(shell pwd):/workdir multiarch/crossbuild
CROSS_ENV_DIR = /usr/mipsel-linux-gnu/bin

CC = $(DOCKER) $(CROSS_ENV_DIR)/cc 
LD = $(DOCKER) $(CROSS_ENV_DIR)/ld
COPY = $(DOCKER) $(CROSS_ENV_DIR)/objcopy
DUMP = $(DOCKER) $(CROSS_ENV_DIR)/objdump

.PHONY: test test_elf test 

test:
	$(CC) -ffreestanding -nostdlib -c $(TEST_DIR)/test.c -o $(TEST_DIR)/test.o
	$(LD) -T $(TEST_DIR)/linker.ld -o $(BIN)/test.elf $(TEST_DIR)/test.o
	$(COPY) -O binary --only-section=.text --only-section=.data $(BIN)/test.elf $(BIN)/test

elf:
	$(DUMP) -D $(BIN)/test.elf

dump:
	$(DUMP) -m mips -b binary --endian=little -D $(BIN)/test
	