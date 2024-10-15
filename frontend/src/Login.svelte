<script>
  let email = "";
  let password = "";
  let message = "";

  async function login() {
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();
    if (response.ok) {
      localStorage.setItem("auth_token", data.token);
      window.location.replace("/");
    } else {
      message = data.error;
    }
  }

  function goToRegister() {
    window.location.href = "#/register";
  }
</script>

<div class="login-container">
  <h2>Login</h2>
  <input type="email" placeholder="Email" bind:value={email} />
  <input type="password" placeholder="Password" bind:value={password} />
  <button on:click={login}>Login</button>
  <button on:click={goToRegister} class="register-button">Register</button>
  <div>{message}</div>
</div>

<style>
  .login-container {
    max-width: 400px;
    margin: 100px auto;
    padding: 20px;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }

  input {
    width: 100%;
    padding: 10px;
    margin: 8px 0;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  button {
    width: 100%;
    padding: 10px;
    margin: 8px 0;
    border: none;
    border-radius: 4px;
    background-color: #6200ea;
    color: white;
    cursor: pointer;
  }

  button:hover {
    background-color: #3700b3;
  }

  .register-button {
    background-color: #03dac5;
    margin-top: 0;
  }

  .register-button:hover {
    background-color: #018786;
  }
</style>
