pragma solidity ^ 0.4 .17;

contract Adoption {
  // address类型指向的是以太坊地址，存储为20个字节的值
  // 定义了一个固定长度16为的数组，也就是16个领养宠物的人对应的以太坊地址
  address[16] public adopters;

  // 领养宠物
  function adopt(uint petId) public returns(uint) {
    // require函数用来检查petId，确保petId在0~15之间，防止数组下标越界
    require(petId >= 0 && petId <= 15);
    // msg.sender表示调用此函数或者智能合约的地址
    adopters[petId] = msg.sender;
    // 返回petId作为确认
    return petId;
  }

  // 获取领养人
  // 确保返回类型是adopters指定的类型->address[16]
  function getAdopters() public view returns(address[16]) {
    return adopters;
  }
}
