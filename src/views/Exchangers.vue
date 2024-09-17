<script>
export default {
    data() {
        return {
            currencies: [],
            exchangers: [],
            inCurr: "",
            outCurr: "",
            description: "",
            inmin: 0,
            paymentVerify: false
        }
    },
    async mounted() {
        let response = await fetch("/api/list-currencies", {
            method: "GET"
        });

        let data = await response.json();
        this.currencies = data.currencies;

        let exchResp = await fetch("/api/list-exchangers", {
            method: "GET"
        });

        let exchData = await exchResp.json();
        this.exchangers = exchData.exchangers;
    },
    methods: {
        openForm() {
            document.getElementById("myForm").style.display = "block";
        },
        closeForm() {
            document.getElementById("myForm").style.display = "none";
        },
        async removeExchanger(id) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "id": id
            });

            let response = await fetch("/api/admin/remove-exchanger", {
                method: "DELETE",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/exchangers";
            }
        },
        async createExchanger() {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "description": this.description,
                "inmin": this.inmin,
                "payment_verification": this.paymentVerify,
                "in_currency": this.inCurr,
                "out_currency": this.outCurr
            });

            let response = await fetch("/api/admin/create-exchanger", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/exchangers"
            }
        }
    }
}
</script>

<template>
    <title>Обменники</title>

    <h2>Обменники</h2>
    <table id="table">
        <tr>
            <th>ID</th>
            <th>Минимум для создания заявки</th>
            <th>Описание</th>
            <th>Требуется ли подтверждение адреса</th>
            <th>Входящая валюта</th>
            <th>Выходящая валюта</th>
            <th>Удалить обменник</th>
        </tr>
        <tr v-for="exchanger in exchangers">
            <td>{{ exchanger.id }}</td>
            <td>{{ exchanger.inmin }}</td>
            <td>{{ exchanger.description }}</td>
            <td>{{ exchanger.require_payment_verification }}</td>
            <td>{{ exchanger.in_currency }}</td>
            <td>{{ exchanger.out_currency }}</td>
            <td><button @click="removeExchanger(exchanger.id)">Удалить</button></td>
        </tr>
    </table>

    <button class="open-button" @click="openForm">Добавить обменник</button>
    <!-- The form -->
    <div class="form-popup" id="myForm">
        <form class="form-container">
            <h3>Добавить обменник</h3>

            <label for="curr"><b>Входящая валюта:</b></label>
            <select id="curr" name="curr" v-model="inCurr">
                <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                    {{ currency.description }}</option>
            </select>

            <label for="curr"><b>Выходящая валюта:</b></label>
            <select id="curr" name="curr" v-model="outCurr">
                <option v-for="currency in currencies" :value="currency.code">{{ currency.code }} -
                    {{ currency.description }}</option>
            </select>

            <label for="balance"><b>Описание:</b></label>
            <input type="text" placeholder="Описание обменника" name="balance" required v-model="description">

            <label for="inmin"><b>Минимум операции:</b></label>
            <input type="number" placeholder="Описание обменника" name="inmin" required v-model="inmin">

            <input type="checkbox" id="payment_verification" name="payment_verification" v-model="paymentVerify">
            <label for="payment_verification">Требовать подтверждение адреса</label><br>

            <button type="button" class="btn" @click="createExchanger">Создать обменник</button>
            <button type="button" class="btn cancel" @click="closeForm">Закрыть редактор</button>
        </form>
    </div>

</template>

<style scoped>
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
</style>
