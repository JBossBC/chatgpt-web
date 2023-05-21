/*
 * Here is the code to fulfill the query:
 */
import React, { useState,useEffect,useRef,useContext } from "react";
import { AlertTitle,Alert } from "@mui/material"

import axios from "axios"



const chat = () => {
  const baseUrl=useContext("BackendURL");
   const [userModal,setUserModal]  =useState(false)
   const [notice,setNotice]=useState(false)
   const [errInfo,setErrInfo]=useState({
    title:"",
    context:"",
})
const [infoList,setInfoList]=useState([]);
const [info,setInfo]=useState(null);
const textareaRef=useRef(null);
const CommitbuttonRef=useRef(null);
const newInfoElement=useRef(null);
function sendInfo(){
  if(info == null || info == ""){
    evalErrorAlert("error", "输入信息不能为空");
    return; 
  }
  //修改前端状态
  setInfoList(pre=>[...pre,packageInfo()]);
  setInfo(null);
  textareaRef.current.value= "";
  CommitbuttonRef.current.disabled=true;
  //向后端发送请求
   axios.get(baseUrl+"/chat").then((response)=>{
    console.log(response);
   })
  setInfoList(pre=>[...pre,packageAnswer("hello")]);
  //前端状态恢复
  CommitbuttonRef.current.disabled=false;
}
useEffect(()=>{
  if(newInfoElement.current!=null){
    console.log("hello");
    newInfoElement.current.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
    }
},[infoList])
function formatNowDate(){
   const now = new Date();
   const options = {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      // weekday: 'long',
      hour: 'numeric',
      minute: 'numeric',
      second: 'numeric'
    };
   const localDate = now.toLocaleDateString('zh-CN', options);
   return localDate
}
function packageAnswer(answer){
  const newElement=(<div key={Date.now()} className="flex w-full mb-6 overflow-hidden flex-row">
    <div className="flex items-center justify-center flex-shrink-0 h-8 overflow-hidden rounded-full basis-8 ml-2">
    <svg className=" h-8 w-8" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" aria-hidden="true" width="1em" height="1em"><path d="M29.71,13.09A8.09,8.09,0,0,0,20.34,2.68a8.08,8.08,0,0,0-13.7,2.9A8.08,8.08,0,0,0,2.3,18.9,8,8,0,0,0,3,25.45a8.08,8.08,0,0,0,8.69,3.87,8,8,0,0,0,6,2.68,8.09,8.09,0,0,0,7.7-5.61,8,8,0,0,0,5.33-3.86A8.09,8.09,0,0,0,29.71,13.09Zm-12,16.82a6,6,0,0,1-3.84-1.39l.19-.11,6.37-3.68a1,1,0,0,0,.53-.91v-9l2.69,1.56a.08.08,0,0,1,.05.07v7.44A6,6,0,0,1,17.68,29.91ZM4.8,24.41a6,6,0,0,1-.71-4l.19.11,6.37,3.68a1,1,0,0,0,1,0l7.79-4.49V22.8a.09.09,0,0,1,0,.08L13,26.6A6,6,0,0,1,4.8,24.41ZM3.12,10.53A6,6,0,0,1,6.28,7.9v7.57a1,1,0,0,0,.51.9l7.75,4.47L11.85,22.4a.14.14,0,0,1-.09,0L5.32,18.68a6,6,0,0,1-2.2-8.18Zm22.13,5.14-7.78-4.52L20.16,9.6a.08.08,0,0,1,.09,0l6.44,3.72a6,6,0,0,1-.9,10.81V16.56A1.06,1.06,0,0,0,25.25,15.67Zm2.68-4-.19-.12-6.36-3.7a1,1,0,0,0-1.05,0l-7.78,4.49V9.2a.09.09,0,0,1,0-.09L19,5.4a6,6,0,0,1,8.91,6.21ZM11.08,17.15,8.38,15.6a.14.14,0,0,1-.05-.08V8.1a6,6,0,0,1,9.84-4.61L18,3.6,11.61,7.28a1,1,0,0,0-.53.91ZM12.54,14,16,12l3.47,2v4L16,20l-3.47-2Z" fill="currentColor"></path></svg>
    </div>
    <div className="overflow-hidden text-sm items-start">
      <p className="text-xs text-[#b4bbc4] text-right">{formatNowDate()}</p>
    <div className="flex items-start gap-1 mt-2 flex-row">
      <div className="text-black text-wrap min-w-[20px] rounded-md px-3 py-2 bg-[#d2f9d1] dark:bg-[#a1dc95]">
         <div className="leading-relaxed break-words">
            <div ref={newInfoElement} className="whitespace-pre-wrap">{answer}</div>
         </div>
      </div>
      <div className="flex flex-col">
           <button className="transition text-neutral-300 hover:text-neutral-800 dark:hover:text-neutral-200">
           <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" class=" iconify iconify--ri" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M12 3c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2Zm0 14c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2Zm0-7c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2Z"></path></svg>
           </button>
      </div>
    </div>
    </div>
  </div>)
  return newElement;


}
function packageInfo(){
  const newElement=(<div key={Date.now()} className="flex w-full mb-6 overflow-hidden flex-row-reverse">
    <div className="flex items-center justify-center flex-shrink-0 h-8 overflow-hidden rounded-full basis-8 ml-2">
      <img src="../../public/user.jpg" className=" h-8 w-8"></img>
    </div>
    <div className="overflow-hidden text-sm items-end">
      <p className="text-xs text-[#b4bbc4] text-right">{formatNowDate()}</p>
    <div className="flex items-end gap-1 mt-2 flex-row-reverse">
      <div className="text-black text-wrap min-w-[20px] rounded-md px-3 py-2 bg-[#d2f9d1] dark:bg-[#a1dc95]">
         <div className="leading-relaxed break-words">
            <div ref={newInfoElement} className="whitespace-pre-wrap">{info}</div>
         </div>
      </div>
      <div className="flex flex-col">
           <button className="transition text-neutral-300 hover:text-neutral-800 dark:hover:text-neutral-200">
           <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" class=" iconify iconify--ri" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M12 3c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2Zm0 14c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2Zm0-7c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2Z"></path></svg>
           </button>
      </div>
    </div>
    </div>
  </div>)
  return newElement;

}
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
function initPage(){
   // get userInfo
   
   //get paramsInfo
}
initPage()
  return (
   <>
        <div className="w-full h-1/6 ">
          {notice?
          <div className="absolute right-1/2"> 
          <Alert severity="error" className="w-full" onClose={()=>{setNotice(false)}} >
               <AlertTitle>{errInfo.title}</AlertTitle>
                 {errInfo.context}
             </Alert>
          </div>:<></>
}
          <div className="flex flex-row border-2 rounded-md float-right mr-6 w-44 h-16">
            <button onClick={()=>{setUserModal(true)}}>
               <img src="../../public/user.jpg" className="w-16 h-16 rounded-md"></img>
               </button>
               <div className="w-28 h-16 flex flex-col align-middle">
                <span className="h-6 w-28 self-start text-center">用户名:</span>
               <span className="h-6 w-28 self-center overflow-hidden">18080705675</span>
               </div>
          </div>
       {/* {header} */}
        </div>
        <div className="flex flex-row w-full h-5/6">
                  <div className=" w-1/3 border-2 m-4 overflow-auto rounded-md">
                      <div className=" mt-4 mb-8">
                        <p className="  text-center text-2xl">模型</p>
                      </div>
                      <div className=" h-1/2 ">
                           <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                     Model
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true} />
                               </div>
                            </div>   
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" flex justify-between ">
                                  <span className=" mr-6">
                                     Max_Tokens
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Temperature
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Top_p
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Presence_penalty
                                  </span>
                                  <input type=" text " className="w-2/3  form-input"disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" flex justify-between">
                                  <span className=" mr-6">
                                  Frequency_penalty
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  N
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true} />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Stream
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between">
                                  <span className=" mr-6">
                                  Stop
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between">
                                  <span className=" mr-6">
                                  Logit_bias
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Role
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Output_format
                                  </span>
                                  <input type=" text " className=" w-2/3 form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className="  flex justify-between ">
                                  <span className=" mr-6">
                                  Keep_context
                                  </span>
                                  <input type=" text " className="w-2/3  form-input" disabled={true}  />
                               </div>
                            </div> 
                      </div>
                  </div>
                  <div className=" w-2/3 flex-col flex justify-between py-8  border-2 rounded-md">
                     <div className=" overflow-auto overflow-y-auto border-b-2  w-full h-full" >
                        {infoList}
                     </div>
                     <div className="flex flex-row w-full h-32">
                         {/* <input type="text" className=" w-11/12 overflow-scroll rounded-md form-input" placeholder="请输入"  onChange={(info)=>setInfo(info)}></input>
     */}
     <textarea ref={textareaRef} className=" h-full w-10/12 resize-none m-0 focus:outline-0" onInput={(event)=>setInfo(event.target.value)} ></textarea>
     <div className=" w-2/12 h-full flex items-end ">
                          <button ref={CommitbuttonRef} className="  bg-sky-500 rounded-lg ml-4 w-16 h-9" onClick={sendInfo}>发送</button>
                          </div>
                     </div>

                  </div>
        </div>
   {userModal?<div className="">
    
   </div>:<></>}
   </>
  );
};


export default chat;