// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

// ERC721 基础模块, ERC721URIStorage 获取URL地址, ERC721Burnable 可燃烧的代币 (未来可以用来做合成 2 个或者多个NFT合成一个新的)
contract DollerKillerToken is ERC721, ERC721URIStorage, ERC721Burnable, Ownable {
    using Counters for Counters.Counter;  // 计数器 NFT系列号使用整数表示

    Counters.Counter private _tokenIdCounter;

    // DollerKillerToken NFT全称, DKT NFT简称
    constructor() ERC721("DollerKillerToken", "DKT") {}

    // 具体NFT资源地址前缀   可能是IPFS 可能是S3
    function _baseURI() internal pure override returns (string memory) {
        return "https://i.imgur.com/";
    }

    // 可以挖矿  
    function safeMint(address to, string memory uri) public onlyOwner {
        // 当前NFT序列 递增
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri); // 设置Mappding tokenID 与URL关联
    }

    // The following functions are overrides required by Solidity.

    // 可燃烧
    function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
    }

    // 获取 nft具体 url
    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}   