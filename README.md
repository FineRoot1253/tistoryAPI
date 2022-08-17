# Tistory API

본 패키지는 셀레니움을 포함하고 있습니다.  
이 패키지의 기능중 셀레니움을 사용하고 싶지 않다면   
직접 Authorization Code를 구하여 서비스를 생성하시면 됩니다.  
개인적으로 혼자 쓸려다 그냥 배포해서 편하게 쓸려고 만들어본 패키지입니다.

## 주의점
**셀레니움 사용시 발생한 불이익에 대해 전혀 책임드릴 수 없습니다.**  
**이점 양해 부탁드리겠습니다.**  
~~이걸 저 혼자 쓸려고 만든 저 자신도 책임을 못지는 걸요...~~  
내부 셀레니움 서비스는 크롬드라이버만 지원합니다.  
크롬드라이버 사용시 사용중인 크롬이 있다면 버전을 맞춰서 사용해야 에러가 나오지 않습니다.  
만약 외부망 셀레니움(도커, aws에 올려둔 셀레니움등등)을 사용하겠다고 한다면  
그냥 셀레니움을 직접 구현해서 쓰시는거나 직접 AuthorizationCode를 구하시는 추천드립니다. 

## Quick Guide

````
userData := model.UserData{
    
}
service := tistoryAPI.NewService(context.Background(), model.UserData{})
````

MIT License

Copyright (c) [2022] [FineRoot1253]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.