<script>
export default {
    data() {
        return {
            cardnum: "",
            incorrect: false
        }
    },
    methods: {
        async fileChosen() {
            let orderstr = localStorage.getItem("order");
            let order = JSON.parse(orderstr);

            var input = document.querySelector('input[type="file"]');

            var data = new FormData();
            data.append('file', input.files[0]);

            let response = await fetch(`/api/validate-card?email=${order.email}&currency=${order.in_currency}&addr=${this.cardnum}`, {
                method: "POST",
                body: data,
                headers: {}
            });

            if (response.ok) {
                let headersList = {
                    "Content-Type": "application/json"
                }

                let response = await fetch("/api/create-order", {
                    method: "POST",
                    body: orderstr,
                    headers: headersList
                });

                if (response.ok) {
                    let resp = await response.json();
                    localStorage.setItem("lastorderid", resp.order_number.toString());
                    localStorage.setItem("transfer_address", resp.transfer_address);
                    localStorage.setItem("out_amount", resp.out_amount.toString());
                    this.$router.push(`/transfer/?addr=${resp.transfer_address}&inamount=${resp.in_amount}&outamount=${resp.out_amount}&ordernum=${resp.order_number}`);
                    return;
                }
            }
            this.incorrect = true;
        }
    }
}
</script>

<template>
    <div class="alert" v-if="incorrect">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Произошла ошибка
    </div>
    <h3>Подтверждение карты.</h3>
    <p>Напишите номер карты с которой будет осуществлен перевод и прикрепите
        фото карты на фоне сайта или листа бумаги с названием сайта. Нам нужно
        подтверждение что вы являетесь владельцем карты.
    </p>
    <form>
        <label for="card">Введите номер карты:</label>
        <input type="text" id="card" name="card" v-model="cardnum">

        <label for="file-upload" class="custom-file-upload">
            Прикрепить скриншот:
        </label>
        <br>
        <input id="file-upload" type="file">
    </form>
    <input type="submit" @click="fileChosen">
</template>

<style scoped>
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