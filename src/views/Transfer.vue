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
            incorrect: false,
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
        async fileChosen() {
            var input = document.querySelector('input[type="file"]');

            var data = new FormData();
            data.append('file', input.files[0]);

            let response = await fetch(`http://localhost:8080/api/confirm-payment?order_id=${this.orderNumber}`, {
                method: "POST",
                body: data,
                headers: {}
            });

            if (response.ok) {
                this.$router.push(`/order?orderid=${this.orderNumber}`);
                return;
            }
            this.incorrect = true;
        }
    }
}
</script>

<template>
    <div class="alert" v-if="incorrect">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Ошибка при загрузке файла, попробуйте еще раз
    </div>
    <h2>Заявка номер {{ orderNumber }} создана.</h2>
    <p>Выполните перевод {{ amountIn }} {{ currencyIn }} на следующий адрес
        <b>{{ transferAddr }}</b> в течении следующих 15 минут, что бы получить
        {{ amountOut }} {{ currencyOut }} по следующему адресу {{ receiveAddr }}.
        Как только платеж будет выполнен - приложите чек перевода. Все
        заявки обрабатываются оператором в ручном режиме и занимают до 10 минут.
        Как только файл будет прикреплен нажмите кнопку отправить данные.
    </p>
    <form>
        <label for="file-upload" class="custom-file-upload">
            Прикрепить скриншот:
        </label>
        <br>
        <input id="file-upload" type="file">
    </form>
    <button @click="fileChosen">Отправить данные</button>
</template>

<style>
.button {
    background-color: grey;
    /* Green */
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
}
</style>