import { useState, useEffect } from 'react'
import axios from "axios";
import {Button, createTheme, ThemeProvider} from '@mui/material'
import {
    BrowserRouter,
    Routes,
    Route,
    Link,
} from "react-router-dom";
import {HomePage} from "./pages/home";
import styles from './App.module.css'

function App() {
    const theme = createTheme({
        palette: {
            mode: "light",
            primary: {
                main: "#0fa6a2",
            },
            secondary: { main: "#8eb8e7" },
            background: {
                paper: "",
            },
        },
        shape: { borderRadius: 4 },
    });
    const [count,setCount] = useState<number>(0)
    const [robotGallery, setRobotGallery] = useState<any>([])
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(()=>{
        document.title = `点击${count}次`
    }, [count]) // 监听count 发生变化时发送副作用

    useEffect(()=>{
        axios.get("https://jsonplaceholder.typicode.com/users").
        then(response=>{
            console.log(response.data)
            setRobotGallery(response.data)
            setLoading(false)
        })
    },[]) // 当为空时 仅当挂载UI时加载

  return (
      <ThemeProvider theme={theme}>
          {/*<ul className={styles.list}>*/}
          {/*    /!*<li><Link to="/">Home</Link></li>*!/*/}
          {/*    <li><Button variant="contained">空投</Button></li>*/}
          {/*</ul>*/}

          <div className="App">
              <BrowserRouter>
                  <Routes>
                      <Route path="/" element={<HomePage/>}/>
                  </Routes>
              </BrowserRouter>
          </div>
      </ThemeProvider>
  )
}

export default App
