'use strict';

var Docker = require('dockerode')
var docker = new Docker({socketPath: '/var/run/docker.sock'})

var app = require('express')()
var http = require('http').Server(app)
var io = require('socket.io')(http)

var Docker = require('dockerode');
var docker = new Docker({
    socketPath: '/var/run/docker.sock'
});

function runExec(container, command, callback) {

    var options = {
	Cmd: ['python', '-c', command],
	Env: ['VAR=ttslkfjsdalkfj'],
	AttachStdout: true,
	AttachStderr: true
    };
    
    container.exec(options, function(err, exec) {
	if (err) return;
	exec.start(function(err, stream) {
	    if (err) return;
	    
	    var chunks = []
	    stream.on('data', function(chunk) {
		chunks.push(chunk.toString().replace(/[^ -~]+/g, ""))
	    })
	    
	    stream.on('end', function() {
		console.log(chunks)
		callback(chunks[chunks.length - 1])
	    })
	});
    });
}

io.on('connection', function(socket) {
    socket.on('run', function(code) {
	docker.createContainer({
	    Image: 'python',
	    Tty: true,
	    Cmd: ['/bin/bash']
	}, function(err, container) {
	    container.start({}, function(err, data) {
		runExec(container, code, (m) => socket.emit('result', m));
	    });
	});
	
    })
})

http.listen(3000, function() {
	console.log('listening on 3000')
})
