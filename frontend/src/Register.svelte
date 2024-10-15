<script>
  let email = "";
  let password = "";
  let message = "";

  async function register() {
    const response = await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();
    if (response.ok) {
      message = "Registration successful! Redirecting to login...";
      setTimeout(() => {
        window.location.href = "#/login";
      }, 2000);
    } else {
      message = data.error;
    }
  }
</script>

<div class="register-container">
  <h2>Register</h2>
  <input type="email" placeholder="Email" bind:value={email} />
  <input type="password" placeholder="Password" bind:value={password} />
  <button on:click={register}>Register</button>
  <div>{message}</div>
</div>

<style>
  .register-container {
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
</style>
