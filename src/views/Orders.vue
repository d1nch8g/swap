<script>
export default {
    data() {
        return {
            orders: [],
            email: "",
        }
    },
    async mounted() {
        let token = localStorage.getItem("token");

        let headersList = {
            "Authorization": `Bearer ${token}`,
        }

        let response = await fetch("/api/operator/finished-orders", {
            method: "GET",
            headers: headersList
        });

        if (response.ok) {
            let resp = await response.json();
            this.orders = resp.orders;
        }
    },
    methods: {
        async orderSearch() {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
            }

            let response = await fetch(`/api/operator/order-search?email=${this.email}`, {
                method: "GET",
                headers: headersList
            });

            if (response.ok) {
                let resp = await response.json();
                this.orders = resp.orders;
                return;
            }

            if (response.status === 404) {
                this.orders = [];
                return;
            }
            
        }
    },
}
</script>

<template>
    <title>Заявки</title>

    <h3>Заявки</h3>

    <form @submit.prevent="orderSearch">
        <label for="email">Пользовательский email:</label>
        <input type="email" id="email" name="email" v-model="email">

        <input type="submit" value="Найти">
    </form>

    <h2>Результаты:</h2>
    <table id="table">
        <tr>
            <th>ID</th>
            <th>Email</th>
            <th>Получаемая валюта</th>
            <th>Получаемое количество</th>
            <th>Отдаваемая валюта</th>
            <th>Отдаваемое количество</th>
            <th>Отправляемый адресс</th>
            <th>Подтверждение платежа</th>
            <th>Статус заявки</th>
        </tr>
        <tr v-for="order in orders">
            <td>{{ order.id }}</td>
            <td>{{ order.email }}</td>
            <td>{{ order.currin }}</td>
            <td>{{ order.amountin }}</td>
            <td>{{ order.currout }}</td>
            <td>{{ order.amountout }}</td>
            <td>{{ order.address }}</td>
            <td><img v-bind:src="'data:image/jpeg;base64,' + order.approvepic" /></td>
            <td>{{ order.status }}</td>
        </tr>
    </table>

</template>

<style>
img {
    max-height: 150px;
}

input[type=email],
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