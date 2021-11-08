module mx_im

go 1.16

require (
	github.com/Shopify/sarama v1.19.0
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.7.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/olivere/elastic/v7 v7.0.29 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	go.etcd.io/etcd v0.0.0-20200402134248-51bdeb39e698
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/tools v0.0.0-20210106214847-113979e3529a // indirect
	google.golang.org/grpc v1.33.2
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c

)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
