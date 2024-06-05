from math import factorial, comb

def binomial_coefficient(n, k):
    return comb(n,k)

def find_n_k(result):
    n_k_pairs = []
    for n in range(result + 1):
        for k in range(n + 1):
            if binomial_coefficient(n, k) == result:
                n_k_pairs.append((n, k))
    return n_k_pairs

def main():
    t = int(input())
    for _ in range(t):
        result = int(input())
        n_k_pairs = find_n_k(result)
        if n_k_pairs:
            print(len(n_k_pairs))
            for n, k in sorted(n_k_pairs):
                print(f"({n}, {k})",end=" ")
            print()
        else:
            print(0)

if __name__ == "__main__":
    main()
