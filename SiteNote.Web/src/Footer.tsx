import './Footer.css'
import SegelImage from '../public/segelfartyg.png'
export default function Footer() {
  return (
    <div className="footer">
        <p>privacy</p>
        <p>terms</p>
        <p>purpose</p>
        <img className="segelImage" src={SegelImage}></img>
    </div>
  )
}
