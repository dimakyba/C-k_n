from math import comb

t = int(input())

for _ in range(t):
    n, k = map(int, input().split())
    print(comb(n, k))
