package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	rbac()
}

func rbac() {
	ef, err := casbin.NewEnforcer("./conf/rbac/model.conf", "./conf/rbac/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforcer error: %v", err)
	}

	check(ef, "admin_user1", "/admin/data1/detail", "get")
	check(ef, "admin_user1", "/admin/data1/detail", "post")
	check(ef, "guest_user1", "/admin/data1/detail", "get")
	check(ef, "guest_user1", "/admin/data1/detail", "post")
	check(ef, "guest_user_no_read1", "/admin/data1/detail", "get")
	check(ef, "guest_user_no_read1", "/admin/data1/detail", "post")
}

func check(ef *casbin.Enforcer, sub, obj, act string) {
	ok, _ := ef.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("sub:%s obj:%s act:%s result:yes\n", sub, obj, act)
	} else {
		fmt.Printf("sub:%s obj:%s act:%s result:no\n", sub, obj, act)
	}
}

// acl PERM(Policy, Effect, Request, Matchers) 模型
func acl() {
	ef, err := casbin.NewEnforcer("./conf/acl/model.conf", "./conf/acl/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforcer error: %v", err)
	}

	check(ef, "admin_user1", "admin_data1", "read")
	check(ef, "admin_user1", "admin_data1", "write")
	check(ef, "admin_user_no_write1", "admin_data1", "read")
	check(ef, "admin_user_no_write1", "admin_data1", "write")
	check(ef, "guest_user1", "admin_data1", "read")
	check(ef, "guest_user1", "admin_data1", "write")
	check(ef, "root", "admin_data1", "read")
	check(ef, "root", "admin_data1", "write")
}
