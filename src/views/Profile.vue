<script>
export default {
    data() {
        return {
            orders: [],
            email: ""
        }
    },
    methods: {
        logout() {
            localStorage.setItem("token", "");
            window.location.href = '/';
        },
        navigateToOrder(id) {
            this.$router.push(`/order?orderid=${id}`);
        }
    },
    async mounted() {
        let token = localStorage.getItem("token");
        let headersList = {
            "Authorization": `Bearer ${token}`
        }

        let response = await fetch("/api/user/list-orders", {
            method: "GET",
            headers: headersList
        });

        let resp = await response.json();
        this.orders = resp.orders;

        let selfInfoResponse = await fetch("/api/user/self-info", {
            method: "GET",
            headers: headersList
        });

        let selfInfo = await selfInfoResponse.json();
        this.email = selfInfo.email; ``
    }
}
</script>

<template>
    <title>Страница профиля</title>

    <div class="card">
        <div class="container">
            <h4><b>{{ email }}</b></h4>
            <p>Это ваш профиль, здесь будут отображены все ваши заявки.</p>
        </div>
    </div>

    <h2>Ваши заявки</h2>
    <table id="table">
        <tr>
            <th>Номер заявки</th>
            <th>Принимаемая валюта</th>
            <th>Количество отправки</th>
            <th>Адрес отправки</th>
            <th>Отправляемая валюта</th>
            <th>Количество получения</th>
            <th>Адрес получения</th>
            <th>Статус</th>
            <th>Страница заявки</th>
        </tr>
        <tr v-for="order in orders">
            <td>{{ order.id }}</td>
            <td>{{ order.in_currency }}</td>
            <td>{{ order.in_amount }}</td>
            <td>{{ order.out_addr }}</td>
            <td>{{ order.out_currency }}</td>
            <td>{{ order.out_amount }}</td>
            <td>{{ order.recv_addr }}</td>
            <td>{{ order.status }}</td>
            <td><button @click="navigateToOrder(order.id)">link</button></td>
        </tr>
    </table>

    <br>
    <button class="button" @click="logout"><span>Выйти</span></button>
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

.card {
    /* Add shadows to create the "card" effect */
    margin: 20px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    transition: 0.3s;
}

/* On mouse-over, add a deeper shadow */
.card:hover {
    box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

/* Add some padding inside the card container */
.container {
    padding: 2px 16px;
}

.button {
    background-color: orangered;
    /* Green */
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
}

table {
    table-layout: fixed;
}

td {
    overflow: hidden;
    text-overflow: ellipsis;
    word-wrap: break-word;
}

@media only screen and (max-width: 480px) {

    /* horizontal scrollbar for tables if mobile screen */
    table {
        overflow-x: auto;
        display: block;
    }
}
</style>