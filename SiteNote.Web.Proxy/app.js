console.log("STARTING WEB PROXY");


const express = require('express');
const path = require('path');
const app = express();

app.use(express.static(path.join('./SiteNote.Web/dist')));

// Start the server
const PORT = process.env.PORT || 5173;
app.listen(PORT, () => {
    console.log(`App listening on port ${PORT}`);
    console.log('Press Ctrl+C to quit.');
});


// For any other route, serve the index.html file.
app.get("*", (req, res) => {
    res.sendFile(path.join(__dirname, "./SiteNote.Web/dist", "index.html"));
  });