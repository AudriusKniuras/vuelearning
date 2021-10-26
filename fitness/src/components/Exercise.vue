<template>
  <div class="rectangles">
    <div class="rectangle">
      <div class="exerciseName">
        <h4 class="exerciseName">Squat</h4>
        <p class="weight">5x5 30kg</p>
      </div>
      <div :key="i" v-for="i in reps.length" class="circle"></div>
    </div>
    <div></div>
    <div class="rectangle">
      <div class="exerciseName">
        <h4 class="exerciseName">Overhead Press</h4>
        <p class="weight">5x5 30kg</p>
      </div>
      <div class="circle"></div>
      <div class="circle"></div>
      <div class="circle"></div>
      <div class="circle"></div>
      <div class="circle"></div>
    </div>
    <div></div>
  </div>
</template>

<script>
export default {
  name: "Exercise",
  props: {
    exercises: Array,
  },
};
</script>

<style>
.rectangles .rectangle {
  border-radius: 10px;
  display: inline-block;
  margin-bottom: 30px;
  margin-right: 5px;
  width: 500px;
  height: 100px;
  border: 1px solid #000;
  background-color: white;
  padding: 20px 20px 10px 20px;
  position: relative;
}

.rectangles .rectangle .circle {
  background: rgb(209, 209, 209);
  border-radius: 100%;
  height: 60px;
  width: 60px;
  position: relative;
  top: 10px;
  /* left: -70px; */
  /* float: left; */
  margin-right: 30px;
  display: inline-block;
  /* align-self: auto; */
  /* align-items: stretch; */
}
.exerciseName {
  margin-top: 0px;
  margin-bottom: 0px;
  text-align: left;
  display: flex;
  justify-content: space-between;
}
.weight {
  margin-top: 0px;
  margin-bottom: 0px;
  text-align: right;
}
</style>
