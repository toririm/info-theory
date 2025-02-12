import math

# p = 0.01
pu = 1
pd = 100

m = int(input())
n = 2**m - 1
print("n", n)
k = n - m
print("k", k)

pe = 0

for i in range(2, n + 1):
    if i > 2:
        print(" + ", end="")
    j = math.comb(n, i) * pu**i * (pd - pu) ** (n - i)
    print(f"{j/pd**n:.10f}", end="")
    pe += j
print()

pe /= pd**n

print(f"{pe:.10f}")
