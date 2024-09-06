<script setup lang="ts">
import { ref } from 'vue';
import { ConnectToChatroom, GetMessages, IsConnected, ListenForMessage, SendMsgToChatRoom } from '../../wailsjs/go/main/App';

let isConnected = ref(false);
const placeholder = "Empty Chat... try to type something!";
let messages = ref<string[]>([]);
let userMessage = ref("");

const _interval = setInterval(async () => {
  isConnected.value = await IsConnected();
  if (isConnected.value) {
    await ListenForMessage();
    console.log(messages);
    messages.value = await GetMessages();
  }
}, 1000);

</script>

<template>
  <div class="flex flex-col">
    <div class="flex flex-col grow">
      <div v-if="isConnected" class="grow">
        {{ messages }}
        <!-- <div v-for="message, i in messages" :key="i">
        <div v-if="message.sender === 'bot'" class="flex justify-start">
          <div class="bg-gray-200 p-2 m-2 rounded-lg chat-bubble">
            {{ message.text }}
          </div>
        </div>
        <div v-else class="flex justify-end">
          <div class="bg-blue-500 text-white p-2 m-2 rounded-lg">
            {{ message.text }}
          </div>
        </div>
      </div> -->
      </div>
      <div v-else class="grow flex justify-center items-center">
        <span>Please
          <button class="btn btn-primary" @click="ConnectToChatroom"> Connect
          </button>
        </span>
      </div>
      <hr class="w-11/12 self-center">
      <div class="flex h-20 justify-between items-center px-10">
        <label class="block w-3/4">
          <textarea
            class="textarea w-full items-center resize-none border-none outline-none focus:outline-none"
            style="line-height: 1;" placeholder="Try sending a message!"
            v-model="userMessage" />

        </label>
        <button class="btn btn-primary grow max-w-40 text-lg"
          @click="SendMsgToChatRoom(userMessage)">Send</button>
      </div>
    </div>

  </div>
</template>