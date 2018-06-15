package main

import (
  "encoding/json"
  "fmt"
  "strconv"
  "errors"

  "github.com/hyperledger/fabric/core/chaincode/shim"
  pb "github.com/hyperledger/fabric/protos/peer"
)

// 定义智能合约结构体
type SmartContract struct {
}

/*
 * 定义"部门"结构体
 * 部门字段包括部门ID和部门名称
 */
type Department struct {
  DepartmentID int `json:department_id` // 部门ID
  DepartmentName string `json:department_name`  // 部门名称
}

/*
 * 定义"员工"结构结构体
 * 员工字段包括员工ID、员工姓名、员工所属部门ID、工作岗位
 */
 type Employee struct {
   EmployeeID int `json:employee_id`  // 员工ID
   EmployeeName string `json:employee_name` // 员工姓名
   DepartmentID int `json:department_id` // 部门ID
   Jobs string `json:jobs`  // 工作岗位
 }

 // 在链码初始化过程中调用Init来初始化任何数据
 func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
   fmt.Println("my_chaincode01 Init")
   return shim.Success(nil)
 }

 // 在链码每个事务中，Invoke会被调用。
 func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
   fmt.Println("my_chaincode01 Invoke")

   function, args := stub.GetFunctionAndParameters()
   if function == "initDepartment" {
     return t.initDepartment(stub, args)
   } else if function == "addEmployee" {
     return t.addEmployee(stub, args)
   } else if function == "deleteEmployee" {
     return t.deleteEmployee(stub, args)
   } else if function == "searchEmployeeInfoByID" {
     return t.searchEmployeeInfoByID(stub, args)
   } else if function == "updateEmployeeInfo" {
     return t.updateEmployeeInfo(stub, args)
   }

   return shim.Error("Invalid Smart Contract function name.")
 }

  // 初始化部门
  func (t *SmartContract) initDepartment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 2 {
      return shim.Error("Incorrect number of arguments. Expecting 2 like (departmentId, departmentName)")
    }

    departmentIdAsString := args[0]
    // 转换为int类型
    departmentIdAsInt, err := strconv.Atoi(args[0])
    if err != nil {
      return shim.Error("first argument must be a numeric string")
    }
    departmentName := args[1]
    // 创建"部门"结构体
    department := &Department{departmentIdAsInt, departmentName}

    // 创建联合主键，用多个列组合作为联合主键
    // Fabric是用U+0000来把各个联合主键的字段拼接起来，因为这个字符太特殊，所以很适合，
    departmentIdKey, err := stub.CreateCompositeKey("Department", []string{"department", departmentIdAsString})
    if err != nil {
		   return shim.Error(err.Error())
	  }

    // 结构体转换为json字符串
    departmentAsBytes, err := json.Marshal(department)
    if err != nil {
      return shim.Error(err.Error())
    }

    // 新增一条"部门"数据
    err = stub.PutState(departmentIdKey, departmentAsBytes)
    if err != nil {
		    return shim.Error(err.Error())
	  }
    return shim.Success(departmentAsBytes)
 }

 // 新增员工
 func (t *SmartContract) addEmployee(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   // 将字符串数组类型数据转换为"员工"结构体
   employee, err := translateEmployeeFromArgs(args)
	 if err != nil {
	  	return shim.Error(err.Error())
	 }
   fmt.Println("employee:", employee)
	 employeeIdAsString := strconv.Itoa(employee.EmployeeID)
   // 检查添加的部门ID是否已经存在，返回所有的部门ID
   departmentIds := queryAllDepartmentIDs(stub)
   fmt.Println("departmentIds:", departmentIds)

   // 是否已经存在该员工
   isExist := false
   if len(departmentIds) > 0 {
     for _, departmentId := range departmentIds {
       // 转换为int类型
       departmentIdAsInt, err := strconv.Atoi(departmentId)
       if err != nil {
    		 return shim.Error("Department Id argument must be a numeric string")
    	 }

       if departmentIdAsInt == employee.DepartmentID {
         isExist = true
         break
       }
     }
   }

   if isExist {
     // 读取账本中的数据
     employeeAsBytes, err := stub.GetState(employeeIdAsString)
     if err != nil {
       return shim.Error(err.Error())
     } else if employeeAsBytes != nil {
       fmt.Println("This employee already exists: " + employeeIdAsString)
       return shim.Error("This employee already exists: " + employeeIdAsString)
     }

     // 结构体转换为json字符串
     employeeAsJsonBytes, err := json.Marshal(employee)
     if err != nil {
       return shim.Error(err.Error())
     }
     // 保存到账本中
     err = stub.PutState(employeeIdAsString, employeeAsJsonBytes)
     if err != nil {
       return shim.Error(err.Error())
     }
     return shim.Success(employeeAsJsonBytes)
   } else {
     fmt.Println("department:" + string(employee.DepartmentID) + " does not exist")
		 return shim.Error("department:" + string(employee.DepartmentID) +  " does not exist")
   }
 }

 // 删除员工
 func (t *SmartContract) deleteEmployee(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 like (employee_id)")
	 }
   employeeIdAsString := args[0]
   employeeAsBytes, err := stub.GetState(employeeIdAsString)
   if err != nil {
     return shim.Error("Failed to get employee info:" + err.Error())
   } else if employeeAsBytes == nil {
     return shim.Error("Employee does not exist")
   }

   err = stub.DelState(employeeIdAsString)
   if err != nil {
     return shim.Error("Failed to delete employee:" + employeeIdAsString + err.Error())
   }
   return shim.Success(nil)
 }

 // 根据员工ID查询员工信息
 func (t *SmartContract) searchEmployeeInfoByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   if len(args) < 1 {
     return shim.Error("Incorrect number of arguments. Expecting 1 like (employee_id)")
   }
   employeeIdAsString := args[0]
	 employeeAsBytes, err := stub.GetState(employeeIdAsString)
	 if err != nil {
		 return shim.Error("Failed to get employee info:" + err.Error())
	 } else if employeeAsBytes == nil {
		 return shim.Error("Employee does not exist")
   }

   fmt.Printf("Search Response:%s\n", string(employeeAsBytes))
 	 return shim.Success(employeeAsBytes)
 }

 // 更新员工信息
 func (t *SmartContract) updateEmployeeInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   // 将字符串数组类型数据转换为"员工"结构体
   employee, err := translateEmployeeFromArgs(args)
   if err != nil {
     return shim.Error(err.Error())
   }

   employeeIdAsString := strconv.Itoa(employee.EmployeeID)
   // 检查添加的部门ID是否存在
   departmentIds := queryAllDepartmentIDs(stub)

   isExist := false
   if len(departmentIds) > 0 {
     for _, departmentId := range departmentIds {
       // 转换为int类型
       departmentIdAsInt, err := strconv.Atoi(departmentId)
       if err != nil {
    		 return shim.Error("Department Id argument must be a numeric string")
    	 }

       if departmentIdAsInt == employee.DepartmentID {
         isExist = true
         break
       }
     }
   }

   if isExist {
     /*
     * State DB是一个Key-Value数据库，如果我们指定的Key在数据库中已经存在，那么就是修改操作。
     * 如果Key不存在，那么就是插入操作。
     */
     employeeAsJsonBytes, err := json.Marshal(employee)
     if err != nil {
       return shim.Error(err.Error())
     }
     // 保存到账本中
     err = stub.PutState(employeeIdAsString, employeeAsJsonBytes)
     if err != nil {
       return shim.Error(err.Error())
     }
     return shim.Success(employeeAsJsonBytes)
   } else {
     fmt.Println("department:" + string(employee.DepartmentID) + " does not exist")
		 return shim.Error("department:" + string(employee.DepartmentID) + " does not exist")
   }
 }

 // 获取所有部门ID
 func queryAllDepartmentIDs(stub shim.ChaincodeStubInterface) []string {
   // 部分复合键的查询GetStateByPartialCompositeKey，一种对Key进行前缀匹配的查询
	compositeKeysIterator, err := stub.GetStateByPartialCompositeKey("Department", []string{"department"})
	if err != nil {
		return nil
	}

	defer compositeKeysIterator.Close()

	departmentIds := make([]string, 0)

	for i := 0; compositeKeysIterator.HasNext(); i++ {
		responseRange, err := compositeKeysIterator.Next()
		if err != nil {
			return nil
		}
    // 拆分复合键SplitCompositeKey，
		_, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil
		}
		departmentId := compositeKeyParts[1]
		departmentIds = append(departmentIds, departmentId)
	}

	return departmentIds
}

 // 将字符串数组类型数据转换为"员工"结构体
 func translateEmployeeFromArgs(args []string) (*Employee, error) {
   if len(args) != 4 {
     return nil, errors.New("Incorrect number of arguments. Expecting 4 like (employeeId, employeeName, departmentId, jobs)")
   }

   // 转换为int类型
   employeeId, err := strconv.Atoi(args[0])
   if err != nil {
		 return nil, errors.New("first argument must be a numeric string")
	 }
   employeeName := args[1]
   departmentId, err := strconv.Atoi(args[2])
   if err != nil {
		 return nil, errors.New("third argument must be a numeric string")
	 }
   jobs := args[3]

   employee := &Employee{employeeId, employeeName, departmentId, jobs}
   return employee, nil
 }

 // Go语言的入口是main函数
 func main() {
   err := shim.Start(new(SmartContract))
   if err != nil {
     fmt.Printf("Error creating new Smart Contract: %", err)
   }
 }
