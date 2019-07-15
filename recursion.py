def pow(base, exp):
    if exp > 0:
        return base * pow(base, exp - 1)
    elif exp == 0:
        return 1

base = int(input("Enter the base number: "))
exp = int(input("Enter the power of the number: "))
print(base, "^", exp, " = ", pow(base,exp), sep = '')