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
            receiveAddr: ""
        }
    },
    methods: {
        getQueryVariablel(variable) {
            var query = window.location.search.substring(1);
            var vars = query.split("&");
            for (var i = 0; i < vars.length; i++) {
                var pair = vars[i].split("=");
                if (pair[0] == variable) {
                    return pair[1];
                }
            }
        }
    },
    mounted() {
        let orderString = localStorage.getItem("order");
        let order = JSON.parse(orderString);

        this.orderNumber = this.getQueryVariable("ordernum");
        this.amountIn = order.amount;
        this.currencyIn = order.in_currency;
        this.amountOut = this.getQueryVariablel("outamount");
        this.currencyOut = order.out_currency;
        this.receiveAddr = order.address;

    }
}
</script>

<template>
    <h2>Заявка номер {{ orderNumber }} создана.</h2>
    <p>Выполните перевод {{ amountIn }} {{ currencyIn }} на следующий адрес
        <b>{{ transferAddr }}</b> в течении следующих 15 минут, что бы получить
        {{ amountOut }} {{ currencyOut }} по следующему адресу {{ receiveAddr }}.
        Все заявки обрабатываются оператором в ручном режиме и занимают до 10 минут.
        Как только платеж будет выполнен нажмите кнопку платеж отправлен, приложив
        скриншот или чек перевода.
    </p>
    <button>Прикрепить скриншот</button>
</template>

<style></style>