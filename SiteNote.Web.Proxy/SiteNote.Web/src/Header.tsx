
import './styles/Header.css'
import { useState } from 'react'
import Logo from '../public/noteladlogo.png'
import Hamburger from '../public/hamburgericon.svg'
import Exit from '../public/cross.svg'
import { useNavigate } from 'react-router-dom'

export default function Header() {

    const [menuSource, setMenuSource] = useState(Hamburger);
    const [popMenuStyle, setPopMenuStyle] = useState({
        display: "none",
        animation: "none"
    });
    const navigate = useNavigate()
    function onMenuClickEventHandler(){

        if(menuSource == Hamburger){
            setMenuSource(Exit)
            setPopMenuStyle({...popMenuStyle, display: "flex", animation: "forwards menuAnimation 0.05s"})
            
        }
        else{
            setMenuSource(Hamburger)
            setPopMenuStyle({...popMenuStyle, animation: "forwards menuAnimationBack 0.15s"})
        }
        
    }


    function navigatorFunction(path: string){
        setMenuSource(Hamburger)
            setPopMenuStyle({...popMenuStyle, animation: "forwards menuAnimationBack 0.15s"})
        navigate(path)
    }


  return (
    <div className='Header'>
        
    <div className="header">
        <div className="headerContent"> 
            <img className="logoImage" src={Logo}></img>
            {/* <p>home</p>
            <p>profile</p>
            <p>login</p> */}
            <img className="hamburger" onClick={onMenuClickEventHandler} src={menuSource}></img>
        </div>
        
    </div>
    <div style={popMenuStyle} className="popMenu">
    <p className='headerItem' onClick={() => navigatorFunction("/")}>home</p>
    <p className='headerItem' onClick={() => navigatorFunction("/profile")}>profile</p>
    <p className='headerItem' onClick={() => navigatorFunction("/login")}>login</p>
</div>

</div>
  )
}
