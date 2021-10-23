<template>
  <div :key="task.id" v-for="task in tasks">
    <!-- emit task.id to upper level  -->
    <Task @delete-task="$emit('delete-task', task.id)" :task="task" />
  </div>
</template>

<script>
import Task from './Task'

export default {
  name: 'Tasks',
  props: {
    tasks: Array
  },
  components: {
    Task,
  },
  emits: ['delete-task']
}
</script>