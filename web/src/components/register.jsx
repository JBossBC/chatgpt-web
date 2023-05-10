import {Link} from "react-router-dom"
import { useState,useEffect } from "react"
import { AlertTitle,Alert } from "@mui/material"

const register = () => {
    const [notice,setNotice]=useState(false)
    const [verification,setverification]=useState(false)
    const [timeSend,setTimeSend]=useState(60)  
    let nowTimes =60
    useEffect(()=>{
        if (verification){
         const intervalTimeSend=setInterval(()=>{
            nowTimes--
                setTimeSend(nowTimes)
            },1000)  
         setTimeout(() => {
            console.log("hello")
            clearInterval(intervalTimeSend)
            setverification(false)
            setTimeSend(60)
            nowTimes=60

         }, 60000);
        }
    },[verification])
    const [user,setUser]=useState({
        username:"",
        password:"",
        phoneNumber: "",
        verification:"",
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
    function toRegister(){
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
    function sendVerification(){
        if (user.username==null ||user.username == ""){
            evalErrorAlert("错误","请先填写账号")
            return
        }
        if (user.password==null ||user.password==""){
            evalErrorAlert("错误","请先填写密码")
            return
        }
           setverification(true);
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
                <div className="absolute right-60 bottom-32 w-96 mr-15 flex float-right content-center form rounded-6">
                    <div className=" justify-center self-center flex-initial bg-white rounded-sm">
                        <div className="flex justify-center pt-10 text-3xl pb-10">
                            注册
                        </div>
                            <div className=" bg-white form-input border-0 flex justify-center">
                             <label htmlFor="usename" className="white mr-3  text-lg align-middle">账号:</label>   
                            <input type="text" id="username" className="w-72 border-1 rounded-full" onChange={(event)=>setUser({...user,username:event.target.value})}></input>
                            </div>
                            <div  className=" bg-white form-input border-0 flex justify-center">
                                <label htmlFor="password"  className="white  mr-3 text-lg align-middle">密码:</label>
                                <input type="password" id="password" className="w-72 border-1 rounded-full" onChange={(event)=>setUser({...user,password:event.target.value})}></input>
                            </div>
                            <div  className=" bg-white form-input border-0 flex justify-center">
                                <label htmlFor="phoneNumber"  className="white  mr-3 text-lg align-middle">手机号:</label>
                                <input type="tel" id="phoneNumber" className="w-72 border-1 rounded-full" onChange={(event)=>setUser({...user,phoneNumber:event.target.value})}></input>
                            </div>
                            <div  className=" bg-white form-input border-0 flex justify-center ">
                                <label htmlFor="verification"  className="white  mr-3 text-lg align-middle ">验证码:</label>
                                <input type="text" id="verification" className="w-42 border-1 rounded-full" onChange={(event)=>setUser({...user,verification:event.target.value})}></input>
                                {!verification?<button className="w-1/5 button rounded-full ml-2" onClick={sendVerification}>发送</button>:<button className="w-1/3 button rounded-full" disabled="disabled">{timeSend}后重试</button>}
                            </div>
                            <div className="flex h-10 mt-10 mb-3 justify-center">
                                <button onClick={toRegister} className="w-4/5  rounded-full text-center button ">注册</button>
                            </div>
                            <Link to="/login">
                            <div className="flex h-10 justify-center">
                            <button className="w-4/5 text-center rounded-full button bg-blue-600">返回</button>
                            </div>
                            </Link>
                    </div>
                </div>
           </div>
        )
};

export default register