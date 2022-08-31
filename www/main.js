// Set-up code editor
var editor = ace.edit("editor");
// editor.setTheme("ace/theme/monokai");
editor.session.setMode("ace/mode/javascript");
// ---

document.addEventListener("DOMContentLoaded", async function (){
	var hk = await getHotkey()
	document.getElementById("run").innerText = "Run - " + hk;
	document.getElementById("stop").innerText = "Stop - " + hk;
	updateSavedList()
})

setInterval(async () => {
	let pos = await getMousePosition();
	mousePosition.innerText = "Mouse: (" + pos.x + ", " + pos.y + ")";
}, 1000)

async function save() {
	error.innerText = "";
	try {
		var code = editor.getValue();
		await saveCode(code);
	} catch (ex) {
		error.innerText = ex;
	}
}

async function run() {
	error.innerText = "";
	try {
		var code = editor.getValue();
		await executeCode(code);
	} catch (ex) {
		error.innerText = ex;
	}
}

async function stop() {
	error.innerText = "";
	try {
		await stopMacros();
	} catch (ex) {
		error.innerText = ex;
	}
}

// Called from Go
function log(msg) {
	logWindow.innerHTML += msg + "<br>";
	logWindow.scrollTop = logWindow.scrollHeight;
}

// Called from Go
function logImage(img) {
	logWindow.innerHTML += "<img src='data:image/jpeg;base64," + img + "'><br>";
	logWindow.scrollTop = logWindow.scrollHeight;
}

function help() {
	var w = window.open("/help", "popup");
}

function getCode() {
	return editor.getValue();
}

function onRun() {
    document.getElementById("run").style.display = 'none';
    document.getElementById("stop").style.display = 'unset';
}

function onStop() {
    document.getElementById("run").style.display = 'unset';
    document.getElementById("stop").style.display = 'none';
}

async function saveHandler() {
	var name = document.getElementById("s-name").value;
	var code = editor.getValue();
	await saveScript(name, code)
	await updateSavedList()
}

async function updateSavedList() {
	var list = document.getElementById("saved-list");
	list.innerHTML = '';
	var scripts = await getSavedScripts();
	for (var id in scripts)
		list.innerHTML += '<li>' +
			'<a class="dropdown-item d-flex justify-content-between" href="#" onclick="loadSavedScript(\''+id+'\')">'+
				id+
			    '<button class="btn-close" data-id="'+id+'"></button></a>' +
			'</li>';
    Array.from(list.getElementsByTagName("button")).forEach(btn => {
        btn.addEventListener('click', e => {
            e.stopPropagation();
            var id = e.target.getAttribute("data-id");
            if (!confirm("Delete " + id + "?"))
                return;
            (async function(){
                await deleteScript(id)
                await updateSavedList()
            })()
        })
    })
}

async function loadSavedScript(id) {
	var scripts = await getSavedScripts();
	if (scripts[id])
		editor.setValue(scripts[id], 1);
}
