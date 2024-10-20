package db

import (
	"testing"
	"time"
	"zproject/user_service/common/utils"
)

func TestUserDBImpl_GetUserByIDs(t *testing.T) {
	u := InitUserDBImpl()
	defer func(st time.Time) {
		t.Logf("cost %v\n", time.Now().Sub(st))
	}(time.Now())

	info, err := u.GetUserByIDs([]int64{
		222,
		82,
		356,
		725,
		505,
		265})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user info is %v\n", utils.JsonMarshal(info))

}
