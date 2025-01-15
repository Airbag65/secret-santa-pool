const express = require('express');
const path = require('path');

const app = express()

app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, "/index.html"))
})

app.listen(1337)
console.log("Started listening on localhost:1337");
