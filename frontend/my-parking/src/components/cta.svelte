<script lang="ts">
  import { onMount } from "svelte";

  import { hashHex256 } from "../lib/hash256";
  import { LoginCreds, registerUser } from "../lib/api";

  let error = false;

  let name;
  let email;
  let password;

  onMount(() => {
    name = document.getElementById("full-name");
    email = document.getElementById("email");
    password = document.getElementById("password");
  });

  async function onSignUp(event: Event) {
    event.preventDefault();

    const user: LoginCreds = {
      name: name.value,
      email: email.value,
      password: await hashHex256(password.value),
    };

    if (await registerUser(user)) {
      error = false;
      window.location.href = "/#/login";
    } else {
      error = true;
    }
  }
</script>

<section class="text-gray-400 bg-gray-900 body-font">
  <div class="container px-5 py-24 mx-auto flex flex-wrap items-center">
    <div class="lg:w-3/5 md:w-1/2 md:pr-16 lg:pr-0 pr-0">
      <h1 class="title-font font-medium text-3xl text-white">
        Your automated future is here.
      </h1>
      <p class="leading-relaxed mt-4">
        Sign up to get started with the easiest way to park your vehicles.
        <br />
        We keep your data safe and secure.
      </p>
    </div>
    <form
      on:submit={onSignUp}
      action=""
      id="Cta"
      class="lg:w-2/6 md:w-1/2 bg-gray-800 bg-opacity-50 rounded-lg p-8 flex flex-col md:ml-auto w-full mt-10 md:mt-0"
    >
      <h2 class="text-white text-lg font-medium title-font mb-5">Sign Up</h2>
      <div class="relative mb-4">
        <label for="full-name" class="leading-7 text-sm text-gray-400"
          >Name</label
        >
        <input
          type="text"
          id="full-name"
          name="full-name"
          required
          class="w-full bg-gray-600 bg-opacity-20 focus:bg-transparent focus:ring-2 focus:ring-purple-900 rounded border border-gray-600 focus:border-purple-500 text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
        />
      </div>
      <div class="relative mb-4">
        <label for="email" class="leading-7 text-sm text-gray-400">Email</label>
        <input
          type="email"
          id="email"
          name="email"
          required
          class="w-full bg-gray-600 bg-opacity-20 focus:bg-transparent focus:ring-2 focus:ring-purple-900 rounded border border-gray-600 focus:border-purple-500 text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
        />
      </div>
      <div class="relative mb-4">
        <label for="full-name" class="leading-7 text-sm text-gray-400"
          >Password</label
        >
        <input
          type="password"
          id="password"
          name="password"
          minlength="8"
          required
          class="w-full bg-gray-600 bg-opacity-20 focus:bg-transparent focus:ring-2 focus:ring-purple-900 rounded border border-gray-600 focus:border-purple-500 text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
        />
      </div>
      <button
        type="submit"
        class="text-white bg-purple-500 border-0 py-2 px-8 focus:outline-none hover:bg-purple-600 rounded text-lg"
        >Sign Up</button
      >
      {#if error}
        <p class="text-xs mt-3 text-red-500">Email already in use.</p>
      {/if}
    </form>
  </div>
</section>
