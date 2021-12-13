package etcdclient_test

import (
	"cs9223-final-project/etcdclient/etcdclientpb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cs9223-final-project/etcdclient"
)

var _ = Describe("Blog", func() {
	Context("User test", func() {
		It("errors in AddBlog", func() {
			blog := &etcdclientpb.Blog{Title: "Test", Content: "Blah", Author: "Alice"}
			_, err := etcdclient.Addblog_(blog)
			Expect(err).To(BeNil())
		})
		It("errors in ListBlog", func() {
			_, blogs := etcdclient.Listblog_(1, 1)
			Expect(len(blogs)).To(Equal(1))
		})
		It("errors in GetBlog", func() {
			_, err := etcdclient.Getblog_(0)
			Expect(err).To(BeNil())
		})
		It("errors in GetUserBlog", func() {
			_, blogs := etcdclient.Getuserblog_("Alice", 1, 1)
			Expect(len(blogs)).To(Equal(1))
		})
	})
})
