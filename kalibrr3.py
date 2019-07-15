def nword(string, word):
    normal = string
    listed = list(normal)
    listed.reverse()
    reverse = ''.join(listed)

    n = len(word)
    count = 0

    idx = 0
    start = 0
    while idx >= 0:
        idx = normal.find(word, start)
        if idx >= 0:
            count += 1
        start = idx+1
    idx = 0
    start = 0
    while idx >= 0:
        idx = reverse.find(word, start)
        if idx >= 0:
            count += 1
        start = idx+1

    return count


def fdiags(grid):
    n = len(grid[0])
    m = len(grid)

    diags = []
    Ll = [(1, n-1)] + [(0, i+1) for i in range(n-1)]
    Lr = [(1, 0)] + [(0, i) for i in range(n-1)]
    for l, r in zip(Ll, Lr):
        diags.append(''.join([grid[l[0]+j][l[1]-k]
                              for j, k in zip(range(m-l[0]), range(l[1]+1))]))
        diags.append(''.join([grid[r[0]+j][r[1]+k]
                              for j, k in zip(range(m-r[0]), range(n-r[1]))]))

    return diags


def wsearch(grid, word):

    n = len(grid[0])
    m = len(grid)
    if n == 1:
        from_cols = sum([nword(''.join(i), word) for i in zip(*grid)])
        return from_cols
    if m == 1:
        from_rows = sum([nword(''.join(i), word) for i in grid])
        return from_rows
    from_diags = sum([nword(i, word) for i in fdiags(grid)])
    from_cols = sum([nword(''.join(i), word) for i in zip(*grid)])
    from_rows = sum([nword(''.join(i), word) for i in grid])
    return from_diags+from_cols+from_rows


T = int(input())
C = 0
result = []
while C < T:
    N = int(input())
    M = int(input())
    grid = [list(input()[:M]) for j in range(N)]
    word = input()
    result.append(wsearch(grid, word))
    C = C + 1

C = 0
while C < T:
    print("Case ", C + 1, ": ", result[C], sep='')
    C = C + 1
