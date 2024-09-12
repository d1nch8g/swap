<script>
export default {
    data() {
        return {
            currencies: [],
            exchangers: [],
        }
    },
    mounted() { 
        
    }
}
</script>

<template>
    <h2>Валюты</h2>
    <table id="table">
        <tr>
            <th>ID баланса</th>
            <th>Описание</th>
            <th>Код</th>
            <th>Адрес</th>
            <th>Баланс</th>
            <th>Удалить</th>
        </tr>
        <tr>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
        </tr>
    </table>

    <h2>Обменники</h2>
    <table id="table">
        <tr>
            <th>ID баланса</th>
            <th>Описание</th>
            <th>Код</th>
            <th>Адрес</th>
            <th>Баланс</th>
            <th>Удалить</th>
        </tr>
        <tr>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
            <td>тест</td>
        </tr>
    </table>
</template>

<style>
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