# Calling Go Functions from Other Languages using C Shared Libraries

다음의 [github](https://github.com/vladimirvivien/go-cshared-examples)를 참고하여 궁금한 부분을 추가적으로 연습해본 내용입니다. 

- 의문 : 개인이 정의한 패키지의 함수로 작성하여도 내보내기가 정상적으로 될까?
- 작성된 예제로 위 내용을 확인 하였습니다. 

## How to use (windows)
#### 1. 패키지는 -buildmode = c-shared 빌드 플래그를 사용해 컴파일되어 공유 객체 바이너리로 만들어 집니다.
#### 완료되면 컴파일러는 C 헤더 파일(go-functions.h) 및 공유 객체 파일(go-functions.so)을 출력합니다.
```
$ go build -o go-functions.so -buildmode=c-shared ./exports/targets.go
```
#### 2. 작성한 C 코드에서 공유 객체 라이브러리를 지정하여 컴파일

```
gcc -o test go-functions.so test.c
```
#### 3. 실행 ! 
```
./test.exe
```
#### The result
```
Using awesome lib from C:
rpc.Dial(tcp, 127.0.0.1:7500)
dial tcp 127.0.0.1:7500: connectex: No connection could be made because the target machine actively refused it.
awesome.Add(12,99) = -1
awesome.Sub(12,99) = -87
awesome.Mul(12,99) = 1188
awesome.Cosine(1) = 0.540302
awesome.Sort(77,12,5,99,28,23): 5,12,23,28,77,99,
Hello from C!
```

## The Go Code
#### targets.go 
외부에서 사용할 함수들 정의

#### functions.go
targets.go 에서 불러서 사용할 다른 패키지에 정의된 함수

#### 고려사항 
- https://golang.org/cmd/cgo
- Not all Go types can be mapped to C types in a useful way. Go struct types are not supported; use a C struct type. Go array types are not supported; use a C pointer.