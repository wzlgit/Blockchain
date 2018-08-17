pragma solidity ^0.4.24;

import 'zeppelin-solidity/contracts/token/ERC20/StandardToken.sol';

contract TutorialToken is StandardToken {
  string public name = 'TutorialToken'; // 代币名称
  string public symbol = 'TT';  // 代币简称
  uint public decimals = 2;     // 代币的最小位数
  uint public INITIAL_SUPPLY = 12000; // 发行代币的数量

  constructor() public {
    totalSupply_ = INITIAL_SUPPLY;
    balances[msg.sender] = INITIAL_SUPPLY;
  }
}
