package main

import (
  "encoding/json"
  "fmt"
  "strconv"
  "bytes"
  "time"

  "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// 定义"弹珠"链码的结构体
type MarblesChaincode struct {
}

/*
 * 定义"弹珠"结构体
 * 字段包括ID、颜色、尺寸和拥有者等
 */
type Marble struct {
  ObjectType string `json:"objectType"` // CouchDB的字段
  ID string `json:"id"` // ID（唯一字符串，将用作键）
  Color string `json:"color"` // 颜色（字符串，CSS颜色名称）
  Size int `json:"size"`  //  尺寸（整型，以毫米为单位）
  Owner string `json:"owner"` //  拥有者（字符串）
}

// 在链码初始化过程中调用Init来初始化任何数据
func (t *MarblesChaincode) Init (stub shim.ChaincodeStubInterface) pb.Response {
  fmt.Println("Marbles Init Success")
  return shim.Success(nil)
}

// 在链码每个交易中，Invoke函数会被调用。
func (t *MarblesChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
  fmt.Println("Marbles Invoke")
  function, args := stub.GetFunctionAndParameters()
  if function == "initMarble" {
    return t.initMarble(stub, args)
  } else if function == "getMarbleInfoByID" {
    return t.getMarbleInfoByID(stub, args)
  } else if function == "deleteMarbleByID" {
    return t.deleteMarbleByID(stub, args)
  } else if function == "changeMarbleOwner" {
    return t.changeMarbleOwner(stub, args)
  } else if function == "getMarbleByRange" {
    return t.getMarbleByRange(stub, args)
  } else if function == "getMarbleByOwner" {
    return t.getMarbleByOwner(stub, args)
  } else if function == "getMarbleHistoryForKey" {
    return t.getMarbleHistoryForKey(stub, args)
  }
  return shim.Error("没找到对应方法~")
}

// 初始化弹珠
func (t *MarblesChaincode) initMarble (stub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 4 {
    return shim.Error("Incorrect number of arguments. Expecting 4")
  }

  marbleObjectType := "marble"
  marbleId := args[0]
  // 查询账本中是否已经存在该弹珠ID的信息
  marbleIdAsBytes, err := stub.GetState(marbleId)
  if err != nil {
    return shim.Error(err.Error())
  }

  if marbleIdAsBytes != nil {
    return shim.Error("该Marble已存在~")
  }

  marbleColor := args[1]
  // 字符串转换为整型
  marbleSize, err := strconv.Atoi(args[2])
  if err != nil {
    return shim.Error("size字段为整型~")
  }
  marbleOwner := args[3]
  // 创建Marble对象
  marble := &Marble{marbleObjectType, marbleId, marbleColor, marbleSize, marbleOwner}
  // 转换为Json格式的内容
  marbleJsonAsBytes, err := json.Marshal(marble)
  // 写入到账本中
  err = stub.PutState(marbleId, marbleJsonAsBytes)
  if err != nil {
    return shim.Error(err.Error())
  }

  return shim.Success(nil)
}

// 通过弹珠ID查询对应的弹珠信息
func (t *MarblesChaincode) getMarbleInfoByID (stub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
  }

  marbleId := args[0]
  // 查询账本中是否已经存在该弹珠ID的信息
  marbleIdAsBytes, err := stub.GetState(marbleId)

  if err != nil {
    return shim.Error(err.Error())
  } else if marbleIdAsBytes == nil {
    return shim.Error("该Marble不存在~")
  }

  return shim.Success(marbleIdAsBytes)
}

// 通过弹珠ID删除对应的弹珠
func (t *MarblesChaincode) deleteMarbleByID (stub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
  }

  marbleId := args[0]
  // 查询账本中是否已经存在该弹珠ID的信息
  marbleIdAsBytes, err := stub.GetState(marbleId)

  if err != nil {
    return shim.Error(err.Error())
  } else if marbleIdAsBytes == nil {
    return shim.Error("该Marble不存在~")
  }

  // 删除账本中的记录
  err = stub.DelState(marbleId)
  if err != nil {
    return shim.Error(err.Error())
  }

  return shim.Success(nil)
}

// 改变弹珠的拥有者
func (t *MarblesChaincode) changeMarbleOwner (stub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

  marbleId := args[0]
  newOwner := args[1]

  // 查询账本中是否已经存在该弹珠ID的信息
  marbleIdAsBytes, err := stub.GetState(marbleId)

  if err != nil {
    return shim.Error(err.Error())
  } else if marbleIdAsBytes == nil {
    return shim.Error("该Marble不存在~")
  }

  marble := Marble{}
  err = json.Unmarshal(marbleIdAsBytes, &marble)
  if err != nil {
    return shim.Error(err.Error())
  }
  marble.Owner = newOwner
  marbleJsonAsBytes, err := json.Marshal(marble)
  if err != nil {
    return shim.Error(err.Error())
  }

  // 写入到账本中
  err = stub.PutState(marbleId, marbleJsonAsBytes)
  if err != nil {
    return shim.Error(err.Error())
  }

  return shim.Success(nil)
}

// 查询指定ID范围的弹珠信息
func (t *MarblesChaincode) getMarbleByRange (stub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

  startKey := args[0]
  endKey := args[1]

  // 查询指定范围内的键值
  resultsIterator, err := stub.GetStateByRange(startKey, endKey)
  if err != nil {
    return shim.Error(err.Error())
  }
  defer resultsIterator.Close()

  var buffer bytes.Buffer
  buffer.WriteString("[")

  bArrayMemberAlreadyWritten := false
  // 遍历弹珠信息，转换为Json格式的字符串数组并返回
  for resultsIterator.HasNext() {
    queryResult, err := resultsIterator.Next()
    if err != nil {
			return shim.Error(err.Error())
		}

    if bArrayMemberAlreadyWritten == true {
      buffer.WriteString(",")
    }

    buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResult.Key)
		buffer.WriteString("\"")

    buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResult.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
  }
  buffer.WriteString("]")
  fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())
  return shim.Success(buffer.Bytes())
}

// 查询某个拥有者的所有弹珠信息
func (t *MarblesChaincode) getMarbleByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  marbleOwner := args[0]
  queryStr := fmt.Sprintf("{\"selector\":{\"owner\":\"%s\"}}",marbleOwner)

  resultsIterator, err := stub.GetQueryResult(queryStr)
  if err != nil {
			return shim.Error(err.Error())
	}

  defer resultsIterator.Close()

  var buffer bytes.Buffer
  buffer.WriteString("[")

  bArrayMemberAlreadyWritten := false
  for resultsIterator.HasNext() {
    queryResult, err := resultsIterator.Next()
    if err != nil {
			return shim.Error(err.Error())
		}

    if bArrayMemberAlreadyWritten == true {
      buffer.WriteString(",")
    }

    buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResult.Key)
		buffer.WriteString("\"")

    buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResult.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
  }
  buffer.WriteString("]")

  fmt.Printf("- getMarbleByOwner queryResult:\n%s\n", buffer.String())
  return shim.Success(buffer.Bytes())
}

// 通过弹珠ID查询所有的交易历史信息
func (t *MarblesChaincode) getMarbleHistoryForKey (stub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

  marbleId := args[0]

  // 返回某个键的所有历史值
  resultsIterator, err := stub.GetHistoryForKey(marbleId)
  if err != nil {
    return shim.Error(err.Error())
  }

  defer resultsIterator.Close()

  var buffer bytes.Buffer
  buffer.WriteString("[")

  bArrayMemberAlreadyWritten := false
  for resultsIterator.HasNext() {
    queryResult, err := resultsIterator.Next()
    if err != nil {
			return shim.Error(err.Error())
		}

    if bArrayMemberAlreadyWritten == true {
      buffer.WriteString(",")
    }

    buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResult.TxId)
		buffer.WriteString("\"")

    buffer.WriteString(", \"Timestamp\":")
    buffer.WriteString("\"")
		buffer.WriteString(time.Unix(queryResult.Timestamp.Seconds, int64(queryResult.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

    buffer.WriteString("{\"Value\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(queryResult.Value))
		buffer.WriteString("\"")

    buffer.WriteString("{\"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(queryResult.IsDelete))
		buffer.WriteString("\"")

		bArrayMemberAlreadyWritten = true
  }
  buffer.WriteString("]")
  fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())
  return shim.Success(buffer.Bytes())
}

// Go语言的入口是main函数
func main() {
  err := shim.Start(new(MarblesChaincode))
  if err != nil {
    fmt.Printf("Error starting Marbles Chaincode - %s", err)
  }
}
