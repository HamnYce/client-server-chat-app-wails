<script lang="ts" setup>
import { ref } from 'vue';
import Navbar from './components/Navbar.vue';
let theme = ref('dark');
function changeTheme() {
  // switch between light and dark
  theme.value = theme.value === 'light' ? 'dark' : 'light';
};

// define prop in composition api vue 3
</script>

<template>
  <div :data-theme="theme" class="min-h-screen min-w-screen flex flex-col">
    <div>
      <Navbar class="h-12" v-on:change-theme="changeTheme" />
    </div>
    <div class="grow flex flex-col">
      <RouterView class="grow" />
    </div>
  </div>

</template>
