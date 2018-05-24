const https = require('https');


// !!!PUT YOUR API KEY HERE!!!
var api_key="YOUR_API_KEY";


if (process.argv.length != 3) {
  console.log("usage: node hubuco_email_verifier.js <email>");
  return;
}

var email = process.argv[2];

var url = "https://api.hubuco.com/api/v3/?api="+api_key+"&"+"email="+email;
https.get(url, function(res){
    var body = '';
    res.on('data', function(chunk){
        body += chunk;
    });
    res.on('end', function(){
        var res = JSON.parse(body);
    switch (res.resultcode) {
    case 1:
        console.log("Ok");
        break;
    case 2:
        console.log("Catch All");
        break;
    case 3:
        console.log("Unknown");
        break;
    case 4:
        console.log("Error: " + res.error);
        break;
    case 5:
        console.log("Disposable");
        break;
    case 6:
        console.log("Invlaid");
        break;
    }
    });
}).on('error', function(e){
      console.log("Got an error: ", e);
});

