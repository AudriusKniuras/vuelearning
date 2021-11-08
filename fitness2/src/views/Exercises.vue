<template>
  <div class="about">
    <h1>Exercises</h1>
    <Exercises :exercises="exercises" />
  </div>
</template>

<script>
import Exercises from "../components/Exercises.vue";

export default {
  name: "ExercisesView",
  components: {
    Exercises,
  },
  data() {
    return {
      exercises: [],
      test: "",
      show: false,
    };
  },
  methods: {
    async fetchTest() {
      const res = await fetch('http://192.168.1.101:10000/allhistory', {
        method: 'GET',
        // mode: 'no-cors'
        // body: JSON.stringify({"key": "Zm9v"})
      })
      const data = await res.json()
      return data
    },
  },
  async created() {
    this.exercises = await this.fetchTest()
  }
};
</script>
