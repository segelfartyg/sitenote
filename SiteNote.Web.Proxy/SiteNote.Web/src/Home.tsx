import { CredentialResponse, GoogleLogin } from '@react-oauth/google'
import noteladHeader from '../public/noteladheader.svg'
import noteladLogo from '../public/noteladlogo.png'
import './styles/Home.css'
import "./styles/common.css"

interface loginProps {
  /** The text to display inside the button */
  onLogin: (credRes: CredentialResponse) => void;
}

export default function Home({onLogin}:loginProps) {
  return (
    <div className='Home WaveBackground'>

    
      <div className='flexCon'>
      <h2 className='secondTitle'>NOTELAD</h2>
      <p className="para">This is the NoteLad homepage. sign in to take charge of your own NoteLad universe!</p>
      </div>

      <div className="flexCon">
      <img className="noteladLogo" src={noteladLogo}></img>
      </div>

      <div className='flexCon'>
      <GoogleLogin onSuccess={credRes => {onLogin(credRes)}} onError={() => console.log("LOGIN FAILED")}></GoogleLogin>
      </div>

    </div>
  )
}
