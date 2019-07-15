def g(str):
    i = 0
    new_str = ""
    while i < len(str) - 1:
        new_str = new_str + str[i + 1]
        i = i + 1
    return new_str

def f(str):
    if len(str) == 0:
        return ""
    elif len(str) == 1:
        return str
    else:
        return f(g(str)) + str[0]

def h(n, str):
    while n != 1:
        if n % 2 == 0:
            n = n/2
        else:
            n = 3*n + 1
        str = f(str)
    return str

def pow(x, y):
    if y == 0:
        return 1
    else:
        return x * pow(x, y-1)

#print(h(1, "fruits")+h(2, "fruits")+h(5, "fruits"))
print(h(pow(2, 22), "fruits"))