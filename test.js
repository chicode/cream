var Docker = require('dockerode');
var docker = new Docker({
  socketPath: '/var/run/docker.sock'
});

function runExec(container, command) {

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
		chunks.push(chunk.toString())
	    })
	    
	    stream.on('end', function() {
		console.log(chunks.join(''))
	    })
	});
    });
}

docker.createContainer({
    Image: 'ubuntu',
    Tty: true,
    Cmd: ['/bin/bash']
}, function(err, container) {
    container.start({}, function(err, data) {
	runExec(container);
    });
});
