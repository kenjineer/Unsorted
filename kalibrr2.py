def test(A, B, K):
    X = 0
    while A <= B:
        if A % K == 0:
            X = X + 1
            A = A + 1
        else:
            A = A + 1
    return X

T = int(input())
C = 0
A = []
B = []
K = []
X = []
while C < T:
    A.append(int(input()))
    B.append(int(input()))
    K.append(int(input()))
    X.append(test(A[C], B[C], K[C]))
    C = C + 1

C = 0
while C < T:
    print("Case ", C + 1, ": ", X[C], sep='')
    C = C + 1
