<script>
export default {
    data() {
        return {
            uuid: "",
            chat: [],
            chatMessage: "",
        }
    },
    async mounted() {
        this.uuid = this.getQueryVariable("uuid");


        setInterval(async () => {
            let response = await fetch(`/api/get-chat-messages/${this.uuid}`, {
                method: "GET",
                headers: {}
            });

            if (response.ok) {
                let resp = await response.json();
                this.chat = resp.messages;
            }
        }, 2000);
    },
    methods: {
        getQueryVariable(variable) {
            var query = window.location.search.substring(1);
            var vars = query.split("&");
            for (var i = 0; i < vars.length; i++) {
                var pair = vars[i].split("=");
                if (pair[0] == variable) {
                    return pair[1];
                }
            }
        },
        async sendMessage() {
            let headersList = {
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "uuid": this.uuid,
                "message": this.chatMessage,
                "outgoing": false,
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
        }
    }
}
</script>

<template>

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

    <form @submit.prevent="sendMessage">
        <input type="text" id="chatmes" name="chatmes" v-model="chatMessage">
        <input type="submit">
    </form>

</template>

<style scoped>
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
    height: 500px;
    overflow-y: scroll;
}


input[type=text],
select {
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
}

input[type=submit] {
    width: 100%;
    background-color: #4CAF50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

input[type=submit]:hover {
    background-color: #45a049;
}
</style>