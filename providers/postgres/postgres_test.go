package userAccount
//
// import (
// 	"context"
// 	"database/sql"
// 	"testing"
//
// 	// pr "github.com/idirall22/user/providers/postgres"
// )
//
// var id = int64(1)
// var username = "Jhon"
// var email = "Jhon@email.com"
// var userPassword = "password"
//
// // Test new
// func testNew(t *testing.T) {
// 	// fmt.Println(testService.provider.(*))
//
// 	user, err := testService.provider.New(context.Background(),
// 		username, email, "", "", userPassword, "")
//
// 	if err != nil {
// 		t.Error("Error create new user, ", err)
// 		return
//
// 	}
// 	if user.Username != username {
// 		t.Error("Error username does not match")
// 	}
// }
//
// // Test get
// func testGet(t *testing.T) {
//
// 	// Test GEt by id
// 	user, err := testService.provider.Get(context.Background(), id, "", "")
//
// 	if err != nil {
// 		t.Error("Error testing by id on Get :", err)
// 		return
// 	}
// 	if user.ID != id {
// 		t.Errorf("Error user returned id does not match %d", id)
// 	}
//
// 	// Test GEt by username
//
// 	user, err = testService.provider.Get(context.Background(), 0, username, "")
//
// 	if err != nil {
// 		t.Error("Error Get :", err)
// 		return
// 	}
// 	if user.Username != username {
// 		t.Errorf("Error user returned username does not match %s", username)
// 	}
//
// 	// Test GEt by email
//
// 	user, err = testService.provider.Get(context.Background(), 0, "", email)
//
// 	if err != nil {
// 		t.Error("Error Get :", err)
// 		return
// 	}
// 	if user.Email != email {
// 		t.Errorf("Error user returned email does not match %s", email)
// 	}
//
// 	// Test GEt by no valid params id:0, username: "", email:""
//
// 	user, err = testService.provider.Get(context.Background(), 0, "", "")
//
// 	if err == nil {
// 		t.Error("Error Get :", err)
// 		return
// 	}
//
// 	if user != nil {
// 		t.Error("Error user returned should be nil")
// 	}
// }
//
// // Test update
// func testUpdate(t *testing.T) {
// 	err := testService.provider.Update(context.Background(), 1, "jhonny", "jons", "")
// 	if err != nil {
// 		t.Error("Error Update: ", err)
// 	}
//
// 	err = testService.provider.Update(context.Background(), 0, "jhonny", "jons", "")
// 	if err != sql.ErrNoRows {
// 		t.Error("Error Update: ", err)
// 	}
// }
//
// // Test delete
// func testDelete(t *testing.T) {
//
// 	//test when id exists
// 	err := testService.provider.Delete(context.Background(), id)
// 	if err != nil {
// 		t.Error("Error to delete a user: ", err)
// 		return
// 	}
//
// 	// test when id not exists
// 	err = testService.provider.Delete(context.Background(), 0)
// 	if err != sql.ErrNoRows {
// 		t.Error("Error should be row not exists but got: ", err)
// 		return
// 	}
// }
