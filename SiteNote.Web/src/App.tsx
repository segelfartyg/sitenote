import './App.css'

import Header from './Header'
import Footer from './Footer'
import Home from './Home'
import Profile from './Profile'
import { BrowserRouter, Routes, Route } from "react-router-dom"
import { CredentialResponse, GoogleLogin } from '@react-oauth/google'


async function login(credRes: CredentialResponse){

  let req = {id_token: credRes.credential}
  const response = await fetch("http://localhost:9000" + "/login", {
      method: "POST", 
      cache: "no-cache", 
      mode: "cors",
      headers: {},
      redirect: "follow", 
      referrerPolicy: "no-referrer", 
      credentials: "include",
      body: JSON.stringify(req), 
    });
  
    if(response.status == 401){
      console.log("unauthorized")
    }
    else{

      console.log(await response.text()); 
    }
    
  }

function App() {

  return (

      <div className='container'>
        <Header/>
        <BrowserRouter>
        <Routes>
        <Route index element={<Home />} />
        <Route path="profile" element={<Profile />} />
        </Routes>
        </BrowserRouter>
        <GoogleLogin onSuccess={credentialres => {login(credentialres)}} onError={() => console.log("LOGIN FAILED")}></GoogleLogin>
        <Footer />
        
      </div>

  )
}

export default App
