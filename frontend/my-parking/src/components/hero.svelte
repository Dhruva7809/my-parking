<script lang="ts">
  import Car from "../assets/undraw_fast_car.svg";

  function onGetStarted() {
    let Cta = document.getElementById("Cta");
    window.scrollTo({
      top: Cta.offsetTop,
      behavior: "smooth",
    });
  }
</script>

<section class="text-gray-400 bg-gray-900 body-font">
  <div
    class="container mx-auto flex px-5 py-24 md:flex-row flex-col items-center"
  >
    <div class="lg:max-w-lg lg:w-full md:w-1/2 w-5/6 md:mb-0 mb-10">
      <img class="object-cover object-center rounded" alt="hero" src={Car} />
    </div>
    <div
      class="lg:flex-grow md:w-1/2 lg:pl-24 md:pl-16 flex flex-col md:items-start md:text-left items-center text-center"
    >
      <h1 class="title-font sm:text-4xl text-3xl mb-4 font-medium text-white">
        Find safe parking near you
        <br class="hidden lg:inline-block" />At the click of a button
      </h1>
      <p class="mb-8 leading-relaxed">
        We know how precious your car is, so we make sure you don't lose it.
        Find the nearest parking spot and get your car back in no time. We
        provide you with the best parking spots in the city at the most
        affordable rates.
      </p>
      <div class="flex justify-center">
        <button
          on:click={onGetStarted}
          class="inline-flex text-white bg-purple-500 border-0 py-2 px-6 focus:outline-none hover:bg-purple-600 rounded text-lg"
          >Get Started</button
        >
      </div>
    </div>
  </div>
</section>
