package etcdclient

import (
	"context"
	"errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"strconv"
)

func put(key string, value string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints(),
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, key, value)
	cancel()
	return err
}

func get(key string) (string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints(),
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return "", err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil || resp == nil{
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", errors.New("key not found")
	}
	return string(resp.Kvs[0].Value), nil
}

func makeClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints(),
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}


func FixedLengthItoa(num int, length int) string {
	numstr := strconv.Itoa(num)
	return strings.Repeat("0", length - len(numstr)) + numstr
}