package etcds

import (
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var EtcdClient *clientv3.Client

func Init() (err error) {
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("etcd.endpoints"),
		DialTimeout: time.Duration(viper.GetInt("etcd.timeout")) * time.Second,
	})
	if err != nil {
		return err
	}
	return
}

func Close() (err error) {
	if EtcdClient != nil {
		err = EtcdClient.Close()
		return err
	}
	return
}
