<template>
  <CrowIcon class="crow h-15" />
</template>

<script lang="ts" setup>
import CrowIcon from '~/assets/crow.svg?component';
</script>

<style scoped>
@keyframes peck {
  0% {
    transform: rotate(5deg) translateX(5%);
  }
  10% {
    transform: rotate(-5deg) translateX(-15%);
  }
  20% {
    transform: rotate(5deg) translateX(5%);
  }
  30% {
    transform: rotate(-5deg) translateX(-15%);
  }
  100% {
    transform: rotate(5deg) translateX(5%);
  }
}

.crow ::v-deep(path) {
  animation: peck 1s ease infinite;
  @apply fill-wp-text-100;
}
</style>
