package backend_test

import (
	"final-project-0b5a2e16-ly2062-yufanren/backend"
	"final-project-0b5a2e16-ly2062-yufanren/backend/backendpb"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Db", func() {
	Context("Db test", func() {

		It("user thread-safe test", func() {
			db := backend.NewDB()
			go func() {
				db.AddUser("bot", "123")
			}()
			go func() {
				db.HasUser("bot")
			}()
			go func() {
				db.ValidUser("bot", "123")
			}()
			db.AddUser("bot", "123")
		})

		It("blog thread-safe test", func() {
			db := backend.NewDB()
			go func() {
				db.AddBlog(&backendpb.Blog{Title: "test", Content: "hi", Author: "bot"})
			}()
			go func() {
				db.GetBlog(0)
			}()
			go func() {
				db.ListBlogs(1, 1)
			}()
			go func() {
				db.ListFollowBlog("bot", 1, 1)
			}()
			db.ListBlogsFromUser("bot", 1, 1)
		})

	})
})
