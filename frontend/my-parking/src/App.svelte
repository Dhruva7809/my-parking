<script lang="ts">
  import Nav from "./components/nav.svelte";
  import Footer from "./components/footer.svelte";

  import Router from "svelte-spa-router";
  import Home from "./routes/Home.svelte";
  import Login from "./routes/Login.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import Book from "./routes/Book.svelte";
  import { onMount } from "svelte";
  import { UserStore } from "./stores/UserStore";
  import { getUser } from "./lib/api";

  onMount(async () => {
    UserStore.set(await getUser());
  });

  let routes = {
    "/": Home,
    "/login": Login,
    "/dashboard": Dashboard,
    "/book": Book,
  };
</script>

<main class="flex flex-col h-screen">
  <Nav />
  <Router {routes} />
  <Footer />
</main>
