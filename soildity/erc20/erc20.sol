// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// 我们在这里使用 openzeppelin 安全的智能合约开发框架  是经过严格审计的库

contract DollarKillerChain is ERC20, ERC20Burnable, Pausable, Ownable {
    // 发行 全名为 DollarKillerChain 简称为 DKC 的 Coin
    constructor() ERC20("DollarKillerChain", "DKC") {
        // 预先挖矿 1000 个
        // 单位10个0
        _mint(msg.sender, 1000 * 10 ** decimals());
    }

    // 暂停智能合约
    function pause() public onlyOwner {
        _pause();
    }

    // 接触暂停
    function unpause() public onlyOwner {
        _unpause();
    }

    // 再挖矿
    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount)
    internal
    whenNotPaused
    override
    {
        super._beforeTokenTransfer(from, to, amount);
    }
}