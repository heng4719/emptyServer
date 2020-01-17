package model

import (
	"goRssMail-go/pkg/DB"
)

func Init() {
	//需要同步的表结构
	if err := DB.Engine.Sync2(
		//new(sys.Employee),
		//new(sys.Role),
		//new(sys.OrgType),
		//new(sys.SysMenu),
		//new(sys.RoleMenu),
		//new(sys.Permission),
		//new(sys.RolePermission),
		//new(base.Bank),
		//new(base.ServiceProvider),
		//new(base.Shop),
		
	); err != nil {
		panic(err)
	}

	// 初始化数据
	//if err := initData(); err != nil {
	//	panic(err)
	//}
}
//
//func initData() (err error) {
//	if err = initRole(); err != nil {
//		return err
//	}
//
//	if err = initAccount(); err != nil {
//		return err
//	}
//
//	return
//}
//
//func initRole() (err error) {
//	count, err := DB.Where("id = 1").Count(&sys.Role{})
//	if err != nil {
//		return fmt.Errorf("init superadmin role err: %v\n", err)
//	}
//
//	if count > 0 {
//		return nil
//	}
//
//	role := &sys.Role{Id: 1, Code: "1001", IsAdmin: 1, Name: "超级管理员", Buildin: 1}
//	_, err = DB.InsertOne(role)
//	return err
//}
//
//// 初始化超级管理员账号
//func initAccount() (err error) {
//	count, err := DB.Where("account=?", "super_admin").Count(&sys.Employee{})
//	if err != nil {
//		glog.Fatalf("init superadmin account err: %v\n", err)
//		panic(err)
//	}
//	if count > 0 {
//		return
//	}
//
//	session := DB.Engine.NewSession()
//	defer session.Close()
//
//	if err = session.Begin(); err != nil {
//		return fmt.Errorf("session begin err: %s", err)
//	}
//
//	password := gosecurity.MD5Password("111111")
//	employee := &sys.Employee{
//		Id:       1,
//		Name:     "超级管理员",
//		Account:  "super_admin",
//		Password: password,
//		Phone:    "18601694368",
//		RoleId:   1,
//		Code:     "1000",
//		RoleName: "超级管理员",
//		Type:     employee_enum.ADMIN,
//	}
//	if _, err = DB.InsertOneTx(session, employee); err != nil {
//		session.Rollback()
//		return err
//	}
//
//	err = session.Commit()
//	return err
//}
