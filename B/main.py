def mod_inverse(a, p):

    return pow(a, p - 2, p)

def binomial_coefficient_mod(n, k, p):
    if k > n:
        return 0

    factorial = [1] * (n + 1)
    inverse_factorial = [1] * (n + 1)


    for i in range(2, n + 1):
        factorial[i] = factorial[i - 1] * i % p


    inverse_factorial[n] = mod_inverse(factorial[n], p)
    for i in range(n - 1, 0, -1):
        inverse_factorial[i] = inverse_factorial[i + 1] * (i + 1) % p


    return factorial[n] * inverse_factorial[k] % p * inverse_factorial[n - k] % p


n, k = map(int, input().split())

print(binomial_coefficient_mod(n, k, 9929))
