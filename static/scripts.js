function onSignIn(googleUser) {
	profile = googleUser.getBasicProfile();
	console.log('Email: ' + profile.getEmail());
	console.log('Name: ' + profile.getName());
}

window.addEventListener("load", function() {
	function send() {
		const XHR = new XMLHttpRequest();

		const FD = new FormData(form);
		XHR.addEventListener("load", function(event) {
			alert(event.target.responseText);
		});

		XHR.open("GET", "");
		var name = profile.getName();
		var splitname = name.split(" ");
		var username= splitname.join("");
		var dataAsString = new URLSearchParams(FormData).toString();
		var data = dataAsString + "&user=" + username;
		console.log(data);
		XHR.send(data);
	}
	const form = document.getElementById("images");
	form.addEventListener("submit", function(event) {
		event.preventDefault();
		send();
	});
});
