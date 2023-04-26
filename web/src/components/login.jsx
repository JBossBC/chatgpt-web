import "./login.css"
import {Link} from "react-router-dom"
import { useState,useEffect } from "react"
import { AlertTitle,Alert } from "@mui/material"
const login =()=>{
    const [notice,setNotice]=useState(false)
    const [user,setUser]=useState({
        username:"",
        password:"",
    });
    const [errInfo,setErrInfo]=useState({
        title:"",
        context:"",
    })
    function evalErrorAlert(title,context){
        setNotice(true)
        setErrInfo({
            ...errInfo,
            tilte:title,
            context:context,
        })
    }
    useEffect(()=>{
        if (notice){
           setTimeout(()=>{
            setNotice(false)
           },3000)
        }
    },[notice])
    function toLogin(){
        console.log(user)
        if (user.username==null ||user.username == ""){
            evalErrorAlert("错误","用户名不能为空")
            return
        }
        if (user.password==null ||user.password==""){
            evalErrorAlert("错误","密码不能为空")
            return
        }
        //TODO业务处理
    }
    return (
        <div className="w-full h-full ">
            <div className="h-1/4">
                {notice?
                <div className="mt-2 w-1/3 h-1/3 flex float-right">
            <Alert severity="error" className="w-full" onClose={()=>{setNotice(false)}} >
               <AlertTitle>{errInfo.title}</AlertTitle>
                 {errInfo.context}
             </Alert>
             </div>:<></>}
            </div>
            <div className="absolute right-60 w-96 mr-15 flex float-right content-center form rounded-6">
                <div className=" justify-center self-center flex-initial bg-white rounded-sm">
                    <div className="flex justify-center text-3xl pb-10">
                        登录
                    </div>
                        <div className=" bg-white form-input border-0 flex justify-center">
                         <label htmlFor="usename" className="white mr-3  text-lg">账号:</label>   
                        <input type="text" id="username" className="w-72 border-1 rounded-full" onChange={(event)=>setUser({...user,username:event.target.value})}></input>
                        </div>
                        <div  className=" bg-white form-input border-0 flex justify-center">
                            <label htmlFor="password"  className="white  mr-3 text-lg">密码:</label>
                            <input type="password" id="password" className="w-72 border-1 rounded-full" onChange={(event)=>setUser({...user,password:event.target.value})}></input>
                        </div>
                        <div className="flex h-10 mt-10 mb-3 justify-center">
                            <button onClick={toLogin} className="w-4/5  rounded-full text-center button ">登录</button>
                        </div>
                        <Link to="/register">
                        <div className="flex h-10 justify-center">
                        <button className="w-4/5 text-center rounded-full button bg-blue-600">注册</button>
                        </div>
                        </Link>
                </div>
            </div>
       </div>
    )

}


export default login