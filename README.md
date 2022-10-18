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

#### 8. 싱글톤 패턴(single ton pattern)
객체의 instance가 오직 하나만 생성되는 패턴.
딱히 문법으로 있는 게 아니라 프로그래머들의 편의로 만든 개념인듯. 코드를 짜서 구현함.
생성자랑 비슷한 개념으로 느껴짐

#### 9. nill 이란?
Go 언어에서 pointer, interface, maps, slices, function type들의 zero value 이다.
zero value란 명시적인 초기값이 없이 선언된 변수가 기본적으로 가지게 되는 값이다.
예를 들어 string은 "" 이 zero value

#### 10. sync package란?
go 언어는 goroutine을 만들기 쉽다.
즉 프로그램을 병렬적으로 실행 가능하다는 뜻이며 sync패키지가 동기적으로 처리해야하는 부분을 제대로 처리하게 도와준다.

*동기적이란?
동기 방식은 서버에 요청을 보낼때 응답이 돌아와야 다음 작업을 수행할 수 있다. 즉 A작업이 모두 진행
되기 전까지는 B작업은 대기해야함. 여기서는 서버가 프로그램을 관리하는 무언가.

비동기 방식은 반대로 요청을 보낼때 응답 상태와 상관없이 다음 동작을 수행할 수 있다. 즉 A작업이 시작
되었을때 동시에 B작업이 실행된다. A작업은 결과값이 나오는대로 출력된다.


비동기 방식을 쉽게 구현할 수 있는 golang 특성상 오히려 동기적 방식으로 진행시키고 싶은 부분을 명확히 지정해주는 기능이 sync package.

#### Once 패키지 중 Do 함수는 확실히 단 한번만 호출되도록 해주는 function 이다.
즉! 운영체제로 따지면 critical section을 확실히 정해주는 셈.