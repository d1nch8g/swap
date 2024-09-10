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
            minAmountToExchnage: 0,
            showCreateOrderButton: true,
            showMinAmountNotification: false,
            showPairNotSupportedNotfication: false,
            allOperatorsAreBusyNotification: false
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
            this.currencyOut = "SBPRUB";
        }
    },
    async updated() {
        let response = await fetch(`http://localhost:8080/api/current-rate?currency_in=${this.currencyIn}&currency_out=${this.currencyOut}&amount=${this.amountIn}`, {
            method: "GET",
            headers: {}
        });

        // handle normal response
        if (response.ok) {
            let amount = await response.json();
            this.amountOut = amount.amount;
            this.showCreateOrderButton = true;
            this.showMinAmountNotification = false;
            this.showPairNotSupportedNotfication = false;
            return;
        }

        // handle not over required minimum error
        if (response.status === 409) {
            let rsp = await response.text();
            this.showCreateOrderButton = false;
            this.showMinAmountNotification = true;
            this.showPairNotSupportedNotfication = false;
            this.amountOut = 0;
            this.minAmountToExchnage = Number(rsp.replace("not over minimum operation ", ""));
            return;
        }

        // handle exchangers pair not supported
        if (response.status === 403) {
            this.showCreateOrderButton = false;
            this.showMinAmountNotification = false;
            this.showPairNotSupportedNotfication = true;
            this.amountOut = 0;
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
        async createOrder() {
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

            localStorage.setItem("order", bodyContent);

            let response = await fetch("http://localhost:8080/api/create-order", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                let resp = await response.json();
                this.$router.push(`/transfer/?addr=${resp.transfer_address}&inamount=${resp.in_amount}&outamount=${resp.out_amount}&ordernum=${resp.order_number}`);
                return;
            }

            if (response.status === 409) {
                this.showCreateOrderButton = false;
                this.showMinAmountNotification = false;
                this.showPairNotSupportedNotfication = false;
                this.allOperatorsAreBusyNotification = true;
                return;
            }

            if (response.status === 403) {
                let resp = await response.json();

                this.$router.push(`/validate-card`);
                // handle validate card operation
            }
        }
    }
}
</script>

<template>
    <form @submit.prevent="createOrder">
        <label for="currin">Отдаете:</label>
        <select id="currin" name="currin" v-model="currencyIn">
            <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                {{ currency.description }}</option>
        </select>
        <input type="number" id="currin" name="currency-in" v-model="amountIn">

        <label for="currout">Получаете:</label>
        <select id="currout" name="currout" v-model="currencyOut">
            <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                {{ currency.description }}</option>
        </select>
        <input type="number" id="currout" name="currency-out" v-model="amountOut" readonly>

        <label for="email">Электронная почта:</label>
        <input type="email" id="email" name="email" v-model="email">

        <label for="address">Адрес получения (номер карты/кошелек):</label>
        <input type="text" id="address" name="address" v-model="address">

        <input type="submit" value="Создать заявку" @click="createOrder" v-if="showCreateOrderButton">
    </form>

    <div class="alert" v-if="showMinAmountNotification">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Слишком маленькая сумма для совершения операции, минимум - {{ this.minAmountToExchnage }}
    </div>

    <div class="alert" v-if="showPairNotSupportedNotfication">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Выбранная валютная пара не поддерживается обменником
    </div>

    <div class="alert" v-if="allOperatorsAreBusyNotification">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Все операторы в данный момент заняты, создайсте заявку позже.
    </div>
</template>

<style scoped>
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
</style>