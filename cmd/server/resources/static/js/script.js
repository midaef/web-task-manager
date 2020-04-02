let xmlHttp = new XMLHttpRequest()

window.onload = function() {
	let table = get()
	console.log(table)
	for (let i = 0; i < table.length; i++) {
		push(table[i].todo, table[i].datetime)
	}
}

function push(task, time) {
	let ul = document.getElementById("tasks")
	let li = document.createElement("li")
	li.appendChild(document.createTextNode(task + " " + time))
	ul.appendChild(li)
}

function add(task) {
	xmlHttp.open("POST", "/add", false)
	xmlHttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
	xmlHttp.send("task=" + task)
	return JSON.parse(xmlHttp.response)
}

function get() {
	xmlHttp.open("GET", "/get", false)
	xmlHttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
	xmlHttp.send()
	return JSON.parse(xmlHttp.response)
}

function onClick() {
	let taskName = document.getElementById("todo").value
	let time = add(taskName).datetime
	push(taskName, time)
}