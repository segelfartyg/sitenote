import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom";


interface getFindingsResponse{
  authenticated: boolean;
  findingsList: Finding[]
}

interface Finding{
  FindingId: string;
  Name: string;
  UserId: string;
  Content: string;
  Link: string;
}


async function getUserFindings(): Promise<getFindingsResponse>{
  const response = await fetch("http://localhost:9000" + "/finding/user/all", {
    method: "GET", 
    cache: "no-cache", 
    mode: "cors",
    redirect: "follow", 
    referrerPolicy: "no-referrer",
    credentials: "include"
  });
  console.log("RETRIEVED USER INFO:")
  console.log(response)

  let result: getFindingsResponse = {
    authenticated: false,
    findingsList: []
  } ;
  
  
  result.findingsList = await response.json();
  console.log(response.status);

  if(response.status == 401){
    result.authenticated = false
    return result;
  }
  result.authenticated = true
  return result;
  
}

export default function Profile() {

  const navigate = useNavigate()
  const [findings, setFindings] = useState<Finding[]>([])
  
  useEffect(() => {

    let authRes = async () => await getUserFindings().then(res => {
      console.log("RESPONSE SIR:")
      console.log(res);
      if(res.authenticated == false){
        navigate("/")
      }
      else{

        let newFindingsList: Finding[] = []
        res.findingsList.forEach(finding => {
          newFindingsList.push(finding)
        })
        setFindings(newFindingsList)
      }
    })
    authRes()
    setFindings(findings)
  }, [])
  

  return (
    <div>
      <h1>NoteLad Profile</h1>
      <ol>
      {findings.map(f => {
        return <li key={f.FindingId}>{f.Name}, {f.Content}, {f.Link}, {f.UserId}</li>
      })}
      </ol>
    </div>
  )
}
