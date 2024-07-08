import './styles/Footer.css'
import SegelImage from '../public/segelfartyg.png'
export default function Footer() {
  return (
    <div className="footer">
        <p className="footerItem">privacy</p>
        <p className="footerItem">terms</p>
        <p className="footerItem">purpose</p>
        {/* <img className="segelImage" src={SegelImage}></img> */}
    </div>
  )
}
