<script setup>
defineProps({
  msg: {
    type: String,
    required: true
  }
})
</script>

<template>
  <h1>{{ msg }}</h1>
  <h1>Hello world componenent!</h1>
</template>

<style scoped></style>
