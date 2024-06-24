
import './Header.css'
import Logo from '../public/noteladlogo.png'
export default function Header() {
  return (
    <div className="header">
        <div className="headerContent"> 
            <img className="logoImage" src={Logo}></img>
            <p>home</p>
            <p>profile</p>
            <p>login</p>
        </div>
    </div>
  )
}
