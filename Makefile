SRC = sample/plus.c
DST = bin/a.out

DOCKER = docker run --rm -v $(shell pwd):/workdir multiarch/crossbuild
CROSS_ENV_DIR = /usr/mipsel-linux-gnu/bin

CC = $(DOCKER) $(CROSS_ENV_DIR)/cc
COPY = $(DOCKER) $(CROSS_ENV_DIR)/objcopy
DUMP = $(DOCKER) $(CROSS_ENV_DIR)/objdump

.PHONY: dump build

all: build

build:
	$(CC) $(SRC) -o $(DST)

dump:
	$(DUMP) -D $(DST)
