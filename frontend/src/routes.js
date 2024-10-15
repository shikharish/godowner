// routes.js
import Home from "./Home.svelte";
import Login from "./Login.svelte";
import Register from "./Register.svelte"; // Import the Register component

export default {
  "/": Home,
  "/login": Login,
  "/register": Register, // Register route
};
