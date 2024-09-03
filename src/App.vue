<script setup>
</script>

<template>
  <header>
    <img class="logo" src="./assets/logo.png" alt="logo">
    <nav>
      <ul class="nav__links">
        <li><a href="/currencies">Currencies</a></li>
        <li><a href="/exchangers">Exchangers</a></li>
        <li><a href="/balances">Balances</a></li>
        <li><a href="/orders">Orders</a></li>
      </ul>
    </nav>
    <a class="cta" href="#"><button>Contact</button></a>
  </header>
</template>

<style scoped>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  background-color: #24252a;
}

li,
a,
button {
  font-weight: 500;
  font-size: 16;
  color: #edf0f1;
  text-decoration: none;
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 30px;
}

.logo {
  cursor: pointer;
  height: 50px;
}

.nav__links {
  list-style: none;
}

.nav__links li {
  display: inline-block;
  padding: 0px 20px;
}

.nav__links li a {
  transition: all 0.3s ease 0s;
}

.nav__links li a:hover {
  color: #0088a9;
}

button {
  padding: 9px 25px;
  background-color: rgba(0, 136, 169, 1);
  border: none;
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.3s ease 0s;
}

button:hover {
  background-color: rgba(0, 136, 169, 0.8);
}
</style>
