<template>
  <div class="about">
    <h1>This is an about page</h1>
    <Exercise :reps="[1, 2, 3, 4, 5]" />
  </div>
</template>

<script>
// @ is an alias to /src
import Exercise from "@/components/Exercise.vue";

export default {
  name: "Home",
  components: {
    Exercise,
  },
};
</script>
