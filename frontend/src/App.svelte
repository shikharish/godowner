<script>
  import { onMount } from "svelte";
  import Router from "svelte-spa-router";
  import routes from "./routes.js";

  let isAuthenticated = false;

  onMount(() => {
    const token = localStorage.getItem("auth_token");
    isAuthenticated = !!token;

    if (
      !isAuthenticated &&
      window.location.pathname !== "#/login" &&
      !window.location.pathname.includes("register")
    ) {
      window.location.href = "#/login";
    }
  });

  function handleLogout() {
    localStorage.removeItem("auth_token");
    window.location.href = "#/login";
  }
</script>

<body>
  <Router {routes} />
</body>

<button on:click={handleLogout}>Logout</button>
