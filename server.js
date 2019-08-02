const http = require('http')

const server = http.createServer((req, res) => {
	console.log(req.url)
	console.log('received')
	res.end('hello')
})
server.listen(4000)
console.log('start success')
