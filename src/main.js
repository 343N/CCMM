import "./app.pcss";
import "./styles.css";
import App from "./App.svelte";

const app = new App({
  target: document.getElementById("app"),
});

print = console.log

export default app;
