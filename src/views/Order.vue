<script>
export default {
    data() {
        return {
            orderNum: 0,
            status: "",
            trailText: "",
            showPayment: false,
            order: null,
            transferAddr: "",
            amountOut: ""
        }
    },
    async mounted() {
        this.orderNum = localStorage.getItem("lastorderid");

        let queryNum = this.getQueryVariable("orderid");
        if (queryNum) {
            this.orderNum = queryNum;
        }
        let headersList = {
            "orderid": this.orderNum.toString(),
        }

        let response = await fetch("/api/order-status", {
            method: "GET",
            headers: headersList
        });

        if (response.ok) {
            let resp = await response.json();
            this.status = resp.status;

            if (this.status === "платеж пользователем осуществлен") {
                this.trailText = "Платеж и подтверждение будут проверены оператором и после средства будут направлены на ваш адрес. Когда выплата будет осуществлена заявка сменит статус на выполненную.";
            }

            if (this.status === "ожидает платежа") {
                this.showPayment = true;
                let order = localStorage.getItem("order")
                this.order = JSON.parse(order);
                this.transferAddr = localStorage.getItem("transfer_address")
                this.amountOut = localStorage.getItem("out_amount")
            }
        }
        //handle unknown errors
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

            let response = await fetch(`/api/confirm-payment?order_id=${this.orderNum}`, {
                method: "POST",
                body: data,
                headers: {}
            });

            if (response.ok) {
                window.location.href = `/order`
                this.$router.push(`/order?orderid=${this.orderNum}`);
                return;
            }
            this.incorrect = true;
        }
    }
}
</script>

<template>
    <h3>Заявка номер {{ orderNum }}.</h3>
    <p>Заявка находится в статусе: <b>{{ status }}</b></p>
    <p v-if="trailText">{{ trailText }}</p>
    <div v-if="showPayment">
        <p>Выполните перевод {{ order.amount }} {{ order.in_currency }} на следующий адрес
            <b>{{ transferAddr }}</b> в течении следующих 15 минут, что бы получить
            {{ amountOut }} {{ order.out_currency }} по следующему адресу {{ order.address }}.
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
    </div>
</template>