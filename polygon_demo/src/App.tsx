import { useState, useEffect } from 'react'
import Robot from "./components/Robot"
import axios from "axios";
import styles from './App.module.css'

function App() {
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
    <div className={styles.app}>
      <span>count: {count}</span>
      <button onClick={()=>{
       setCount(count+1)
      }}>Click</button>

        {
            loading ?
                (<div>Loding</div>) : (
                    <div className={styles.robotList}>
                        {robotGallery.map((r)=>(
                            <Robot id={r.id} email={r.email} name={r.name}></Robot>
                        ))}
                    </div>
                )
        }
    </div>
  )
}

export default App
