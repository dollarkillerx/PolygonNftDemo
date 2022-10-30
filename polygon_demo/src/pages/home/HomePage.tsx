import React, {useEffect, useState} from "react"
import styles from './HomePage.module.css'
import Typography from '@mui/material/Typography';
import axios from "axios"
import {Tabs, Tab, Box, Button, Alert} from '@mui/material'
import {ethers} from "ethers"

interface TabPanelProps {
    index: number;
    value: number;
}

function TabPanel(props: TabPanelProps) {
    const { value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`simple-tabpanel-${index}`}
            aria-labelledby={`simple-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box sx={{ p: 3 }}>
                    <Switch value={value}></Switch>
                </Box>
            )}
        </div>
    );
}

function Switch(props: TabPanelProps) {
    const { value, index, ...other } = props;

    switch (value) {
        case 0:
            return Home()
        case 1:
            return Airdrop()
        case 2:
            return Other()
    }
}

function Home() {
    // 存储eth账户信息
    const [account, setAccount] = useState();
    const [balance, setBalance] = useState();
    const [provider, setProvider] = useState();
    const [nfts, setNfts] = useState<any>([]);

    // 链接钱包
    function connectOnClick() {
        //  浏览器安装了 metamask 会向浏览器注入JS代码 window.ethereum
        if (!window.ethereum) {
            alert("请安装Metamask钱包")
            return
        }

        const connection = async() => {
            // 如果您想使用原生API可以参考这个: https://docs.metamask.io/guide/rpc-api.html#ethereum-json-rpc-methods
            // const accounts = await ethereum.request({
            //     method: 'eth_requestAccounts'
            // })
            // 我们这里使用 ethers  它可以兼容多个钱包 : https://docs.ethers.io/
            // 初始化metamask钱包
            const providerWeb3 = new ethers.providers.Web3Provider(window.ethereum);
            // 登录钱包获取授权
            providerWeb3.send("eth_requestAccounts", []).
                then((accounts) => {
                    // 注意 生产环境需要判断错误
                    setAccount(accounts[0])
                    setProvider(providerWeb3)

                // login sign
                const message = "Login to the website"
                ethereum.request({
                    method: "personal_sign",
                    params: [
                        message,
                        accounts[0]
                    ]
                }).then((signature) => {
                    console.log(signature)
                    console.log(message)
                    console.log(accounts[0])
                })


                axios.get(`/api/v1/address/${accounts[0]}/nfts`).
                then(data=>{
                    console.log(data)

                        if (data.data.code != "0") {
                            alert("请求失败")
                            return
                        }
                        setNfts(data.data.data)
                    console.log(data.data.data)
                }).catch(err=>{
                    alert("错误: "+ err)
                })
            })
        }

        connection()
    }

    // 副作用函数 初始化页面时 加载
    useEffect(()=>{
        if (!window.ethereum) {
            alert("请安装Metamask钱包")
            return
        }

        // metamask钱包 原生API 监听账户变动
        ethereum.on("accountsChanged",function (accountsChange) {
            setAccount(accountsChange[0])
        })
    },[])

    // 副作用函数 监听 account变化
    useEffect(()=>{
       if (!window.ethereum || !account || !provider) {
           return
       }

       provider.getBalance(account).then((result) => {
           setBalance(ethers.utils.formatEther(result))
           console.log(balance)
       })
    },[account])

    return (
        <div>

            {
                account == undefined ? (
                    <Button variant="contained" onClick={connectOnClick}>Login</Button>
                ): (<span>已经链接钱包: {account}  余额: {balance}</span>)
            }

            <hr/>

            {
                nfts.length == 0 ? (
                    <span>当前账户 占无https://polygonscan.com/token/0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf  NFT</span>
                ): (

                    <div>
                        <span>https://polygonscan.com/token/0xa10a5e7e2b2cc6fcaf4e74b00a163b10ee060ebf  NFT</span>

                        {nfts.map((r)=>(
                            <li>NTF ID:  {r} </li>
                        ))}

                    </div>


                )
            }

        </div>
    )
}

function Airdrop() {
    // 副作用函数 初始化页面时 加载
    useEffect(()=>{
        if (!window.ethereum) {
            alert("请安装Metamask钱包")
            return
        }
    },[])

    function getAirdrop() {
        if (!window.ethereum) {
            alert("请安装Metamask钱包")
            return
        }

        const getAir = async() => {
            const providerWeb3 = new ethers.providers.Web3Provider(window.ethereum);
            providerWeb3.send("eth_requestAccounts", []).
            then((accounts) => {
                if (accounts == undefined || accounts.length == 0) {
                    alert("请授权账户")
                    return
                }

                // 注意 生产环境需要判断错误
                let account = accounts[0]
                axios.post("/api/v1/airdrop", {
                    "address": account
                }).then(data => {
                    console.log(data)
                    if (data.data.code == "0") {
                        alert("空投领取成功 请前往opensea 查看空投")
                    }
                }).catch((err)=> {
                    alert("错误: " + err)
                })

            })
        }

        getAir()
    }

    return (
       <div>
           <div>查看空投信息，前往opensea 隐藏的NFT</div>
           <Button variant="contained" onClick={getAirdrop}>领取空投</Button>
       </div>
    )
}

function Other() {
    const [nfts, setNfts] = useState<any>([]);

    function otherNet() {
        axios.get("/api/v1/nft_holder").then(data => {
            console.log(data)
            if (data.data.code != "0") {
                alert("获取失败")
                return
            }
            setNfts(data.data.data)
        }).catch((err)=> {
            alert("错误: " + err)
        })
    }

    useEffect(()=>{
        otherNet()
    },[])

    return (
        <div>
            <div>NFT 持有者 </div>

            {nfts.map((r)=>(
                <li>NTF ID:  {r.key} ,   持有者: {r.value}</li>
            ))}
        </div>
    )
}

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

export const HomePage: React.FC = () => {
    const [value, setValue] = React.useState(0);

    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    return (
        <Box sx={{ width: '100%' }}>
            <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
                    <Tab label="Home" {...a11yProps(0)} />
                    <Tab label="空投" {...a11yProps(1)} />
                    <Tab label="查询" {...a11yProps(2)} />
                </Tabs>
            </Box>

            <TabPanel value={value} index={0}></TabPanel>
            <TabPanel value={value} index={1}></TabPanel>
            <TabPanel value={value} index={2}></TabPanel>
        </Box>
    );
}