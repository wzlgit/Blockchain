package main

/*
* 单元测试命令：
* go test -v my_chaincode01_test.go my_chaincode01.go
*/

import (
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func initDepartment(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("initDepartment"), []byte(args[0]), []byte(args[1])})

	if res.Status != shim.OK {
		fmt.Println("InitDepartment failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func addEmployee(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addEmployee"), []byte(args[0]), []byte(args[1]), []byte(args[2]), []byte(args[3])})

	if res.Status != shim.OK {
		fmt.Println("AddEmployee failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func deleteEmployee(t *testing.T, stub *shim.MockStub, employeeId string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("deleteEmployee"), []byte(employeeId)})

	if res.Status != shim.OK {
		fmt.Println("DeleteEmployee :", employeeId, ", failed :", string(res.Message))
		t.FailNow()
	}
}

func searchEmployeeInfoByID(t *testing.T, stub *shim.MockStub, employeeId string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("searchEmployeeInfoByID"), []byte(employeeId)})
	if res.Status != shim.OK {
		fmt.Println("searchEmployeeInfoByID :", employeeId, ", failed :", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("SearchEmployeeInfoByID :" , employeeId, " failed to get value")
		t.FailNow()
	}
}

func updateEmployeeInfo(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("updateEmployeeInfo"), []byte(args[0]), []byte(args[1]), []byte(args[2]), []byte(args[3])})

	if res.Status != shim.OK {
		fmt.Println("UpdateEmployeeInfo failed:", args[0], string(res.Message))
		t.FailNow()
	}
}


// 测试"初始化部门"
func TestInitDepartment(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	initDepartment(t, stub, []string{"1", "department_software"})
	initDepartment(t, stub, []string{"2", "department_test"})
}

// 测试"新增员工"，部门ID不存在时创建会失败
func TestAddEmployee(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	initDepartment(t, stub, []string{"1", "department_software"})
	addEmployee(t, stub, []string{"1", "Wenzil", "1", "Software Engineer"})
	// ID为"2"的部门没有创建，返回错误，所以先注释掉这一行
	// addEmployee(t, stub, []string{"2", "Test", "2", "Test Engineer"})
}

// 测试"查询员工信息"
func TestSearchEmployeeInfoByID(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	initDepartment(t, stub, []string{"2", "department_test"})
	addEmployee(t, stub, []string{"2", "Test", "2", "Test Engineer"})
	searchEmployeeInfoByID(t, stub, "2")
}

// 测试"删除员工"
func TestDeleteEmployee(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	initDepartment(t, stub, []string{"3", "department_ui"})
	addEmployee(t, stub, []string{"3", "Han Meimei", "3", "UI Designer"})
	deleteEmployee(t, stub, "3")
}

// 测试更新"员工信息"
func TestUpdateEmployeeInfo(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)
	initDepartment(t, stub, []string{"4", "department_blockchain"})
	addEmployee(t, stub, []string{"7", "Li Lei", "4", "Blockchain Designer"})
	updateEmployeeInfo(t, stub, []string{"7", "Li Lei", "4", "Blockchain Senior Designer"})
	searchEmployeeInfoByID(t, stub, "7")
}
