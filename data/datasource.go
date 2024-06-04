package data

import (
	"context"
	"fmt"
	"time"

	"aleksandersh.dev/etcd-tui/domain"
	etcd "go.etcd.io/etcd/client/v3"
)

const (
	requestTimeOut = 5 * time.Second
)

type EtcdDataSource struct {
	cli *etcd.Client
}

func NewEtcdDataSource(cli *etcd.Client) *EtcdDataSource {
	return &EtcdDataSource{cli: cli}
}

func (d *EtcdDataSource) GetEntityList(ctx context.Context) (*domain.EntityList, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeOut)
	defer cancel()

	resp, err := d.cli.Get(ctx, "", etcd.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to load keys: %w", err)
	}

	list := make([]domain.Entity, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		list = append(list, *domain.NewEntity(string(kv.Key), string(kv.Value)))
	}
	return domain.NewEntityList(list), nil
}

func (d *EtcdDataSource) SaveKeyValue(ctx context.Context, key string, value string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := d.cli.Put(ctx, key, value); err != nil {
		return fmt.Errorf("failed to save value: %w", err)
	}
	return nil
}

func (d *EtcdDataSource) DeleteKey(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := d.cli.Delete(ctx, key); err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}
	return nil
}
