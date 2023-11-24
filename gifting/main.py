import random
import sys

source = sys.argv[1]
gifters = []

with open(source, "r") as f:
    for line in f:
        gifters.append(line.strip())

if len(gifters) % 2 != 0:
    print("Number of gifters must be even")
    sys.exit(1)

random.shuffle(gifters)

pairs = []
for i in range(len(gifters)//2):
    pairs.append((gifters[i], gifters[-i-1]))

for pair in pairs:
    print(f"{pair[0]}, {pair[1]}")
