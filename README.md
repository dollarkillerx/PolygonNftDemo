# PolygonNftDemo

Polygon NFT Demo 

### 开发环境

- Ganache EVM虚拟机
- Remix IDE
- Blockscout  本地测试链浏览器 

### 链

- Polygon (Eth Layer2链)

### ERC20とは　 什么是ERC20

我们先从传统的银行  到第一代区块链(BTC)   

假设全世界只有一家银行

小林去吃烧烤 使用银行卡支付 给店家  花费50   

| 操作  | 账户       | 余额   |
| --- | -------- | ---- |
| 查询  | 小林       | 100  |
| 查询  | 店家       | 1000 |
| 转账  | 小林 -> 店家 | 50   |
| 查询  | 小林       | 50   |
| 查询  | 店家       | 1050 |

付款的本质 就是 银行内部账本的修改，  数字有1000 修改成 1050 账户就多出50



其本质就是账本



假设我来发行银行卡， 该银行卡只能在 在本银行开户的店铺使用 

我自己给自己开了一张卡，余额 修改为 10000 是不是我可以放肆消费了

但可惜的是，没有商铺来我这里开户，因为我没有信用 我可以肆意的修改余额

如果我是马云的儿子 以阿里巴巴的信用为担保 您愿意来我这里开户吗？



第一代区块链(BTC) 本质上就是一个 账本， 但是这个账本 和 我刚刚所谈到的又有许多不同点

1. 账本透明  所有操作和修改 都会公开  （我给我自己无中生有的钱 马上就会被发现）

2. (PoW)工作量证明    (这里不细讲，主要是“任何人”都有记账的权利)
   
   

但是第一代区块链(BTC) 功能单一 只能转账， 而且BTC 价值太高了

新一代区块链  ETH  



功能就不仅仅只有转账的功能 它提供一个EVM虚拟机  (简单讲一下，我们也可以把EVM虚拟机理解为JAVA虚拟机 就是你可以写Solidity 的代码 然后在 ETH链上执行)

那我们就可以在链上去开发自己的应用俗称为(Dapp)



我们在BTC链上转账 就只能转账BTC  

ETH 赋予了我们 可以再链上开发应用的能力，那我们就可以在ETH链上做抽象开发出我们自己的代币     V神(ETH的父亲) 提出 ERC20 标准 实现ERC20标准 就可以在ETH上拥有我们自己的代币



简单的看看ERC20的实现

核心就是账户和余额的映射

```solidity
contract ERC20 is Context, IERC20, IERC20Metadata {
    mapping(address => uint256) private _balances; // 这就是一个简单的MAP key为 账户的地址 value 为 拥有代币的数量

    string private _name;  // 代币全称 
    string private _symbol;    // 代币简称
}
```

我们通过一个简单的   `mapping`  就可以实现一个账本 依托于 ETH的能力就变成了分布式账本  



### NFT とは　 什么是NFT    (ERC721 标准)

我们刚刚看见了 ERC20 代币  赋予了每一个人都有发币的权利 其价值取决于发行货币着的信用，如果是南非小国家的央行发行， 应为是在区块链上所有的交易记录都公开透明 是否可以高效的打击贪腐？？？



NFT  翻译成汉语就是 非同質化代幣  

通俗的立即就是 不是货币但是具有价值  比如艺术品



ERC721为 ETH链上的NFT标准   这次简单一点我们看看上层的抽象实践



```solidity
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

    // 可以挖矿     （内部的实现就是 更新底层两个 1. NFT ID 与 资源URL的关系 2. NFT ID 与 账户地址的关系）
    function safeMint(address to, string memory uri) public onlyOwner {
        // 当前NFT序列 递增
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri); // 设置Mappding tokenID 与URL关联
    }

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
```

其本质还是一个记账(MAP)，  ERC20 是代币和账户地址的关系， ERC721 的 资产(艺术品)和账户的关系   



我们更具以上代码部署了一个名为 DollerKillerToken 的NFT 地址 为 0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf

地址: https://polygonscan.com/address/0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf



我们调用挖矿函数: 地址 0xe916857755caee785ec694f096583f1be994892d 挖了 4个NFT

https://polygonscan.com/token/0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf?a=0xe916857755caee785ec694f096583f1be994892d



给地址 0xd1f84ec6652b315e77ebffa649f6742754574796 挖 1个 NFT:

https://polygonscan.com/token/0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf?a=0xd1f84ec6652b315e77ebffa649f6742754574796







### 部署后基础信息

#### NFT:

- 合约地址: 0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf
- 浏览器: https://polygonscan.com/address/0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf





FQA: 

- 前端获取 MetaMask 钱包信息
  
  - 阅读代码 `polygon_demo`

- 获取用户持有那些NFT    （这里以polygon链上的数据为例）
  
  - 用户的NFT的关系 不是在用户这里 我们最多只能获得用户的address
  
  - 用户和NFT的关系存储在   NFT只能合约的MAP当中
  
  - 最简单的方法就是爬取区块链浏览器已经分析好的数据 （自己做也可以就是要面对链上的全部数据和及时的更新链上数据 ）

- 获取某个NFT 目前所载的钱包地址 
  
  - 调用ERC721 `ownerOf` 函数

- NFT 上架 
  
  - https://docs.opensea.io/docs/metadata-standards     修改uri返回json 适配opensea metadata的标准

- NFT 派奖
  
  - 调用ERC721 `safeMint` 函数
  
  - 玩法的策划是非常重要的， 可以参考下 唯一艺术　と　THETAN ARENA　的发展史        
    
    - 合成      三个NFT 合成一个 NFT
    
    - VIP权益     经验加成
    
    - 饥饿营销   每周出少量空头    然后4~5倍的价格回收   
    
    - NFT本来不值钱， 但是认为他能赚钱，暴力的人多了。它就值钱了。
