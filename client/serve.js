const express = require('express');
const path = require('path');

const app = express()

app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, "/index.html"))
})

app.get('/index.js', (req, res) => {
    res.sendFile(path.join(__dirname, 'index.js'))
})

app.get("/pool", (req, res) => {
    res.sendFile(path.join(__dirname, "/pool.html"))
})

app.get('/pool.js', (req, res) => {
    res.sendFile(path.join(__dirname, 'pool.js'))
})

app.listen(1337)
console.log("Started listening on localhost:1337");
