import noteladHeader from '../public/noteladheader.svg'
import noteladLogo from '../public/noteladlogo.png'
import './Home.css'
export default function Home() {
  return (
    <div className='Home'>

    
      <div className='flexCon'>
      <h2 className='secondTitle'>NOTELAD</h2>
      <p className="para">This is the NoteLad homepage. sign in to take charge of your own NoteLad universe.</p>
      </div>


      <div className="flexCon">
      {/* <img className="noteladHeader" src={noteladHeader}></img> */}
      <img className="noteladLogo" src={noteladLogo}></img>
      </div>


      {/* <div className='flexCon'>
      <p className="para">This is the NoteLad homepage, sign in to take charge of your own NoteLad universe.</p>
      </div> */}
    </div>
  )
}
