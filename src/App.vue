<script>
import { RouterLink, RouterView } from 'vue-router'
import Chat from './components/Chat.vue'

export default {
  data() {
    return {
      showLogin: true,
      showOperator: false,
      showAdmin: false,
    }
  },
  mounted() {
    let token = localStorage.getItem("token");
    if (token) {
      this.showLogin = false;
      let operator = localStorage.getItem("operator");
      if (operator) {
        this.showOperator = true;
      }
      let admin = localStorage.getItem("admin");
      if (admin) {
        this.showAdmin = true;
      }
    }
  }
}
</script>

<template>
  <div class="header">
    <a href="/">
      <img src="./assets/logo.svg" alt="Company logo" href="/">
    </a>

    <div class="header-right">
      <a>
        <RouterLink to="/">Обмен</RouterLink>
      </a>
      <a>
        <RouterLink to="/contacts">Контакты</RouterLink>
      </a>
      <a>
        <RouterLink to="/rules">Правила</RouterLink>
      </a>
      <a>
        <RouterLink to="/amlkyc">AML/KYC</RouterLink>
      </a>
      <a v-if="showLogin">
        <RouterLink to="/login">Войти</RouterLink>
      </a>
      <a v-if="showLogin">
        <RouterLink to="/register">Регистрация</RouterLink>
      </a>
      <a v-if="!showLogin">
        <RouterLink to="/profile">Профиль</RouterLink>
      </a>
      <a v-if="showOperator">
        <RouterLink to="/operator">Оперирование</RouterLink>
      </a>
      <a v-if="showOperator">
        <RouterLink to="/chats">Чаты</RouterLink>
      </a>
      <a v-if="showOperator">
        <RouterLink to="/orders">Заявки</RouterLink>
      </a>
      <a v-if="showOperator">
        <RouterLink to="/card-confirmations">Подтверждения карт</RouterLink>
      </a>
      <a v-if="showAdmin">
        <RouterLink to="/currencies">Валюты</RouterLink>
      </a>
      <a v-if="showAdmin">
        <RouterLink to="/exchangers">Обменники</RouterLink>
      </a>
    </div>
  </div>

  <RouterView />
</template>

<style scoped>
img {
  height: 32px;
}

.header {
  overflow: hidden;
  background-color: #d4d8d9;
  padding: 10px 5px;
  border-radius: 8px;
}

/* Style the header links */
.header a {
  float: left;
  color: black;
  text-align: center;
  padding: 6px;
  text-decoration: none;
  font-size: 18px;
  line-height: 25px;
  border-radius: 4px;
}

/* Style the logo link (notice that we set the same value of line-height and font-size to prevent the header to increase when the font gets bigger */
.header a.logo {
  font-size: 25px;
  font-weight: bold;
}

/* Change the background color on mouse-over */
.header a:hover {
  background-color: #ddd;
  color: black;
}

/* Style the active/current link*/
.header a.active {
  background-color: dodgerblue;
  color: white;
}

/* Float the link section to the right */
.header-right {
  float: right;
}

/* Add media queries for responsiveness - when the screen is 500px wide or less, stack the links on top of each other */
@media screen and (max-width: 500px) {
  .header a {
    float: none;
    display: block;
    text-align: left;
  }

  .header-right {
    float: none;
  }
}
</style>
