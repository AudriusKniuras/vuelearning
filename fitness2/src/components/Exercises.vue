<template>
  <div class="rectangles">
    <teleport to="body">
      <div class="modal" v-if="isOpen">
        <weight-modal
          :ex="modalObject"
          title="Change weight:"
          @closeModal="closeModal(modalObject.id)"
          @loseWeight="loseWeight(modalObject.id)"
          @addWeight="addWeight(modalObject.id)"
        />
      </div>
    </teleport>
    <div :key="ex.id" v-for="ex in exercises" class="rectangle">
      <div class="exerciseName">
        <h4 class="exerciseName">{{ ex.name }}</h4>
        <button @click="weightToModal(ex)" class="weight">weight: {{ ex.weight }}</button>
      </div>
      <div
        @click="onClick(ex.id - 1, i - 1)"
        :key="i"
        v-for="i in ex.repsArr.length"
        class="circle"
        :style="ex.started[i - 1] ? { background: 'green' } : null">
        <p :style="{ color: 'white' }">{{ ex.repsArr[i - 1] }}</p>
      </div>
    </div>
    <br>
    <button>Hi</button>
  </div>
</template>



<script>
import { ref } from "vue";
import WeightModal from './WeightModal.vue';


export default {
  components: { WeightModal },
  name: "Exercise",
  props: {
    exercises: Array,
  },
  data() {
    return {
      ex_changed: Array,
      modalObject: null,
      isOpen: ref(false),
    };
  },
  methods: {
    onClick(id, circle_id) {
      this.ex_changed = this.exercises;
      if (this.ex_changed[id].started[circle_id] === false) {
        this.ex_changed[id].started[circle_id] = true;
      } else {
        this.ex_changed[id].sets[circle_id] -= 1;
        if (this.ex_changed[id].sets[circle_id] === -1) {
          this.ex_changed[id].started[circle_id] = false;
          this.ex_changed[id].sets[circle_id] = this.ex_changed[id].initialReps;
        }
      }
    },
    weightToModal(exercise) {
      this.isOpen = true
      this.modalObject = exercise
    },
    addWeight(id) {
      this.ex_changed = this.exercises;

      this.ex_changed = this.ex_changed
        .filter((ex) => ex.id === id)
        .map((ex) => ex.weight += 2.5)
    },
    loseWeight(id) {
      this.ex_changed = this.exercises;

      this.ex_changed = this.ex_changed
        .filter((ex) => ex.id === id)
        .map((ex) => ex.weight -= 2.5)
    },
    async fetchExercise(id) {
      const res = await fetch(`api/exercises/${id}`)
      const data = await res.json()
      return data
    },
    async saveExercise(id) {
      const ex = await this.fetchExercise(id)
      const updEx = {...ex, weight: this.exercises.filter((e) => e.id === id)[0].weight}

      const res = await fetch(`api/exercises/${id}`, {
        method: 'PUT',
        headers: {
          'Content-type': 'application/json',
        },
        body: JSON.stringify(updEx)
      })
    },
    closeModal(id) {
      this.isOpen = false
    }
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
  margin-right: 30px;
  display: inline-block;
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
.root {
  position: relative;
}

.modal {
  position: absolute;
  top: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.1);
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal > div {
  background-color: #fff;
  padding: 50px;
  border-radius: 10px;
}
</style>
