package svc

import (
	"cds/galaxy/internal/config"
	"cds/galaxy/internal/model"
	"cds/tools/canalx"
	"cds/tools/debeziumx"
	"log"

	"go.etcd.io/etcd/clientv3"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       *model.UserModel
	GroupModel      *model.GroupModel
	PermissionModel *model.PermissionModel
	DmModel         *model.DmModel
	RtuModel        *model.RtuModel
	ConnectorModel  *model.ConnectorModel
	EtcdClient      *clientv3.Client
	DebeziumClient  *debeziumx.Debeziumx
	CanalClient     *canalx.Canalx
}

func NewServiceContext(config config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config:          config,
		UserModel:       model.NewUserModel(config.Mysql),
		GroupModel:      model.NewGroupModel(config.Mysql),
		PermissionModel: model.NewPermissionModel(config.Mysql),
		DmModel:         model.NewDmModel(config.Mysql),
		RtuModel:        model.NewRtuModel(config.Mysql),
		ConnectorModel:  model.NewConnectorModel(config.Mysql),
	}
	var e error
	ctx.EtcdClient, e = clientv3.New(clientv3.Config{
		Endpoints: config.EtcdConfig,
	})
	if e != nil {
		log.Fatal(e)
	}
	ctx.DebeziumClient, e = debeziumx.NewDebeziumx(config.Debezium)
	if e != nil {
		log.Fatal(e)
	}
	ctx.CanalClient = canalx.NewCanalx(config.CanalConfig.IP, config.CanalConfig.Port, config.CanalConfig.UserName, config.CanalConfig.Password)
	return ctx
}
