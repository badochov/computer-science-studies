def score(i, o):
    i_n, i_m, i_k = i
    o_n, o_m, o_k = o
    return o_m * i_n * i_k + o_n * i_m * i_k + o_k * i_m * i_n

print(score((4,6,5),  (2,1,1)))
