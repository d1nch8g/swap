<script>
import { RouterLink, RouterView } from 'vue-router'

export default {
  data() {
    return {
      showLogin: true
    }
  },
  mounted() {
    let token = localStorage.getItem("token");
    if (token) {
      this.showLogin = false;
    }
  },
}
</script>

<template>
  <div class="header">
    <img src="./assets/logo.png" alt="Company logo">

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
      <a v-if="showLogin">
        <RouterLink to="/login">Войти</RouterLink>
      </a>
      <a v-if="showLogin">
        <RouterLink to="/register">Регистрация</RouterLink>
      </a>
      <a v-if="!showLogin">
        <RouterLink to="/orders">Заявки</RouterLink>
      </a>
    </div>
  </div>

  <RouterView />
</template>

<style scoped>
img {
  height: 32px;
  margin: 8px;
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
