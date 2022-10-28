// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;  // 声明当前合约版本

contract Bank { // ERC20 就是在链上再记账, 通过mapping 实现主要内容
    mapping(address => uint256) bill; // 账单
    address owner; // 记录合约创建者地址

    constructor() {
        owner = msg.sender;
    }

    // 存钱
    function saveMoney() public payable {  // payable 表示可以接受eth
        // msg.sender 获取调用当前合约的地址
        // msg.value 存钱的金额
        bill[msg.sender] = bill[msg.sender] + msg.value;
    }

    // 查询余额
    function checkBalances() public view returns (uint256) {
        // public 公开的
        // view 只读 (不消耗gas)
        return bill[msg.sender];
    }

    // 获取银行总存款金额
    function bankBalances() public view returns (uint256) {
        // 获取当前智能合约的余额
        return address(this).balance;
    }

    // AMT 取钱
    function withdrawMoney(uint256 amount) public returns(bool) {
        // require 如果条件为真就 继续执行代码 反之抛出异常
        require(bill[msg.sender] >= amount, "Insufficient balance");

        // payable 调用智能合约转账到 调用者
        (bool isSend,) = payable(msg.sender).call{value: amount, gas: 5000}("");
        bill[msg.sender] =  bill[msg.sender] - amount;
        return isSend;
    }

    // 内部转账
    function internalTransfer(address _to, uint256 amount) public {
        require(bill[msg.sender] >= amount,"Insufficient balance");
        bill[msg.sender] =  bill[msg.sender] - amount;
        bill[_to] =  bill[_to] + amount;
    }

    // 通过智能合约进行中转 转账
    function transfer(address payable _to) public payable returns(bool) {
        (bool isSend,) = _to.call{value: msg.value, gas: 5000}("");
        return isSend;
    }
}