with open("day_6.data") as FILE:
    data = list(map(int, FILE.read().split(",")))

final = dict().fromkeys([i for i in range(9)], 0)
for i in data:
    final[i] += 1

for day in range(256):
    curr = final[0]
    for i in range(1, 9):
        final[i-1] = final[i]
    final[8] = curr
    final[6] += curr

print(sum(list(final.values())))