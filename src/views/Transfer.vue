<script>
export default {
    data() {
        return {
            orderNumber: 0,
            transferAddr: "",
            amountIn: 0,
            currencyIn: "",
            amountOut: 0,
            currencyOut: "",
            receiveAddr: "",
            fileUpload: null,
        }
    },
    mounted() {
        let orderString = localStorage.getItem("order");
        let order = JSON.parse(orderString);

        this.orderNumber = this.getQueryVariable("ordernum");
        this.transferAddr = this.getQueryVariable("addr");
        this.amountIn = order.amount;
        this.currencyIn = order.in_currency;
        this.amountOut = this.getQueryVariable("outamount");
        this.currencyOut = order.out_currency;
        this.receiveAddr = order.address;

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
        fileChosen(f) {
            // let file = this.$refs.file.files[0];
            console.log(f);
            console.log(f.target.files[0]);
        }
    }
}
</script>

<template>
    <h2>Заявка номер {{ orderNumber }} создана.</h2>
    <p>Выполните перевод {{ amountIn }} {{ currencyIn }} на следующий адрес
        <b>{{ transferAddr }}</b> в течении следующих 15 минут, что бы получить
        {{ amountOut }} {{ currencyOut }} по следующему адресу {{ receiveAddr }}.
        Как только платеж будет выполнен нажмите приложите чек перевода. Все
        заявки обрабатываются оператором в ручном режиме и занимают до 10 минут.
    </p>
    <form>
        <label for="file-upload" class="custom-file-upload">
            Прикрепить скриншот:
        </label>
        <br>
        <input id="file-upload" type="file" v-on:change="fileChosen">
    </form>
</template>

<style></style>