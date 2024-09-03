<script setup lang="ts">
import { ref } from 'vue';

const isLightMode = ref(false);
</script>

<template>
  <div class="navbar bg-base-100">
    <div class="flex-1">
      <RouterLink to="/" class="btn btn-ghost text-xl">ChatApp</RouterLink>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal px-1">
        <li>
          <button
            @click="() => { $emit('change-theme'); isLightMode = !isLightMode; }">
            <i v-if="isLightMode" class="fa-solid fa-moon"></i>
            <i v-else class="fa-solid fa-sun"></i>
          </button>
        </li>
        <li>
          <RouterLink to="/">Home</RouterLink>
        </li>
        <li>
          <RouterLink to="/about">About</RouterLink>
        </li>
        <li>
          <RouterLink to="/chat">ChatRoom</RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>