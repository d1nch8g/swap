<script>
import ValidateCard from './ValidateCard.vue';

export default {
    data() {
        return {
            currencies: [],
            balances: [],
            orders: [],
            cardConfirmations: [],
            finorders: [],
            currencyCode: "",
            address: "",
            balance: 0
        }
    },
    async mounted() {
        let currResponse = await fetch("/api/list-currencies", {
            method: "GET"
        });

        let data = await currResponse.json();
        this.currencies = data.currencies;

        let token = localStorage.getItem("token");

        let headersList = {
            "Authorization": `Bearer ${token}`,
        }

        let response = await fetch("/api/operator/list-balances", {
            method: "GET",
            headers: headersList
        });

        if (response.ok) {
            let data = await response.json();
            this.balances = data.balances;
        }

        let ordersResponse = await fetch("/api/operator/get-orders", {
            method: "GET",
            headers: headersList
        });

        if (ordersResponse.ok) {
            let resp = await ordersResponse.json();
            this.orders = resp.orders;
        }

        let cardConfirmations = await fetch("/api/operator/get-card-confirmations", {
            method: "GET",
            headers: headersList
        });

        if (cardConfirmations.ok) {
            let resp = await cardConfirmations.json();
            this.cardConfirmations = resp.card_confirmations;
        }

        let finordersResponse = await fetch("/api/operator/finished-orders", {
            method: "GET",
            headers: headersList
        });

        if (finordersResponse.ok) {
            let resp = await finordersResponse.json();
            this.finorders = resp.orders;
        }

    },
    methods: {
        async removeBalance(id) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "id": id
            });

            let response = await fetch("/api/operator/remove-balance", {
                method: "DELETE",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/operator";
            }
        },
        openForm() {
            document.getElementById("myForm").style.display = "block";
        },
        closeForm() {
            document.getElementById("myForm").style.display = "none";
        },
        async updateBalance() {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "balance_id": 0,
                "currency_code": this.currencyCode,
                "balance": this.balance,
                "address": this.address
            });

            let response = await fetch("/api/operator/update-balance", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/operator";
            }

        },
        async closeOrder(id) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "order_id": id
            });

            let response = await fetch("/api/operator/execute-order", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/operator"
            }

        },
        async cancelOrder(id) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "order_id": id,
            });

            let response = await fetch("/api/operator/cancel-order", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/operator";
            }
        },
        async approveCard(id) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "confirmation_id": id,
            });

            let response = await fetch("/api/operator/approve-card", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/operator";
            }
        },
        async cancelCard(id) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "confirmation_id": id,
            });

            let response = await fetch("/api/operator/cancel-card", {
                method: "DELETE",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/operator";
            }
        }
    }
}
</script>

<template>
    <h2>Активные заявки</h2>
    <table id="table">
        <tr>
            <th>ID</th>
            <th>Email</th>
            <th>Получаемая валюта</th>
            <th>Получаемое количество</th>
            <th>Отдаваемая валюта</th>
            <th>Отдаваемое количество</th>
            <th>Отправляемый адресс</th>
            <th>Картинка подтверждение платежа</th>
            <th>Подтвердить перечисление</th>
            <th>Отменить заявку</th>
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
            <td><button @click="closeOrder(order.id)">Подтвердить и закрыть</button></td>
            <td><button @click="cancelOrder(order.id)">Отменить</button></td>
        </tr>
    </table>

    <h2>Подтверждения карт</h2>
    <table id="table">
        <tr>
            <th>Адрес</th>
            <th>Картинка</th>
            <th>Подтвердить</th>
            <th>Удалить</th>
        </tr>
        <tr v-for="cardConfirmation in cardConfirmations">
            <td>{{ cardConfirmation.address }}</td>
            <td><img v-bind:src="'data:image/jpeg;base64,' + cardConfirmation.image" /></td>
            <td><button @click="approveCard(cardConfirmation.id)">Подтвердить</button></td>
            <td><button @click="cancelCard(cardConfirmation.id)">Удалить</button></td>
        </tr>
    </table>

    <h2>Балансы</h2>
    <table id="table">
        <tr>
            <th>ID баланса</th>
            <th>Описание</th>
            <th>Код</th>
            <th>Адрес</th>
            <th>Баланс</th>
            <th>Удалить</th>
        </tr>
        <tr v-for="balance in balances">
            <td>{{ balance.id }}</td>
            <td>{{ balance.description }}</td>
            <td>{{ balance.code }}</td>
            <td>{{ balance.address }}</td>
            <td>{{ balance.balance }}</td>
            <td><button @click="removeBalance(balance.id)">Удалить</button></td>
        </tr>
    </table>

    <h2>Выполненные заявки</h2>
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
        <tr v-for="order in finorders">
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

    <button class="open-button" @click="openForm">Создать баланс</button>
    <!-- The form -->
    <div class="form-popup" id="myForm">
        <form class="form-container">
            <h1>Баланс</h1>

            <label for="curr"><b>Код валюты:</b></label>
            <select id="curr" name="curr" v-model="currencyCode">
                <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                    {{ currency.description }}</option>
            </select>

            <label for="address"><b>Адрес:</b></label>
            <input type="text" placeholder="Введи ваш адрес (кошелек/номер карты)" name="address" required
                v-model="address">

            <label for="balance"><b>Баланс:</b></label>
            <input type="number" placeholder="Введи значение баланса" name="balance" required v-model="balance">

            <button type="button" class="btn" @click="updateBalance">Создать или обновить баланс</button>
            <button type="button" class="btn cancel" @click="closeForm">Закрыть редактор</button>
        </form>
    </div>
</template>

<style scoped>
img {
    max-height: 150px;
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

/* Button used to open the contact form - fixed at the bottom of the page */
.open-button {
    background-color: #555;
    color: white;
    padding: 16px 20px;
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
    padding: 16px 20px;
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

input[type=text],
input[type=email],
input[type=number],
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

div {
    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
}

/* The alert message box */
.alert {
    padding: 20px;
    background-color: #f44336;
    /* Red */
    color: white;
    margin-bottom: 15px;
}


/* The alert message box */
.alert-green {
    padding: 20px;
    background-color: green;
    /* Red */
    color: white;
    margin-bottom: 15px;
}

/* The close button */
.closebtn {
    margin-left: 15px;
    color: white;
    font-weight: bold;
    float: right;
    font-size: 22px;
    line-height: 20px;
    cursor: pointer;
    transition: 0.3s;
}

/* When moving the mouse over the close button */
.closebtn:hover {
    color: black;
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