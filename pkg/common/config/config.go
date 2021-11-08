package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"mx_im/pkg/base"
)

var Config config

type config struct {
	ServerIP string `yaml:"serverip"`

	Api struct {
		GinPort []int `yaml:"restApiPort"`
	}
	Sdk struct {
		WsPort []int `yaml:"clientSdkWsPort"`
	}

	Etcd struct {
		EtcdSchema string   `yaml:"etcdSchema"`
		EtcdAddr   []string `yaml:"etcdAddr"`
	}
	Kafka struct {
		Ws2mschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		Ms2pschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		ConsumerGroupID struct {
			MsgToMongo string `yaml:"msgToMongo"`
			MsgToMySql string `yaml:"msgToMySql"`
			MsgToPush  string `yaml:"msgToPush"`
		}
	}
	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
		DBTableName    string   `yaml:"DBTableName"`
		DBMsgTableNum  int      `yaml:"dbMsgTableNum"`
		DBMaxOpenConns int      `yaml:"dbMaxOpenConns"`
		DBMaxIdleConns int      `yaml:"dbMaxIdleConns"`
		DBMaxLifeTime  int      `yaml:"dbMaxLifeTime"`
	}
	Mongo struct {
		DBAddress           []string `yaml:"dbAddress"`
		DBDirect            bool     `yaml:"dbDirect"`
		DBTimeout           int      `yaml:"dbTimeout"`
		DBDatabase          string   `yaml:"dbDatabase"`
		DBSource            string   `yaml:"dbSource"`
		DBUserName          string   `yaml:"dbUserName"`
		DBPassword          string   `yaml:"dbPassword"`
		DBMaxPoolSize       int      `yaml:"dbMaxPoolSize"`
		DBRetainChatRecords int      `yaml:"dbRetainChatRecords"`
	}
	Redis struct {
		DBAddress     string `yaml:"dbAddress"`
		DBMaxIdle     int    `yaml:"dbMaxIdle"`
		DBMaxActive   int    `yaml:"dbMaxActive"`
		DBIdleTimeout int    `yaml:"dbIdleTimeout"`
		DBPassWord    string `yaml:"dbPassWord"`
	}

	RpcPort struct {
		RpcAuthPort        []int `yaml:"rpcAuthPort"`
	}
	RpcRegisterName struct {
		RpcAuthName               string `yaml:"rpcAuthName"`
	}
	Log struct {
		StorageLocation       string   `yaml:"storageLocation"`
		RotationTime          int      `yaml:"rotationTime"`
		RemainRotationCount   uint     `yaml:"remainRotationCount"`
		RemainLogLevel        uint     `yaml:"remainLogLevel"`
		ElasticSearchSwitch   bool     `yaml:"elasticSearchSwitch"`
		ElasticSearchAddr     []string `yaml:"elasticSearchAddr"`
		ElasticSearchUser     string   `yaml:"elasticSearchUser"`
		ElasticSearchPassword string   `yaml:"elasticSearchPassword"`
	}
	ModuleName struct {
		LongConnSvrName string `yaml:"longConnSvrName"`
		MsgTransferName string `yaml:"msgTransferName"`
		PushName        string `yaml:"pushName"`
	}
	LongConnSvr struct {
		WsPort       []int `yaml:"wsPort"`
		WebsocketMaxConnNum int   `yaml:"websocketMaxConnNum"`
		WebsocketMaxMsgLen  int   `yaml:"websocketMaxMsgLen"`
		WebsocketTimeOut    int   `yaml:"websocketTimeOut"`
	}

	Push struct {
		Tpns struct {
			Ios struct {
				AccessID  string `yaml:"accessID"`
				SecretKey string `yaml:"secretKey"`
			}
			Android struct {
				AccessID  string `yaml:"accessID"`
				SecretKey string `yaml:"secretKey"`
			}
		}
	}
	Manager struct {
		AppManagerUid []string `yaml:"appManagerUid"`
		Secrets       []string `yaml:"secrets"`
	}
	Secret           string `yaml:"secret"`
	MultiLoginPolicy struct {
		OnlyOneTerminalAccess                                  bool `yaml:"onlyOneTerminalAccess"`
		MobileAndPCTerminalAccessButOtherTerminalKickEachOther bool `yaml:"mobileAndPCTerminalAccessButOtherTerminalKickEachOther"`
		AllTerminalAccess                                      bool `yaml:"allTerminalAccess"`
	}
	TokenPolicy struct {
		AccessSecret string `yaml:"accessSecret"`
		AccessExpire int64  `yaml:"accessExpire"`
	}
	MessageCallBack struct {
		CallbackSwitch bool   `yaml:"callbackSwitch"`
		CallbackUrl    string `yaml:"callbackUrl"`
	}
}

func init() {
	confDir:= base.ConfigDir()
	bytes, err := ioutil.ReadFile(confDir+ "config.yaml")
	//bytes, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
		return
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err)
		return
	}

}
