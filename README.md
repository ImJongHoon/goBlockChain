## 단순 블록체인 구현 중 배우는 것


#### 1. 블록체인의 hash함수는 단방향 함수이며 결정론적 함수이다.
#### 2. 모든 블록이 data를 기반으로 한 hash 값을 가지고 있으므로 중간에서 변경될 경우 변경된 부분 이후에 모든 hash가 바뀜을 알 수 있다.

#### 3. Hash 알고리즘으로 SHA256알고리즘을 많이 사용한다.
#### go에 이미 구현되어있는 이 알고리즘은 sha256.Sum256() 으로 사용하며 인자와 반환값 모두 byte 타입이다.
#### 이 때, String도 byte의 slice의 일종이다.


#### Hash를 해줄 때 %x 로 16진수를 사용해야한다.


#### 4. Sprint함수
```
fmt.Sprint("%", hash)

이 함수는 print하는 대신 %이후에 오는 자료형을 String으로 반환함.
```

#### 5. len(slice명) 을 넣는 것으로 slice의 길이를 반환받을 수 있다.

#### 6. slice의 append함수는 메소드형식이 아닌 반환 형식으로 사용한다.
```
b.blockList = append(b.blockList, newBlock)
```

#### 7. Refectoring
함수는 하나당 하나의 역할을 하도록.
또 package를 사용해서 정리해도 좋다.