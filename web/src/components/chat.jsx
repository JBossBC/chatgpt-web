/*
 * Here is the code to fulfill the query:
 */
import React, { useState,useEffect } from "react";
import { AlertTitle,Alert } from "@mui/material"
const chat = () => {
  const inputCSS="bg-white "
   const [userModal,setUserModal]  =useState(false)
   const [notice,setNotice]=useState(false)
   const [errInfo,setErrInfo]=useState({
    title:"",
    context:"",
})
const [infoList,setInfoList]=useState([]);
const [info,setInfo]=useState(null);
function sendInfo(){
  setInfoList(pre=>[...pre,packageInfo(info)],()=>{
   setInfo(null);
  });
  //向后端发送请求
}
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
function renderInfoList(){

}
function packageInfo(info){
  return (<div key={Date.now()} className="flex w-full mb-6 overflow-hidden flex-row-reverse">
    <div className="flex items-center justify-center flex-shrink-0 h-8 overflow-hidden rounded-full basis-8 ml-2">
      <img src="../../public/user.jpg" className=" h-8 w-8"></img>
    </div>
    <div className="overflow-hidden text-sm items-end">
      <p className="text-xs text-[#b4bbc4] text-right">{formatNowDate()}</p>
    <div className="flex items-end gap-1 mt-2 flex-row-reverse">
      <div className="text-black text-wrap min-w-[20px] rounded-md px-3 py-2 bg-[#d2f9d1] dark:bg-[#a1dc95]">
         <div className="leading-relaxed break-words">
            <div className="whitespace-pre-wrap">{info}</div>
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
        </div>
        <div className="flex flex-row w-full h-5/6">
                  <div className=" w-1/3 border-2 m-4 overflow-auto">
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
                  <div className=" w-2/3 flex-col flex justify-between">
                     <div className=" overflow-auto overflow-y-auto">
                        {infoList}
                     </div>
                     <div className=" flex flex-row">
                         <input type="text" className=" w-11/12 overflow-scroll rounded-md form-input" placeholder="请输入"  onChange={(info)=>setInfo(info)}></input>
                          <button className=" w-1/12  bg-sky-500 rounded-lg ml-4" onClick={sendInfo}>发送</button>
                     </div>

                  </div>
        </div>
   {userModal?<div className="">
    
   </div>:<></>}
   </>
  );
};

export default chat;