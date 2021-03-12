function onSignIn(googleUser) {
	profile = googleUser.getBasicProfile();
	console.log('Email: ' + profile.getEmail());
	console.log('Name: ' + profile.getName());
}

function signOut() {
	var session = gapi.auth2.getAuthInstance();
	session.signOut().then(function () {
		console.log('Signed out user.');
	});
}
