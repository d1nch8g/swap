<template>
    <h2>Входящие заявки</h2>
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

    <h2>Балансы</h2>
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