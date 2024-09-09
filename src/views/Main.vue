<script>
export default {
    data() {
        return {
            currencies: [],
            currencyIn: "",
            amountIn: 0,
            currencyOut: "",
            amountOut: 0
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
        updateOnChange() { 

        }
    }
}
</script>

<template>
    <form>
        <label for="currin">Отдаете</label>
        <select id="currin" name="currin" v-model="currencyIn">
            <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                {{ currency.description }}</option>
        </select>
        <input type="text" id="currin" name="currency-in" v-model="amountIn">

        <label for="currout">Получаете</label>
        <select id="currout" name="currout" v-model="currencyOut">
            <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                {{ currency.description }}</option>
        </select>
        <input type="text" id="currout" name="currency-out" v-model="amountOut" readonly>


        <input type="submit" value="Создать заявку" @click="createOrder">
    </form>
</template>

<style scoped>
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

div {
    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
}
</style>