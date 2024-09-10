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
        // list executed or processed orders

    }
}
</script>

<template>
    <h2>Ваши заявки</h2>
    <table id="table">
        <tr>
            <th>id</th>
            <th>Exchanger </th>
            <th>Требуется подтверждение платежа</th>
        </tr>
        <tr>
            <td>0</td>
            <td>Maria Anders</td>
            <td>Germany</td>
        </tr>
    </table>

    <br>

    <button class="button" style="vertical-align:middle" @click="logout"><span>Выйти </span></button>
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

.button {
    display: inline-block;
    border-radius: 4px;
    background-color: #f4511e;
    border: none;
    color: #FFFFFF;
    text-align: center;
    font-size: 28px;
    padding: 20px;
    width: 200px;
    transition: all 0.5s;
    cursor: pointer;
    margin: 5px;
}

.button span {
    cursor: pointer;
    display: inline-block;
    position: relative;
    transition: 0.5s;
}

.button span:after {
    content: '\00bb';
    position: absolute;
    opacity: 0;
    top: 0;
    right: -20px;
    transition: 0.5s;
}

.button:hover span {
    padding-right: 25px;
}

.button:hover span:after {
    opacity: 1;
    right: 0;
}
</style>