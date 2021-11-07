<template>
  <div class="root">
    <h2>{{ title }}</h2>
    <div></div>
    <p>{{ ex.weight }}</p>
    <button @click="$emit('addWeight', ex.id)">+</button>
    <br><br>
    <button @click="$emit('loseWeight', ex.id)">-</button>
    <br><br>
    <button @click="$emit('closeModal', ex.id)">Close</button>
  </div>
</template>

<script setup>
import { defineProps } from "vue";

defineProps({
  title: String,
  msg: String,
  ex: Object,
})

</script>