import './App.css'

import Header from './Header'
import Footer from './Footer'
import Home from './Home'
import Profile from './Profile'
import { BrowserRouter, Routes, Route } from "react-router-dom"


function App() {

  return (
    <>
      <div>
        <Header/>
        <BrowserRouter>
        <Routes>
        <Route index element={<Home />} />
        <Route path="profile" element={<Profile />} />
        </Routes>
        </BrowserRouter>
        <Footer />
        
      </div>
    </>
  )
}

export default App
