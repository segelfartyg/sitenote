let titleInputField = window.document.getElementById("titleField");
let urlInputField = window.document.getElementById("urlField");
let loginDiv = window.document.getElementById("loginDiv");
let loggedInDiv = window.document.getElementById("loggedInDiv");
let newFindingDiv = window.document.getElementById("newFindingDiv");
let findingsDiv = window.document.getElementById("findingsDiv");
let postFindingBtn = window.document.getElementById("postFindingBtn");
let contentInputField = window.document.getElementById("contentField");

// INITIALIZING THE EXTENSION
(async () => {

  console.log("STARTING EXTENSION...")

  titleInputField.readOnly = true;
  urlInputField.readOnly = true;

  postFindingBtn.addEventListener("click", async (e) => {
    await createFinding()
  })

  let userInfo = await getAuthorizedUserInfo()

  if(userInfo == "unauthorized"){
    clientLogin(false)
  }
  else{
    clientLogin(true)
    userIdDiv.innerHTML = userInfo;
    await getDomainFindings();
    selectTextField()
  }

  getLastUsedUrl();

})();




async function setup(){

  let userInfo = await getAuthorizedUserInfo()

  if(userInfo == "unauthorized"){
    clientLogin(false)
  }
  else{
    userIdDiv.innerHTML = userInfo;
    selectTextField()
  }

}


function selectTextField() {

  const input = document.getElementById("contentField");
  
  console.log("FOCUSING...")
  
  input.focus();
  input.select();
}

async function getAuthorizedUserInfo(){
  const response = await fetch(server_url + "/getUser", {
    method: "GET", 
    cache: "no-cache", 
    mode: "cors",
    redirect: "follow", 
    referrerPolicy: "no-referrer",
    credentials: "include"
  });

  console.log("RETRIEVED USER INFO:")

  let resRaw = await response.text();
  console.log(response.status);

  if(response.status == 401){
    return "unauthorized";
  }

  return resRaw;
  
}

async function getDomainFindings(){

  const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});

  console.log(tab)
  let url = new URL(tab.url);
  let domain = url.hostname;

  let reqBody = {
    domain: domain
  }

  const response = await fetch(server_url + "/finding/user/domain/all", {
    method: "POST", 
    cache: "no-cache", 
    body: JSON.stringify(reqBody),
    mode: "cors",
    redirect: "follow", 
    referrerPolicy: "no-referrer",
    credentials: "include"
  });

  let jsonRes = await response.json()

  if(jsonRes != null){
    console.log(jsonRes)
    populateFindingDiv(jsonRes)
  }
  else{
    console.log("NO FINDINGS ON DOMAIn")
  }

}

async function populateFindingDiv(findingList){
  const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});
  currentUrl = tab.url

  let currentPageFinding = findingList.find(finding => finding.Link === tab.url)
  console.log("CURRENT PAGE FINDING")
  console.log(currentPageFinding)

  newFindingDiv.style = "display:none"
  findingsDiv.style = "display:block"

  findingsDiv.innerHTML = `
  <div>
    <h1>${currentPageFinding.Name}</h1>
    <p>${currentPageFinding.Content}</p>
  </div>`


}


async function cookieExists(){
  var cookieJar = await chrome.cookies.getAllCookieStores();
  console.log(cookieJar); // 'static' memory between function calls
  return true;
}

async function getLastUsedUrl(){
  const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});
  console.log("LAST USED URL:")
  console.log(tab.url);
  titleInputField.value = tab.title;
  urlInputField.value = tab.url;
}

async function createFinding(){
  
  
  let url = new URL(urlInputField.value);
  let domain = url.hostname;

  let reqBody = {
    name: titleInputField.value,
    link: urlInputField.value,
    content: contentInputField.value,
    domain: domain
  }


  const response = await fetch(server_url + "/finding/create", {
    method: "POST", 
    cache: "no-cache", 
    body: JSON.stringify(reqBody),
    mode: "cors",
    redirect: "follow", 
    referrerPolicy: "no-referrer",
    credentials: "include"
  });

  console.log(response)
}
