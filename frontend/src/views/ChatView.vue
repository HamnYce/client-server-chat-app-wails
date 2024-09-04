<script setup lang="ts">
const placeholder = "Empty Chat... try to type something!";
const messages = [
  {
    id: 1,
    text: "Hello, how can I help you?",
    sender: "bot",
  },
  {
    id: 2,
    text: "I need help with my order",
    sender: "user",
  },
  {
    id: 3,
    text: "Sure, what is your order number?",
    sender: "bot",
  },
];
</script>

<template>
  <div class="flex flex-col">
    <div class="bg-neutral-500 grow">
      TODO: Chat box goes here with v-for the messages
    </div>
    <div class="flex h-20 justify-between items-center">
      <label class="block w-full">
        <textarea
          class="textarea w-full items-center resize-none border-none outline-none focus:outline-none"
          style="line-height: 1;"
          placeholder="Try sending a message!"></textarea>
      </label>
      <button class="btn btn-primary rounded-full">Send</button>
    </div>
  </div>
</template>