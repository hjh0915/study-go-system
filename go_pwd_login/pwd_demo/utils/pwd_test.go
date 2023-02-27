package utils

import (
	"testing"
)

func TestA(t *testing.T) {
	t.Log("TestA")
}

func TestPasswordHash(t *testing.T) {
	t.Log("--> TestHelloWorld ")

	pwd := "123456"
	hash := PasswordHash(pwd)

	t.Log("--> 输入密码:", pwd)
	t.Log("--> 生成hash:", hash)

	match := PasswordVerify(pwd, hash)
	t.Log("--> 验证结果:", match)
}

func TestPasswordVerify(t *testing.T) {
	t.Log("--> TestpwdVerify ")

	pwd := "123456789"
	hash := "$2a$04$6H4qvBoDFXjV8rLDR1wwmOrgH.q6y954TQg7fPG0G4iReT3gLXUW."

	match := PasswordVerify(pwd, hash)
	t.Log("--> TestpwdVerify 验证结果:", match)
	if match == false {
		t.Errorf("TestpwdVerify failed. Got false, expected true.")
	}
}
