<script setup lang="ts">
import { ref, onUnmounted } from 'vue';
import { ConnectToChatroom, DisconnectFromChatroom, GetMessages, IsConnected, ListenForMessage, SendMsgToChatRoom } from '../../wailsjs/go/main/App';
import { Message } from '../types';
import ChatBubble from '../components/ChatBubble.vue';

let isConnected = ref(false);
let messages = ref<Message[]>([]);
let userMessage = ref("");
let userName = ref("");

const _interval = setInterval(async () => {
  isConnected.value = await IsConnected();
  if (isConnected.value) {
    await ListenForMessage();
    messages.value = await GetMessages() as Message[];
  }
}, 1000);

function sendMessage() {
  if (userMessage.value === "") {
    return;
  }
  SendMsgToChatRoom(userMessage.value);
  userMessage.value = "";
}

async function connectToServer() {
  if (userName.value === "") {
    alert("Please enter a username");
    return;
  }
  await ConnectToChatroom();
  if (await IsConnected()) {
    SendMsgToChatRoom("SET_NAME:" + userName.value);
  }
}

onUnmounted(() => {
  clearInterval(_interval);
  DisconnectFromChatroom();
});
</script>

<template>
  <div class="flex flex-col">
    <div class="flex flex-col grow">

      <div v-if="isConnected" class="grow">
        <div v-for="message, i in messages" :key="i">
          <ChatBubble :message="message" :username="userName" />
        </div>
      </div>
      <div v-else class="grow flex justify-center items-center">
        <input type="text" v-model="userName" placeholder="username"
          class="input input-bordered w-full max-w-xs">
        <button class="btn btn-primary" @click="connectToServer"> Connect
        </button>
      </div>

      <div v-if="isConnected" class="divider" />

      <div v-if="isConnected" class="flex h-20 px-10 items-center">
        <div>
          <button class="btn btn-error text-xs"
            @click="DisconnectFromChatroom()">Disconnect</button>
        </div>
        <div class="flex justify-between items-center grow">
          <label class="block w-3/4">
            <textarea
              class="textarea w-full items-center resize-none border-none outline-none focus:outline-none"
              style="line-height: 1;" placeholder="Try sending a message!"
              v-model="userMessage" />

          </label>
          <button class="btn btn-primary grow max-w-40 text-lg"
            @click="sendMessage">Send</button>
        </div>
      </div>

    </div>

  </div>
</template>