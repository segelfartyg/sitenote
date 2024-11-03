import './styles/App.css'

import Header from './Header'
import Footer from './Footer'
import Home from './Home'
import Profile from './Profile'
import { BrowserRouter, Routes, Route, useNavigate } from "react-router-dom"
import { CredentialResponse, GoogleLogin } from '@react-oauth/google'
import { NOTELAD_BASE_API } from './consts/endpoints'
import Finding from './Finding'
import PostItNotes from './PostItNotes'




function App() {

  const navigate = useNavigate();

  async function login(credRes: CredentialResponse){

  
    let req = {id_token: credRes.credential}
    const response = await fetch(NOTELAD_BASE_API + "/login", {
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
        navigate("/")
      }
      else{
        navigate("/profile")
      }
      
    }

  return (

      <div className='container'>
        <Header/>
        <Routes>
        <Route index element={<Home onLogin={login}/>} />
        <Route path="profile" element={<Profile />} />
        <Route path="finding/:findingId" element={<Finding />} />
        <Route path="postit" element={<PostItNotes />} />
        </Routes>
        <Footer />
        
      </div>

  )
}

export default App
