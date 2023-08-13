const socket = new WebSocket("ws://localhost:8080/chat");
const msgInput = document.getElementById("msg-input");
const sendBtn = document.getElementById("send-btn");
const msgContainer = document.getElementById("msg-container");
socket.addEventListener("open", function () {
    console.log("Connected to server.");
});
socket.addEventListener("message", function (e) {
    const text = e.data;
    displayMsg(text);
});
sendBtn.addEventListener("click", function (e) {
    const text = msgInput.value;
    if (text.trim() !== "") {
        sendMsg(text);
        msgInput.value = "";
    }
});
function sendMsg(text) {
    const msg = {text:text};
    socket.send(JSON.stringify(msg));
}
function displayMsg(msg) {
    const msgEl = document.createElement("p");
    const data = JSON.parse(msg);
    msgEl.textContent = data.text;
    msgContainer.appendChild(msgEl);
}
