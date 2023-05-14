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

  return (
   <>
   <div className="w-full h-full">
        <div className="w-full h-1/6 min-h-full bg-blue-300">
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
                  <div className=" w-1/3 border-2 m-4 overflow-clip">
                      <div>
                        <p className="  text-center text-2xl">模型</p>
                      </div>
                      <div className=" h-1/2 ">
                           <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                     Model
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true} />
                               </div>
                            </div>   
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-row justify-between ">
                                  <span className=" mr-6">
                                     Max_Tokens
                                  </span>
                                  <input type=" text " className=" form-input"disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Temperature
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Top_p
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Presence_penalty
                                  </span>
                                  <input type=" text " className=" form-input"disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Frequency_penalty
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  N
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true} />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Stream
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Stop
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Logit_bias
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Role
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Output_format
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                            <div className=" mt-2 mb-2 px-2 ">
                               <div className=" felx flex-col ">
                                  <span className=" mr-6">
                                  Keep_context
                                  </span>
                                  <input type=" text " className=" form-input" disabled={true}  />
                               </div>
                            </div> 
                      </div>
                  </div>
                  <div className=" w-2/3 flex-col flex">
                     <div className=" h-5/6">
                        {}
                     </div>
                     <div className=" h-1/6 flex flex-row">
                         <input type="text" className=" overflow-scroll" value={(info)=>sendInfo(info)}></input>
                         <button className="" onClick={sendInfo}>发送</button>
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