package account_test

//import (
//	"testing"

//	"github.com/brendanjcarlson/visql/server/src/pkg/config"
//	"github.com/brendanjcarlson/visql/server/src/pkg/database"
//	"github.com/brendanjcarlson/visql/server/src/pkg/domains/account"
//)

//var repo *account.Repository

//func init() {
//	config.MustLoadEnv(".env")
//	client := database.MustConnect()
//	repo = account.NewRepository(client)
//}

//func TestCreate(t *testing.T) {
//	t.Skip("tests not yet implemented")
//	ne := &account.NewEntity{
//		FullName: "Test User",
//		Email:    "test@user.com",
//		Password: "password",
//	}

//	e, err := repo.Create(ne)
//	if err != nil {
//		t.Errorf("expected nil error\ngot: %v\n", err.Error())
//	}
//	if e == nil {
//		t.Errorf("expected non-empty entity\ngot: %v\n", e)
//	}
//}

//func TestFindByEmail(t *testing.T) {
//	t.Skip("tests not yet implemented")
//}

//func TestFindById(t *testing.T) {
//	t.Skip("tests not yet implemented")
//}

//func TestUpdate(t *testing.T) {
//	t.Skip("tests not yet implemented")
//}

//func TestUpdateOnLogin(t *testing.T) {
//	t.Skip("tests not yet implemented")
//}

//func TestDelete(t *testing.T) {
//	t.Skip("tests not yet implemented")
//}
