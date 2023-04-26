import { useState } from 'react'
import {BrowserRouter,Route,Routes} from "react-router-dom"
import './App.css'
import login from "./components/login"
import register from "./components/register"
function App() {

  return (
       <BrowserRouter>
    <Routes>
        <Route path='/' Component={login} />
        <Route path='/login' Component={login}/>
        <Route path="/register" Component={register}/>
        </Routes>
    </BrowserRouter>
  )
}

export default App
