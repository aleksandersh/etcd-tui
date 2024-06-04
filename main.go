package main

import (
	"context"
	"log"
	"time"

	"github.com/aleksandersh/etcd-tui/cli"
	"github.com/aleksandersh/etcd-tui/data"
	"github.com/aleksandersh/etcd-tui/domain"
	"github.com/aleksandersh/etcd-tui/tui"
	etcd "go.etcd.io/etcd/client/v3"
)

const (
	connectionTimeOut = 5 * time.Second
)

func main() {
	args := cli.GetArgs()

	etcdCfg := etcd.Config{
		Endpoints:   args.Endpoints,
		Username:    args.Username,
		Password:    args.Password,
		DialTimeout: connectionTimeOut,
	}
	cli, err := etcd.New(etcdCfg)
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v, %v", etcdCfg.Endpoints, err)
	}
	defer cli.Close()

	log.Printf("etcd connected: %v", etcdCfg.Endpoints)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataSource := data.NewEtcdDataSource(cli)
	list, err := dataSource.GetEntityList(ctx)
	if err != nil {
		log.Fatalf("failed to load keys: %v", err)
	}
	cfg := domain.NewConfig(args.Title)
	tui.RunApp(ctx, cfg, dataSource, list)
	if err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}
