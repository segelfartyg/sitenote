
import './Header.css'
import { useState } from 'react'
import Logo from '../public/noteladlogo.png'
import Hamburger from '../public/hamburger.svg'
import Exit from '../public/exit.svg'

export default function Header() {

    const [menuSource, setMenuSource] = useState(Hamburger);
    const [popMenuStyle, setPopMenuStyle] = useState("none");

    function onMenuClickEventHandler(){

        if(menuSource == Hamburger){
            setMenuSource(Exit)
            setPopMenuStyle("flex")
            
        }
        else{
            setMenuSource(Hamburger)
            setPopMenuStyle("none")
        }
        
    }


  return (
    <div className="header">
        <div className="headerContent"> 
            <img className="logoImage" src={Logo}></img>
            {/* <p>home</p>
            <p>profile</p>
            <p>login</p> */}
            <img className="hamburger" onClick={onMenuClickEventHandler} src={menuSource}></img>
        </div>
        <div style={{display: popMenuStyle}} className="popMenu">
            <p>home</p>
            <p>profile</p>
            <p>login</p>
        </div>
    </div>
  )
}
