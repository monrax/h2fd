const express = require('express')
const app = express()
app.use(express.static('app'))
app.listen(8000, () => console.log('Serving at http://localhost:8000'))