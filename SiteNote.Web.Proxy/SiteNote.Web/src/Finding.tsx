import './styles/Finding.css'
import NoteLogoDefault from '../public/singlen.svg'
import ExitLogo from '../public/cross.svg'
import { useEffect, useState } from 'react'
import { useNavigate, useParams} from 'react-router-dom';
import { NOTELAD_BASE_API } from './consts/endpoints';
import EditLogo from '../public/editicon.svg'
import ShareIcon from "../public/share.svg"

export default function Finding() {


    interface Finding{
        FindingId: string;
        Name: string;
        UserId: string;
        Content: string;
        Link: string;
      }

    const { findingId } = useParams()

    const navigate = useNavigate()

    const [overlayStyling, setOverlayStyling] = useState(
        {display: "none"}
        )

    const [currentChosenFinding, setCurrentChosenFinding] = useState<Finding>()

    const onRemoveClickHandler = () => {
        if(overlayStyling.display == "none"){
            setOverlayStyling(
                {display: "flex"}
            )
        }
        else {
            setOverlayStyling(
                {display: "none"}
            )
        }
    }

    const onSuccessClickHandler = async () => {

        await deleteFinding()
        setOverlayStyling(
            {display: "none"}
        )
    }

    const onDeniedClickHandler =  () => {
        setOverlayStyling(
            {display: "none"}
        )
    }


    async function deleteFinding(){
        let body = {
            findingId: findingId
          }
      
          const response = await fetch(NOTELAD_BASE_API + "/finding/user/delete", {
            method: "POST", 
            cache: "no-cache", 
            mode: "cors",
            body: JSON.stringify(body),
            redirect: "follow", 
            referrerPolicy: "no-referrer",
            credentials: "include"
          });
        
          if(response.status == 401){
            navigate("/")
          }
          if(response.status == 200){
            navigate("/profile")
          }
          
    }
  

    useEffect(() => {
                
        let req = {findingId: findingId};
        

          (async () => {
            const data = await fetch(NOTELAD_BASE_API + "/finding/user", {
                method: "POST", 
                cache: "no-cache", 
                mode: "cors",
                redirect: "follow", 
                referrerPolicy: "no-referrer",
                credentials: "include",
                body:JSON.stringify(req)
              });
              let resJson = await data.json()

              let FindingRes: Finding = {
                FindingId: resJson.FindingId,
                Content: resJson.Content,
                Name: resJson.Name,
                Link: resJson.Link,
                UserId: resJson.UserId

            };

            setCurrentChosenFinding(FindingRes)

         })();

        
      }, [])
    

  return (
    <div className="Finding WaveBackground">
        <div className="flexCon specificFindingCon">
            <h2 className='findingHeader'>{currentChosenFinding?.Name}</h2>
            <p className='findingPara'>{currentChosenFinding?.Content}</p>
            <a href={currentChosenFinding?.Link} className='findingPara'>{currentChosenFinding?.Link}</a>


            <div className='breakArea'></div>
            <div className="findingPropArea">
            <img className="findingImage" src={NoteLogoDefault}></img>
            <img className="findingImage" src={ShareIcon}></img>
            <img className="findingImage" src={EditLogo}></img>
            </div>
            
            <img onClick={onRemoveClickHandler} className="removeFindingImage" src={ExitLogo}></img> 


            <div style={overlayStyling} className="popupOverlay">
            <div className='popupMenu'>
                DO YOU WANT TO DELETE THIS FINDING?
                <button onClick={onSuccessClickHandler}>yup</button>
                <button onClick={onDeniedClickHandler}>nope</button>
            </div>

            
        </div>
       
        </div>
    </div>
  )
}
