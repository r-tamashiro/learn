import time, datetime
from concurrent.futures import ThreadPoolExecutor

def task(num):
    start = datetime.datetime.now()
    time.sleep(2)
    print('number: ', num)
    return (num, start, datetime.datetime.now())

# スレッドを2つ起動
with ThreadPoolExecutor(2) as e:
    ret = e.map(task, [1, 2, 3, 4, 5])

for (num, start, end) in ret:
    print(f'#{num} start: {start} {(end - start).total_seconds()}')
