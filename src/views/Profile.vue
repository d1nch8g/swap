<script>
export default {
    data() {
        return {
            orders: [],
        }
    },
    methods: {
        logout() {
            localStorage.setItem("token", "");
            window.location.href = '/';
        },
        addCurrency() { }
    },
    async mounted() {
        let token = localStorage.getItem("token");
        let headersList = {
            "Authorization": `Bearer ${token}`
        }

        let response = await fetch("http://localhost:8080/api/user/list-orders", {
            method: "GET",
            headers: headersList
        });

        let resp = await response.json();
        this.orders = resp.orders;

    }
}
</script>

<template>
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
        </tr>
    </table>

    <br>

    <button class="button" style="vertical-align:middle" @click="logout"><span>Выйти </span></button>
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

.button {
    display: inline-block;
    border-radius: 4px;
    background-color: #f4511e;
    border: none;
    color: #FFFFFF;
    text-align: center;
    font-size: 28px;
    padding: 20px;
    width: 200px;
    transition: all 0.5s;
    cursor: pointer;
    margin: 5px;
}

.button span {
    cursor: pointer;
    display: inline-block;
    position: relative;
    transition: 0.5s;
}

.button span:after {
    content: '\00bb';
    position: absolute;
    opacity: 0;
    top: 0;
    right: -20px;
    transition: 0.5s;
}

.button:hover span {
    padding-right: 25px;
}

.button:hover span:after {
    opacity: 1;
    right: 0;
}
</style>