<script>
export default {
    methods: {
        logout() {
            localStorage.setItem("token", "");
            window.location.href = '/';
        },
        addCurrency() { }
    },
    async mounted() {
        // list 
    }
}
</script>

<template>
    <h2>Ваши заявки</h2>
    <table id="table">
        <tr>
            <th>id</th>
            <th>Минимальная операция</th>
            <th>Требуется подтверждение платежа</th>
        </tr>
        <tr>
            <td>0</td>
            <td>Maria Anders</td>
            <td>Germany</td>
        </tr>
    </table>
    <button @click="addCurrency">Добавить валюту</button>

    <br>
    <button @click="logout">Выйти</button>
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
</style>