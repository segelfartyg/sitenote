let urlInputField = window.document.getElementById("urlField");

(async () => {
    // see the note below on how to choose currentWindow or lastFocusedWindow
    const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});
    console.log(tab.url);

   
    urlInputField.value = tab.url;

    // ..........
  })();