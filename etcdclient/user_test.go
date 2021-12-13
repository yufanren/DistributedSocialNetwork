package etcdclient_test

import (
	"cs9223-final-project/etcdclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	BeforeEach(func() {
		etcdclient.Register_("Alice", "123")
		etcdclient.Register_("Bob", "321")
	})

	Context("User test", func() {
		It("errors in Login", func() {
			Expect(etcdclient.Login_("Alice", "123")).To(BeNil())
			Expect(etcdclient.Login_("Bob", "123")).ToNot(BeNil())
		})
		It("errors in Follow", func() {
			defer etcdclient.Follow_("Alice", "Bob")
			Expect(etcdclient.Follow_("Alice", "Bob")).To(BeNil())
		})
		It("errors in IsFollow", func() {
			etcdclient.Follow_("Alice", "Bob")
			defer etcdclient.Follow_("Alice", "Bob")
			Expect(etcdclient.Isfollow_("Alice", "Bob")).To(Equal(true))
			Expect(etcdclient.Isfollow_("Bob", "Alice")).To(Equal(false))
		})
		It("errors in Getuserfollowed", func() {
			etcdclient.Follow_("Alice", "Bob")
			defer etcdclient.Follow_("Alice", "Bob")
			followedNumber, followedUser := etcdclient.Getuserfollowed_("Alice")
			Expect(followedNumber).To(Equal(1))
			Expect(followedUser[0]).To(Equal("Bob"))
		})
		It("errors in Getuserfollower", func() {
			etcdclient.Follow_("Alice", "Bob")
			defer etcdclient.Follow_("Alice", "Bob")
			followerNumber, followerUser := etcdclient.Getuserfollower_("Bob")
			Expect(followerNumber).To(Equal(1))
			Expect(followerUser[0]).To(Equal("Alice"))
		})
	})

})
