<script lang="ts">
  import { getUser, login, LoginCreds } from "../lib/api";
  import { hashHex256 } from "../lib/hash256";

  import { onMount } from "svelte";

  import { UserStore } from "../stores/UserStore";

  let invalid = false;

  let email;
  let password;

  onMount(() => {
    email = document.getElementById("email");
    password = document.getElementById("password");
    if ($UserStore != null) {
      window.location.href = "/#/dashboard";
    }
  });

  async function initiateLogin(event) {
    event.preventDefault();

    const user: LoginCreds = {
      name: "",
      email: email.value,
      password: await hashHex256(password.value),
    };

    if (await login(user)) {
      UserStore.set(await getUser());
      invalid = false;
      window.location.href = "/#/dashboard";
    } else {
      invalid = true;
    }
  }
</script>

<section class="text-gray-400 bg-gray-900 body-font flex-grow">
  <div class="container px-5 py-24 mx-auto">
    <div class="flex flex-col flex-1 text-center w-full mb-12">
      <h1 class="sm:text-3xl text-2xl font-medium title-font mb-4 text-white">
        Login to continue to your Dashboard
      </h1>
      {#if invalid}
        <p class="lg:w-2/3 mx-auto leading-relaxed text-base text-red-500">
          Invalid login Credentials
        </p>
      {/if}
    </div>
    <form
      on:submit={initiateLogin}
      class="flex lg:w-2/3 w-full sm:flex-row flex-col mx-auto px-8 sm:px-0 items-end sm:space-x-4 sm:space-y-0 space-y-4"
    >
      <div class="relative sm:mb-0 flex-grow w-full">
        <label for="email" class="leading-7 text-sm text-gray-400">Email</label>
        <input
          required
          type="email"
          id="email"
          name="email"
          class="w-full bg-gray-800 bg-opacity-40 rounded border border-gray-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-900 focus:bg-transparent text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
        />
      </div>
      <div class="relative sm:mb-0 flex-grow w-full">
        <label for="password" class="leading-7 text-sm text-gray-400"
          >Password</label
        >
        <input
          required
          type="password"
          id="password"
          name="password"
          class="w-full bg-gray-800 bg-opacity-40 rounded border border-gray-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-900 focus:bg-transparent text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
        />
      </div>
      <button
        class="text-white bg-purple-500 border-0 py-2 px-8 focus:outline-none hover:bg-purple-600 rounded text-lg"
        >Login</button
      >
    </form>
  </div>
</section>
