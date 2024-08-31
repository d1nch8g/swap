<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/contacts">Contacts</RouterLink>
      <RouterLink to="/description">Description</RouterLink>
      <RouterLink to="/order">Order</RouterLink>
      <RouterLink to="/transfer">Transfer</RouterLink>
      <RouterLink to="/verify">Verify</RouterLink>
    </nav>
  </header>

  <RouterView />
</template>

<script>
</script>