<script>
export default {
    data() {
        return {
            email: "",
            password: "",
            repeatPassword: "",
            showPasswordNotification: false,
            emailSent: false
        }
    },
    methods: {
        async register() {
            if (this.password !== this.repeatPassword) {
                this.showPasswordNotification = true;
                return;
            }

            let headersList = {
                "Content-Type": "application/json"
            }

            let bodyContent = JSON.stringify({
                "email": this.email,
                "password": this.password
            });

            let response = await fetch("/api/create-user", {
                method: "POST",
                body: bodyContent,
                headers: headersList
            });

            if (response.ok) {
                this.emailSent = true;
            }

        }
    }
}
</script>

<template>
    <title>Регистрация нового аккаунта</title>
    <meta property="og:title" content="Регистрация нового аккаунта" />

    <meta property="og:image" content="../assets/logo.svg" />


    <div class="alert" v-if="showPasswordNotification">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Пароли не совпадают
    </div>

    <div class="alert-green" v-if="emailSent">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Уведомление отправлено на почту, перейдите по ссылке что бы подтвердить email
    </div>


    <form>
        <div class="container">
            <h1>Регистрация</h1>
            <p>Заполните поля ниже для создания аккаунта.</p>
            <hr>

            <label for="email"><b>Почта</b></label>
            <input type="email" placeholder="example@email.com" name="email" required v-model="email">

            <label for="psw"><b>Пароль</b></label>
            <input type="password" placeholder="Введите пароль" name="psw" required v-model="password">

            <label for="psw-repeat"><b>Повторите пароль</b></label>
            <input type="password" placeholder="Повторите пароль" name="psw-repeat" required v-model="repeatPassword">

            <p>Создавая аккаунт вы соглашаетесь с правилами использования: <a href="/rules"
                    style="color:dodgerblue">Правила
                    использования</a>.</p>

            <div class="clearfix">
                <button type="submit" class="signupbtn" @click.prevent="register">Регистрация</button>
            </div>
        </div>
    </form>
</template>


<style>
* {
    box-sizing: border-box
}

/* Full-width input fields */
input[type=email],
input[type=password] {
    width: 100%;
    padding: 15px;
    margin: 5px 0 22px 0;
    display: inline-block;
    border: none;
    background: #f1f1f1;
}

input[type=email]:focus,
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