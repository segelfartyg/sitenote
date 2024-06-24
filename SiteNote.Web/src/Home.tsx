import noteladHeader from '../public/noteladheader.svg'
import noteladLogo from '../public/noteladlogo.png'
import './Home.css'
export default function Home() {
  return (
    <div>
      <div className="flexCon">
      <img className="noteladHeader" src={noteladHeader}></img>
      <h1>Welcome to notelad.com</h1>
      <img className="noteladLogo" src={noteladLogo}></img>
      </div>
    </div>
  )
}
