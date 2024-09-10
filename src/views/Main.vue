<script>
export default {
    data() {
        return {
            currencies: [],
            currencyIn: "",
            amountIn: 0,
            currencyOut: "",
            amountOut: 0,
            email: "",
            address: "",
            notovermin: false,
            minvalue: 0
        }
    },
    async mounted() {
        let response = await fetch("http://localhost:8080/api/list-currencies", {
            method: "GET"
        });

        let data = await response.json();
        this.currencies = [];
        data.currencies.forEach((currency) => {
            this.currencies.push(currency);
        });

        this.currencyIn = this.getQueryVariable("currin");
        this.currencyOut = this.getQueryVariable("currout");

        if (!this.currencyIn) {
            this.currencyIn = "TON";
        }
        if (!this.currencyOut) {
            this.currencyOut = "RUB";
        }
    },
    async updated() {
        let response = await fetch(`http://localhost:8080/api/current-rate?currency_in=${this.currencyIn}&currency_out=${this.currencyOut}&amount=${this.amountIn}`, {
            method: "GET",
            headers: {}
        });

        if (response.ok) {
            let amount = await response.json();
            this.amountOut = amount.amount;
            this.notovermin = false;
            this.amount = 0;
            return;
        }

        if (response.status === 409) {
            let rsp = await response.text();
            this.notovermin = true;
            this.amount = 0;
            this.minvalue = Number(rsp.replace("not over minimum operation ", ""));
            return;
        }

    },
    methods: {
        createOrder() {
            console.log("create order triggered")
        },
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
        async handleSubmit() {
            let headersList = {
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "email": this.email,
                "in_currency": this.currencyIn,
                "out_currency": this.currencyOut,
                "amount": this.amountIn,
                "address": this.address
            });

            let response = await fetch("http://localhost:8080/api/create-order", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            let data = await response.text();
            console.log(data);

        }
    }
}
</script>

<template>
    <form @submit.prevent="handleSubmit">
        <label for="currin">Отдаете:</label>
        <select id="currin" name="currin" v-model="currencyIn">
            <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                {{ currency.description }}</option>
        </select>
        <input type="text" id="currin" name="currency-in" v-model="amountIn">

        <label for="currout">Получаете:</label>
        <select id="currout" name="currout" v-model="currencyOut">
            <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                {{ currency.description }}</option>
        </select>
        <input type="text" id="currout" name="currency-out" v-model="amountOut" readonly>

        <label for="email">Электронная почта:</label>
        <input type="email" id="email" name="email" v-model="email">

        <label for="address">Адрес получения (карта/кошелек):</label>
        <input type="text" id="address" name="address" v-model="address">

        <input type="submit" value="Создать заявку" @click="createOrder">
    </form>

    <div class="alert" v-if="notovermin">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Слишком маленькая сумма для совершения операции, минимум - {{ this.minvalue }}
    </div>
</template>

<style scoped>
input[type=text],
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
</style>