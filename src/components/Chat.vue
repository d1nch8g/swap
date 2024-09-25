<script>
export default {
    data() {
        return {
            chatMessage: "",
            chat: [],
            uuid: "",
        }
    },
    methods: {
        openForm() {
            document.getElementById("myForm").style.display = "block";
        },
        closeForm() {
            document.getElementById("myForm").style.display = "none";
        },
        async sendMessage() {
            let headersList = {
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "uuid": this.uuid,
                "message": this.chatMessage,
                "outgoing": true,
            });

            let response = await fetch("/api/send-chat-message", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                let response = await fetch(`/api/get-chat-messages/${this.uuid}`, {
                    method: "GET",
                    headers: {}
                });

                if (response.ok) {
                    let resp = await response.json();
                    this.chat = resp.messages;
                }
            }

            this.chatMessage = "";
        },
        uuidv4() {
            return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
                .replace(/[xy]/g, function (c) {
                    const r = Math.random() * 16 | 0,
                        v = c == 'x' ? r : (r & 0x3 | 0x8);
                    return v.toString(16);
                });
        },
    },
    async mounted() {
        let uuid = localStorage.getItem("uuid");

        if (!uuid) {
            uuid = this.uuidv4();
            localStorage.setItem("uuid", uuid);
        }

        this.uuid = uuid;

        setInterval(async () => {
            let response = await fetch(`/api/get-chat-messages/${uuid}`, {
                method: "GET",
                headers: {}
            });

            if (response.ok) {
                let resp = await response.json();
                if (resp.messages.length > this.chat.length) {
                    // later add play sound
                }
                this.chat = resp.messages;
            }
        }, 2000);
    }
}
</script>

<template>
    <!-- A button to open the popup form -->
    <button class="open-button" @click="openForm">Открыть чат</button>

    <!-- The form -->
    <div class="form-popup" id="myForm">
        <form class="form-container" @submit.prevent="sendMessage">
            <h3>Онлайн чат с платформой</h3>

            <div class="chat">
                <div class="container">
                    <p>Здравствуйте!</p>
                </div>

                <div class="container">
                    <p>Это онлайн поддержка, чем мы можем вам помочь?</p>
                </div>

                <div v-for="message in chat">
                    <div :class="(message.outgoing) ? 'container darker' : 'container'">
                        <p>{{ message.message }}</p>
                        <span :class="(message.outgoing) ? 'time-left' : 'time-right'">
                            {{ message.time }}
                        </span>
                    </div>
                </div>
            </div>

            <input type="text" v-model="chatMessage">

            <button type="submit" class="btn">Отправить сообщение</button>
            <button type="button" class="btn cancel" @click="closeForm">Закрыть чат</button>
        </form>
    </div>
</template>

<style scoped>
* {
    box-sizing: border-box;
}

/* Button used to open the contact form - fixed at the bottom of the page */
.open-button {
    background-color: #555;
    color: white;
    padding: 16px 5px;
    border: none;
    cursor: pointer;
    opacity: 0.8;
    position: fixed;
    bottom: 23px;
    right: 28px;
    width: 280px;
}

/* The popup form - hidden by default */
.form-popup {
    display: none;
    position: fixed;
    bottom: 0;
    right: 15px;
    border: 3px solid #f1f1f1;
    z-index: 9;
}

/* Add styles to the form container */
.form-container {
    max-width: 300px;
    padding: 10px;
    background-color: white;
}

/* Full-width input fields */
.form-container input[type=text],
.form-container input[type=password] {
    width: 100%;
    padding: 15px;
    margin: 5px 0 22px 0;
    border: none;
    background: #f1f1f1;
}

/* When the inputs get focus, do something */
.form-container input[type=text]:focus,
.form-container input[type=password]:focus {
    background-color: #ddd;
    outline: none;
}

/* Set a style for the submit/login button */
.form-container .btn {
    background-color: #04AA6D;
    color: white;
    padding: 16px 5px;
    border: none;
    cursor: pointer;
    width: 100%;
    margin-bottom: 10px;
    opacity: 0.8;
}

/* Add a red background color to the cancel button */
.form-container .cancel {
    background-color: red;
}

/* Add some hover effects to buttons */
.form-container .btn:hover,
.open-button:hover {
    opacity: 1;
}

body {
    margin: 0 auto;
    max-width: 800px;
    padding: 0 5px;
}

.container {
    border: 2px solid #dedede;
    background-color: #f1f1f1;
    border-radius: 5px;
    padding: 10px;
    margin: 10px 0;
}

.darker {
    border-color: #ccc;
    background-color: #ddd;
}

.container::after {
    content: "";
    clear: both;
    display: table;
}

.time-right {
    float: right;
    color: #aaa;
}

.time-left {
    float: left;
    color: #999;
}

.chat {
    height: 300px;
    overflow-y: scroll;
}
</style>