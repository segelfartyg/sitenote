import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom";
import { NOTELAD_BASE_API } from "./consts/endpoints";
import "./styles/Profile.css"
import "./styles/common.css"
import NoteLogoDefault from '../public/singlen.svg'
import EditLogo from '../public/editicon.svg'
import ExitLogo from '../public/cross.svg'
import LinkIcon from '../public/link.svg'





interface getFindingsResponse{
  authenticated: boolean;
  findingsList: Finding[]
}


interface deleteFindingResponse{
  findingId: string
}

interface Finding{
  FindingId: string;
  Name: string;
  UserId: string;
  Content: string;
  Link: string;
}



async function getUserFindings(): Promise<getFindingsResponse>{
  const response = await fetch(NOTELAD_BASE_API + "/finding/user/all", {
    method: "GET", 
    cache: "no-cache", 
    mode: "cors",
    redirect: "follow", 
    referrerPolicy: "no-referrer",
    credentials: "include"
  });
  let result: getFindingsResponse = {
    authenticated: false,
    findingsList: []
  } ;

  if(response.status == 401){
    result.findingsList = []
    result.authenticated = false
    return result;
  }
  result.findingsList = await response.json();
  result.authenticated = true
  return result;
  
}

export default function Profile() {

  const navigate = useNavigate()
  const [findings, setFindings] = useState<Finding[]>([])
  
  useEffect(() => {

    let authRes = async () => await getUserFindings().then(res => {
      if(res.authenticated == false){
        navigate("/")
      }
      else{

        if(res.findingsList != null){

          let newFindingsList: Finding[] = []
          res.findingsList.forEach(finding => {
            newFindingsList.push(finding)
          })
          setFindings(newFindingsList)
        }
       
      }
    })
    authRes()
    setFindings(findings)
  }, [])


  async function deleteUserFinding(findingId: string){

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
    let result = await response.json();
    let currentFindings = findings;
    let newList = currentFindings.filter(f => f.FindingId != result)

    setFindings(newList)
  }
  

  return (
    <div className="Profile WaveBackground">
      <div className="flexCon">
      <h1 className="profileHeader">Your findings</h1>
      </div>   
      {findings.map(f => {
        return <div className="flexCon noteCon" key={f.FindingId}>
        <img className="noteLogo" src={NoteLogoDefault}></img>
        <p className="nameFindingPara">{f.Name}</p>
        <a href={"/finding/" + f.FindingId}><img className="editLogo" src={EditLogo} onClick={() => {navigate("/finding/" + f.FindingId)}}></img></a>
        <a href={"/profile"}><img className="exitLogo" onClick={e => deleteUserFinding(f.FindingId)} src={ExitLogo}></img></a>
        <a href={f.Link}><img className="linkLogo" src={LinkIcon} /></a>
        </div>
      })}

      </div>
  )
}
