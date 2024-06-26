import noteladHeader from '../public/noteladheader.svg'
import noteladLogo from '../public/noteladlogo.png'
import './Home.css'
export default function Home() {
  return (
    <div>
      <div className="flexCon">
      <img className="noteladHeader" src={noteladHeader}></img>
      <img className="noteladLogo" src={noteladLogo}></img>
      </div>
    </div>
  )
}
