import { createApp } from 'vue';
import App from './App.vue';
import './index.css';
import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from './views/HomeView.vue';
import AboutView from './views/AboutView.vue';
import ChatView from './views/ChatView.vue';

import '@fortawesome/fontawesome-free/css/all.css';


const routes = [
  { path: "/", component: ChatView },
  { path: "/about", component: AboutView },
  { path: "/chat", component: ChatView },

];
const router = createRouter({
  history: createWebHashHistory(),
  routes
});

createApp(App)
  .use(router)
  .mount('#app');
