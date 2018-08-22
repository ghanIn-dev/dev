# ELK

- Elasticsearch + Logstash + Kibana + Filebeat
- Elasticsearch는 아파치 루신기반 실시간 분산 검색엔진
- Logstash는 각종 로그를 가져와 json 형태로 가공하여 Elasticsearch로 전달
- Kibana는 Elasticsearch에 저장된 데이터를 사용자가에 차트형태로 보여주는 시각화 솔류션
- Filebeat는 로그내용을 수집하여 Logstash로 전달하는 역할을 함

![elk그림](https://okdevtv.com/md/elk/images/elk_arch.jpg)

## 1. Filebeat
### filebeat.yml 편집
- -path 설정 (와일드카드 * 이용가능)
- tag 는 파일을 보낼때 붙이는 분류기준과 같다. 받는 쪽 (logstash) 에서 tag로 구분하여 처리한다.
- General의 name은 서버의 이름을 기록한다.
- elasticsearch가 default로 활성화 되어있는데 이를 주석처리하여야한다. 우리는 logstash로 전송할 것 이기때문에 output.logstash를 활성화시키고 host를 입력한다.
- include_lines에서 정규표현식으로 특정줄만 출력되도록 할 수 있다. ex( 'reply')
- 윈도우에 서비스로 등록해두어야한다. 서비스를 등록하기위해선 install-service-filebeat 를 우선 실행해야한다.

## 2. logstash
- logstash는 30번(10.27.13.30)에 설치되어있다.
- 설정파일은 /etc/logstash/conf.d에 생성해서 사용한다.
- 설정파일은 input, filter, output 구조로 구성되어있다.
- 설정파일 내용 중 index를 설정할 수 있는데 index는 반드시 소문자만 사용한다. index는 로그분류기준이된다.
- systemctl start logstash : 상시실행가능

### logstash 문법
- disect{mapping =>{"message" => "%{name}}}  로그내용을 매핑하는 과정
- mutate{ convert => ["mid", "integeer"]} 매핑한 값들의 자료형을 선언
- mutate remove에서 timestamp는 지우면안됨 에러남
- output{if "filebeat-tag" in [tags]{elasticsearch{hosts=>[""]index=>"index-name"}} 

### 해야될 것
- doc-type 붙이기 
    - logstash add_field 찾아보기
    - tag로 각 데이터 분류해서 파싱하기

## 3. elastcisearch
- 엘라스틱서치헤드 10.27.13.31/9200

## 4. kibana
- discover탭에서 index를 찾아서 클릭하면 확인이 가능하다.
- kibana 사용시 주의할 점
    - 화면 오른쪽위에 특정 시간을 설정하는 곳에서 today로 해놔야 실시간 반영정보를 확인할 수 있다.......기본은 today로 ...




