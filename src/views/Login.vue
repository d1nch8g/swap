<script>
export default {
    data() {
        return {
            email: "",
            password: "",
            incorrect: false,
            logged: false
        }
    },
    methods: {
        async login() {
            let headersList = {
                "email": this.email,
                "password": this.password
            }

            let response = await fetch("http://localhost:8080/api/login?Email=support%40inswap.in&Password=password", {
                method: "POST",
                headers: headersList
            });

            if (response.ok) {
                let data = await response.text();
                localStorage.setItem("token", data);
                this.incorrect = false;
                this.logged = true;
                window.location.href = '/';
                return;
            }
            this.incorrect = true;
            this.logged = false;
        }
    }
}
</script>

<template>
    <div class="alert" v-if="incorrect">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Неправильный логин или пароль
    </div>
    <div class="alert-green" v-if="logged">
        <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
        Вы успешно вошли!
    </div>
    <form method="post">
        <div class="container">
            <label for="uname"><b>Почта</b></label>
            <input type="email" placeholder="example@email.com" name="uname" required v-model="email">

            <label for="psw"><b>Пароль</b></label>
            <input type="password" placeholder="Ваш пароль" name="psw" required v-model="password">

            <button type="submit" @click.prevent="login">Войти</button>
        </div>
    </form>
</template>

<style scoped>
/* Full-width inputs */
input[type=email],
input[type=password] {
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    box-sizing: border-box;
}

/* Set a style for all buttons */
button {
    background-color: #d2d6d5;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    cursor: pointer;
    width: 100%;
}

/* Add a hover effect for buttons */
button:hover {
    opacity: 0.8;
}

/* Extra style for the cancel button (red) */
.cancelbtn {
    width: auto;
    padding: 10px 18px;
    background-color: #f44336;
}

/* Center the avatar image inside this container */
.imgcontainer {
    text-align: center;
    margin: 24px 0 12px 0;
}

/* Avatar image */
img.avatar {
    width: 40%;
    border-radius: 50%;
}

/* Add padding to containers */
.container {
    padding: 16px;
}

/* The "Forgot password" text */
span.psw {
    float: right;
    padding-top: 16px;
}

/* Change styles for span and cancel button on extra small screens */
@media screen and (max-width: 300px) {
    span.psw {
        display: block;
        float: none;
    }

    .cancelbtn {
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