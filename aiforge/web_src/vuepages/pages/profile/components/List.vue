<template>
  <div>
    <div v-for="(item, index) in params" :key="item.id">
      <Item :data="item" :key="item.ID" />
    </div>
  </div>
</template>

<script>
import Item from './Item.vue';
export default {
  name: "List",
  props: {
    params: { type: Array, default: () => ([]) },
  },
  components: { Item },
  data() {
    return {

    }
  },
}
</script>