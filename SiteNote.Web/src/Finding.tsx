import './styles/Finding.css'
import NoteLogoDefault from '../public/singlen.svg'
import ExitLogo from '../public/cross.svg'
import { useEffect, useState } from 'react'
import { useParams} from 'react-router-dom';
import { NOTELAD_BASE_API } from './consts/endpoints';


export default function Finding() {


    interface Finding{
        FindingId: string;
        Name: string;
        UserId: string;
        Content: string;
        Link: string;
      }

    const { findingId } = useParams()

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

    const onSuccessClickHandler = () => {
        setOverlayStyling(
            {display: "none"}
        )
    }

    const onDeniedClickHandler = () => {
        setOverlayStyling(
            {display: "none"}
        )
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
            <h2>{currentChosenFinding?.Name}</h2>
            <p className='findingPara'>{currentChosenFinding?.Content}</p>
            <img className="findingImage" src={NoteLogoDefault}></img>
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
