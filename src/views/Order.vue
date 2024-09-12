<script>
export default {
    data() {
        return {
            orderNum: 0,
            status: "",
            trailText: ""
        }
    },
    async mounted() {
        this.orderNum = this.getQueryVariable("orderid");
        let headersList = {
            "orderid": this.orderNum.toString(),
        }

        let response = await fetch("http://localhost:8080/api/order-status", {
            method: "GET",
            headers: headersList
        });

        if (response.ok) {
            let resp = await response.json();
            this.status = resp.status;

            if (this.status === "платеж пользователем осуществлен") {
                this.trailText = "Платеж и подтверждение будут проверены оператором и после средства будут направлены на ваш адрес. Когда выплата будет осуществлена заявка сменит статус на выполненную."
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
    }
}
</script>

<template>
    <h3>Заявка номер {{ orderNum }}.</h3>
    <p>Заявка находится в статусе: <b>{{ status }}</b></p>
    <p v-if="trailText">{{ trailText }}</p>
</template>