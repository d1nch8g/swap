import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import Chat from "./components/Chat.vue";

const app = createApp(App);

app.use(router);

app.mount("#app");
app.component("Chat", Chat);
