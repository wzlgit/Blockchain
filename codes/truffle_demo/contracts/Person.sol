pragma solidity ^0.4.23;
/*
   pragma:版本声明
   solidity：开发语言
   0.4.23：当前合约的版本 ，向上兼容，(0.4.23~0.49)均可编译
*/

// contract 关键字声明，与面向对象编程的类相似
contract Person {
    uint age; // 状态变量，默认为internal
    string internal name = "人类";  // internal表示只能在当前合约或者合约的子类中使用
    string public job = "IT";    // public声明的属性，会自动生成同名的get函数
    string private phone = "13512345678"; // private表示只能在当前合约使用
    Address workAddress;

    struct Address {  // 声明结构体
      string country;
      string province;
      string city;
    }

    // 构造函数，用constructor修饰，初始化时会自动被调用
    constructor() public {
        age = 0;
        workAddress = Address('中国','汉东省','京州');
    }

    // set方法，方法默认为public，能够被外部合约访问
    function setPhone(string p) public {
        phone = p;
    }
    // get方法
    function getPhone() constant public returns (string) {
        return phone;
    }

    // msg.sender 返回操作当前合约的账户地址
    function getCurrentAddress() view public returns (address) {
        return msg.sender;
    }

    function test() pure public returns (string) {
      return "test";
    }

    function getWorkCity() constant public returns (string) {
        return workAddress.city;
    }

    // 普通函数
    function kill() public{
        // 析构函数，调用时销毁当前合约
        selfdestruct(msg.sender);
    }
}

// is表示继承
contract Student is Person {
    // 函数默认声明为public，不写public编译会出现警告
    constructor() public {
        age = 18;
        name = "学生仔";
    }

    function sayHello() pure public returns (string) {
        // 不能直接用字符串拼接，会报错
        // return (name + age);
        return "hello";
    }
}
