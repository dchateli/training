package inmemory_test

import (
	"github.com/dchateli/training/davidDb"
	"github.com/dchateli/training/davidDb/inmemory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inmemory", func() {

	var (
		testDb *inmemory.InMemoryDb
	)
	BeforeEach(func(){
		testDb = &inmemory.InMemoryDb{}
	})

	Describe("AddUser", func() {

		var (
			returnedUser db.User
			returnedError error
			userToCreate db.User
		)

		JustBeforeEach(func(){
			returnedUser, returnedError = testDb.AddUser(userToCreate)
		})

		Context("default", func() {
			BeforeEach(func(){
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}
			})
			It("should add a new user", func() {
				Expect(returnedError).To(BeNil())
				Expect(returnedUser.Name).To(Equal("name"))
				Expect(returnedUser.Description).To(Equal("description"))
				Expect(returnedUser.Id).NotTo(BeZero())
			})

		})

		Context("when the id is specified", func() {
			BeforeEach(func(){
				userToCreate = db.User{
					Id: "my super id",
					Name: "name",
					Description: "description",
				}
			})
			It("should add a new user", func() {
				Expect(returnedError).To(BeNil())
				Expect(returnedUser.Name).To(Equal("name"))
				Expect(returnedUser.Description).To(Equal("description"))
				Expect(returnedUser.Id).NotTo(BeZero())
				Expect(returnedUser.Id).NotTo(Equal("my super id"))
			})

		})

	})

	Describe("DeleteUser", func() {

		var (
			returnedError error
			userToCreate db.User
		)

		JustBeforeEach(func(){

			testDb.UserDB = append(testDb.UserDB, userToCreate)
		})

		Context("default", func() {


			BeforeEach(func(){
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}
			})

			JustBeforeEach(func(){
				 returnedError = testDb.DeleteUser(userToCreate.Id)
			})

			It("should delete user", func() {

				Expect(returnedError).To(BeNil())

				Expect(testDb.UserDB).To(HaveLen(0))

			})

		})

		Context("Delete nonexistent user",func(){


			BeforeEach(func(){
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}
			})

			JustBeforeEach(func(){

				returnedError = testDb.DeleteUser("5")
			})

			It("should return an error ", func() {

				Expect(testDb.UserDB).To(HaveLen(1))
				Expect(returnedError.Error()).To(Equal("User not found"))

			})
		})

	})

	Describe("ListUser", func() {

		var (
			returnedError error
			userToCreate db.User
			userToCreate2 db.User
			returnUserList []db.User
		)

		JustBeforeEach(func(){

			testDb.UserDB = append(testDb.UserDB, userToCreate,userToCreate2)
		})

		Context("default", func() {

			BeforeEach(func(){
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}

				userToCreate2 = db.User{
					Name: "name2",
					Description: "description2",
				}
			})

			JustBeforeEach(func(){

				returnUserList,returnedError = testDb.ListUser()
			})


			It("should list all user", func() {

			Expect(returnedError).To(BeNil())
			Expect(testDb.UserDB).To(HaveLen(2))
			Expect(returnUserList[0].Name).To(Equal("name"))

			})
		})
	})

	Describe("GetUser", func() {
		var (
			returnedError error
			userToCreate db.User
			returnedUser db.User
		)

		JustBeforeEach(func(){

			testDb.UserDB = append(testDb.UserDB, userToCreate)
		})

		Context("default", func() {
			BeforeEach(func() {
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}
			})

			JustBeforeEach(func(){
				returnedUser,returnedError = testDb.GetUser(userToCreate.Id)
			})

			It("should return a user", func() {

				Expect(returnedError).To(BeNil())
				Expect(returnedUser.Name).To(Equal("name"))
				Expect(returnedUser.Description).To(Equal("description"))
				Expect(returnedUser.Id).To(Equal(userToCreate.Id))

			})


		})

		Context("Nonexistent user selected ",func(){

			BeforeEach(func(){
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}
			})

			JustBeforeEach(func(){
				returnedUser,returnedError = testDb.GetUser("toto")
			})

			It("should return an error", func() {

				Expect(returnedError.Error()).To(Equal("User not found (getUser Request)"))

			})
		})
	})

	Describe("UpdateUser",func(){

		var (
			returnedError 	error
			userToCreate 	db.User
			returnedUser 	db.User
			updatedUser 	db.User
		)

		JustBeforeEach(func(){

			testDb.UserDB = append(testDb.UserDB, userToCreate)
		})

		Context("default",func(){

			BeforeEach(func(){
				userToCreate = db.User{
					Id:"5",
					Name: "name",
					Description: "description",
				}

				updatedUser = db.User{

					Name: "nameUpdate",
					Description: "descriptionUpdate",
				}

			})

			JustBeforeEach(func(){
				returnedUser,returnedError = testDb.UpdateUser(userToCreate.Id,updatedUser)
			})

			It("should Update user ", func() {

				Expect(testDb.UserDB).To(HaveLen(1))
				Expect(returnedError).To(BeNil())
				Expect(returnedUser.Name).To(Equal("nameUpdate"))
				Expect(returnedUser.Description).To(Equal("descriptionUpdate"))
				Expect(returnedUser.Id).NotTo(BeZero())

			})
		})

		Context("Update nonexistent user", func() {

			BeforeEach(func(){
				userToCreate = db.User{
					Name: "name",
					Description: "description",
				}

				updatedUser = db.User{

					Name: "nameUpdate",
					Description: "descriptionUpdate",
				}

			})

			JustBeforeEach(func(){
				returnedUser,returnedError = testDb.UpdateUser("wrongId",updatedUser)
			})

			It("should return error ", func() {

				Expect(testDb.UserDB).To(HaveLen(1))
				Expect(returnedError.Error()).To(Equal("User not found (UpdateUser request)"))

			})
		})

	})


})
