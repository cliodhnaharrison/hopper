function onSignIn(googleUser) {
	profile = googleUser.getBasicProfile();
	console.log('Email: ' + profile.getEmail());
	console.log('Name: ' + profile.getName());
	var name = profile.getName();
	var splitname = name.split(" ");
	var uname = splitname[0] + splitname[1];
	document.cookie = "username=" + uname.toLowerCase() + ";";
}

function signOut() {
	var session = gapi.auth2.getAuthInstance();
	session.signOut().then(function () {
		console.log('Signed out user');
	});
}

function discard() {
	$.ajax({
		url: "/discard"
	}).done(function() {
		console.log("Container discarded");
	});
}
