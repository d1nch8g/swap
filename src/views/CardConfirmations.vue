<script>
export default {
    data() {
        return {
            cardConfirmations: [],
            email: ""
        }
    },
    methods: {
        async findCardConfirmations() {
            let token = localStorage.getItem("token");

            let headersList = {
                "Authorization": `Bearer ${token}`
            }

            let response = await fetch(`/api/operator/card-confirmations?email=${this.email}`, {
                method: "GET",
                headers: headersList
            });

            if (response.ok) {
                let data = await response.json();
                this.cardConfirmations = data.confirmations;
            }
        }
    }
}
</script>

<template>
    <title>Подтверждения карт</title>

    <h3>Подтвержденные карты</h3>

    <form @submit.prevent="findCardConfirmations">
        <label for="email">Пользовательский email:</label>
        <input type="email" id="email" name="email" v-model="email">

        <input type="submit" value="Найти">
    </form>



    <h2>Найденные карты:</h2>
    <table id="table">
        <tr>
            <th>ID</th>
            <th>Номер карты</th>
            <th>Картинка</th>
        </tr>
        <tr v-for="cardConfirmation in cardConfirmations">
            <td>{{ cardConfirmation.id }}</td>
            <td>{{ cardConfirmation.address }}</td>
            <td><img v-bind:src="'data:image/jpeg;base64,' + cardConfirmation.image" /></td>
        </tr>
    </table>
</template>

<style scoped>
img {
    max-height: 150px;
}

input[type=email],
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

table {
    table-layout: fixed;
}

td {
    overflow: hidden;
    text-overflow: ellipsis;
    word-wrap: break-word;
}

@media only screen and (max-width: 480px) {

    /* horizontal scrollbar for tables if mobile screen */
    table {
        overflow-x: auto;
        display: block;
    }
}
</style>