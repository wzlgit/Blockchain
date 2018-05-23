pragma solidity ^0.4.18;

contract SimpleStorage {
  // 用于存储图片的哈希值
  string storedData;

  function set(string x) public {
    storedData = x;
  }

  function get() public view returns (string) {
    return storedData;
  }
}
