package etcdclient_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEtcdClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EtcdClient Suite")
}
