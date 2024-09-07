<script setup lang="ts">
import { Message } from '../types';
defineProps(
  {
    message: {
      type: Object as () => Message,
      required: true,
    },
    username: {
      type: String,
      required: true,
    },
  }
);
</script>

<template>
  <div class="chat"
    :class="message.Sender === username ? 'chat-end' : 'chat-start'">
    <div class="chat-header">
      {{ message.Sender }}
      <time class="text-xs opacity-50">{{ message.Time }}</time>
    </div>
    <div class="chat-bubble">{{ message.Message }}</div>
  </div>
</template>