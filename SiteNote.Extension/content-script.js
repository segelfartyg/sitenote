let urlInputField = window.document.getElementById("urlField");
let loginDiv = window.document.getElementById("loginDiv");
let loggedInDiv = window.document.getElementById("loggedInDiv");
let postFindingBtn = window.document.getElementById("postFindingBtn");
let contentInputField = window.document.getElementById("contentField");





(async () => {

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
  }

  getLastUsedUrl();

})();

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

async function cookieExists(){
  var cookieJar = await chrome.cookies.getAllCookieStores();
  console.log(cookieJar); // 'static' memory between function calls
  return true;
}

async function getLastUsedUrl(){
  const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});
  console.log("LAST USED URL:")
  console.log(tab.url);
  urlInputField.value = tab.url;
}

async function createFinding(){
  

  let reqBody = {
    name: urlInputField.value,
    link: urlInputField.value,
    content: contentInputField.value
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
