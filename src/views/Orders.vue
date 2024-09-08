<script>
export default {
    methods: {
        logout() {
            localStorage.setItem("token", "");
            window.location.href = '/';
        }
    }
}
</script>

<template>
    <button @click="logout">Выйти</button>
</template>

<style scoped>
.button {
    transition-duration: 0.4s;
}

.button:hover {
    background-color: #04AA6D;
    /* Green */
    color: white;
}
</style>