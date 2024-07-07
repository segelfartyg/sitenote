var auth_url = 'https://accounts.google.com/o/oauth2/auth?';

var client_id = '1010731658636-aeejci8n3gctj78iqdehtti3qfqpn568.apps.googleusercontent.com'
var redirect_url = chrome.identity.getRedirectURL(); 
var server_url = 'http://localhost:9000'
var auth_params = {
    client_id: client_id,
    redirect_uri: redirect_url,
    response_type: 'id_token',
    scope: 'openid https://www.googleapis.com/auth/userinfo.email',
    login_hint: 'real_email@gmail.com' 
};

const url = new URLSearchParams(Object.entries(auth_params));
url.toString();
auth_url += url;

const button = document.getElementById("loginBtn")
const getUserBtn = document.getElementById("getUserBtn")

let userIdDiv = document.getElementById("userIdDiv");

button.addEventListener("click", (e) => {
    chrome.identity.launchWebAuthFlow({url: auth_url, interactive: true}, function(responseUrl) { 
        let idToken = responseUrl.substring(responseUrl.indexOf('id_token=') + 9);
        idToken = idToken.substring(0, idToken.indexOf('&'));
        console.log("ID TOKEN RETRIEVED:")
        console.log(idToken)
        login(idToken).then((res) => {

            if(res == "unauthorized"){
                clientLogin(false)
                userIdDiv.innerHTML = "Not logged in. Sign in to use NoteLad"
            }
            else{
                clientLogin(true)
                userIdDiv.innerHTML = res
            }

            
        });
    });
})

async function login(idToken){
let req = {id_token: idToken}
const response = await fetch(server_url + "/login", {
    method: "POST", 
    cache: "no-cache", 
    mode: "cors",
    headers: {},
    redirect: "follow", 
    referrerPolicy: "no-referrer", 
    credentials: "include",
    body: JSON.stringify(req), 
  });

  if(response.status == 401){
    return "unauthorized"
  }
  else{
    return response.text(); 
  }
  
}

function clientLogin(loggedIn){
    if(!loggedIn){
        console.log("NOT LOGGED IN.")
        loginDiv.style.display = "flex";
        loggedInDiv.style.display = "none";
    }
    else{
        console.log("LOGGED IN.")
        loginDiv.style.display = "none";
        loggedInDiv.style.display = "flex";
    }
}

// async function getUserId(){
//     const response = await fetch(server_url + "/getUser", {
//         method: "GET", 
//         cache: "no-cache", 
//         mode: "cors",
//         redirect: "follow", 
//         referrerPolicy: "no-referrer",
//         credentials: "include"
//       });
//       console.log("RETRIEVED USER INFO:")
//       console.log(response.text()); 
// }