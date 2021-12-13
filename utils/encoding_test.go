package utils_test

import (
	"cs9223-final-project/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Student struct {
	Name string
	Courses []string
}

var _ = Describe("Encoding", func() {

	var (
		student *Student
		newStudent *Student
	)

	BeforeEach(func() {
		student = &Student{
			Name: "Alice",
			Courses: []string{"CS6233", "CS6913", "CS9223"},
		}
		newStudent = new (Student)
		encoded, _ := utils.Encode(student)
		utils.Decode(encoded, newStudent)
	})

	Context("Encode then decode", func() {
		It("should be the same", func() {
			Expect(newStudent.Name).To(Equal(student.Name))
			Expect(newStudent.Courses).To(Equal(student.Courses))
		})
	})

})
