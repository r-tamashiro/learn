
# 素数の計算
def primse(maxvalue):
    ret = []
    for v in range(2, maxvalue):
        for p in ret:
            if v %p == 0:
                break
        else:
            ret.append(v)
    return ret

print(f'{primse(100)}')
