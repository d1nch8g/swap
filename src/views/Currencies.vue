<script>
export default {
    data() {
        return {
            currencies: [],
            code: "",
            description: "",
        }
    },
    async mounted() {
        let response = await fetch("/api/list-currencies", {
            method: "GET"
        });

        let data = await response.json();
        this.currencies = data.currencies;
    },
    methods: {
        openForm() {
            document.getElementById("myForm").style.display = "block";
        },
        closeForm() {
            document.getElementById("myForm").style.display = "none";
        },
        async removeCurrency(code) {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "code": code
            });

            let response = await fetch("/api/admin/remove-currency", {
                method: "DELETE",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/currencies";
            }
        },
        async createCurrency() {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "code": this.code,
                "description": this.description
            });

            let response = await fetch("/api/admin/create-currency", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                window.location.href = "/currencies";
            }
        }
    }
}
</script>

<template>
    <title>Валюты</title>

    <h2>Валюты</h2>
    <table id="table">
        <tr>
            <th>ID</th>
            <th>Название</th>
            <th>Описание</th>
            <th>Удаление</th>
        </tr>
        <tr v-for="currency in currencies">
            <td>{{ currency.id }}</td>
            <td>{{ currency.code }}</td>
            <td>{{ currency.description }}</td>
            <td><button @click="removeCurrency(currency.code)">Удалить</button></td>
        </tr>
    </table>

    <button class="open-button" @click="openForm">Добавить валюту</button>
    <!-- The form -->
    <div class="form-popup" id="myForm">
        <form class="form-container">
            <h3>Добавить валюту</h3>

            <label for="address"><b>Код:</b></label>
            <input type="text" placeholder="Код валюты, например BTC" name="address" required v-model="code">

            <label for="balance"><b>Описание:</b></label>
            <input type="text" placeholder="Описание валюты" name="balance" required v-model="description">

            <button type="button" class="btn" @click="createCurrency">Создать валюту</button>
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