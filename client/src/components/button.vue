<template>
  <div :class="classCustom">
    <q-btn @click="triggerGetKey(title)" round :color="bgColor" :label="title" size="lg" />
  </div>
</template>

<script>
export default {
  name: 'Key',
  props: {
    title: {
      type: String,
      required: true
    },
    classCustom: {
      type: String,
      default: 'col-1'
    },
    bgColor: {
      type: String,
      default: 'primary'
    },
    getKey: { type: Function }
  },
  data () {
    return {}
  },
  methods: {
    triggerGetKey (value) {
      console.log(value)
      this.$emit('getKey', value)
    }
  }
}
</script>
