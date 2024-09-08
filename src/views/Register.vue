<script>
</script>

<template>
    <form action="action_page.php" style="border:1px solid #ccc">
        <div class="container">
            <h1>Регистрация</h1>
            <p>Заполните поля ниже для создания аккаунта.</p>
            <hr>

            <label for="email"><b>Почта</b></label>
            <input type="text" placeholder="example@email.com" name="email" required>

            <label for="psw"><b>Пароль</b></label>
            <input type="password" placeholder="Введите пароль" name="psw" required>

            <label for="psw-repeat"><b>Повторите пароль</b></label>
            <input type="password" placeholder="Повторите пароль" name="psw-repeat" required>

            <p>Создавая аккаунт вы соглашаетесь с правилами использования: <a href="/rules"
                    style="color:dodgerblue">Правила
                    использования</a>.</p>

            <div class="clearfix">
                <button type="submit" class="signupbtn">Регистрация</button>
            </div>
        </div>
    </form>
</template>


<style>
* {
    box-sizing: border-box
}

/* Full-width input fields */
input[type=text],
input[type=password] {
    width: 100%;
    padding: 15px;
    margin: 5px 0 22px 0;
    display: inline-block;
    border: none;
    background: #f1f1f1;
}

input[type=text]:focus,
input[type=password]:focus {
    background-color: #ddd;
    outline: none;
}

hr {
    border: 1px solid #f1f1f1;
    margin-bottom: 25px;
}

/* Set a style for all buttons */
button {
    background-color: #bfc0c0;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    cursor: pointer;
    width: 100%;
    opacity: 0.9;
}

button:hover {
    opacity: 1;
}

/* Extra styles for the cancel button */
.cancelbtn {
    padding: 14px 20px;
    background-color: #f44336;
}

/* Float cancel and signup buttons and add an equal width */
.cancelbtn,
.signupbtn {
    float: left;
    width: 100%;
}

/* Add padding to container elements */
.container {
    padding: 16px;
}

/* Clear floats */
.clearfix::after {
    content: "";
    clear: both;
    display: table;
}

/* Change styles for cancel button and signup button on extra small screens */
@media screen and (max-width: 300px) {

    .cancelbtn,
    .signupbtn {
        width: 100%;
    }
}
</style>