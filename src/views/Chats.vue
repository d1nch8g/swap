<script>
export default {
    data() {
        return {
            chats: [],
        }
    },
    async mounted() {
        let token = localStorage.getItem("token");

        let headersList = {
            "Authorization": `Bearer ${token}`
        }

        let response = await fetch("/api/operator/unresolved-chats", {
            method: "GET",
            headers: headersList
        });

        if (response.ok) {
            let resp = await response.json();
            this.chats = resp.chats;
        }
    },
    methods: {
        async resolveChat(uuid) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "uuid": uuid
            });

            let response = await fetch("/api/operator/resolve-chat", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/chats";
            }
        }
    }
}
</script>

<template>
    <title>Неразрешенные чаты</title>

    <h2>Неразрешенные чаты</h2>
    <table id="table">
        <tr>
            <th>ID</th>
            <th>UUID</th>
            <th>Ссылка на чат</th>
            <th>Закрыть чат</th>
        </tr>
        <tr v-for="chat in chats">
            <td>{{ chat.id }}</td>
            <td>{{ chat.uuid }}</td>
            <td><a :href="'/chat?uuid=' + chat.uuid">Ссылка</a></td>
            <td><button @click="resolveChat(chat.uuid)">Закрыть</button></td>
        </tr>
    </table>
</template>


<style scoped>
#table {
    font-family: Arial, Helvetica, sans-serif;
    border-collapse: collapse;
    width: 100%;
}

#table td,
#table th {
    border: 1px solid #ddd;
    padding: 8px;
}

#table tr:nth-child(even) {
    background-color: #f2f2f2;
}

#table tr:hover {
    background-color: #ddd;
}

#table th {
    padding-top: 12px;
    padding-bottom: 12px;
    text-align: left;
    background-color: lightslategray;
    color: white;
}
</style>