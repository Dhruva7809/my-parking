<script lang="ts">
  import { UserStore } from "../stores/UserStore";

  import { onMount } from "svelte";

  import { getBookings, getUser } from "../lib/api";

  onMount(async () => {
    UserStore.set(await getUser());
    console.log($UserStore);
    if ($UserStore == null) {
      window.location.href = "/#/login";
    }
  });

  let getbookings = getBookings();
</script>

<section class="text-gray-400 bg-gray-900 body-font flex-grow">
  <div class="container px-5 py-24 mx-auto">
    {#await getbookings then bookings}
      <div class="flex flex-col flex-1 text-center w-full mb-12">
        <h1 class="sm:text-3xl text-2xl font-medium title-font mb-4 text-white">
          {#if bookings.filter((booking) => booking.userid == $UserStore.id.toString()).length > 0}
            Your Bookings
          {:else}
            You have no Bookings :(
            <br />
            <button
              on:click={() => (window.location.href = "/#/book")}
              class="text-white bg-purple-500 border-0 py-2 px-8 mt-5 focus:outline-none hover:bg-purple-600 rounded text-lg"
              >New Booking</button
            >
          {/if}
        </h1>
      </div>
      {#if bookings.length > 0}
        <div class="flex items-center justify-center h-auto p-5">
          <div class="container">
            <div class="flex justify-center">
              <div class="bg-gray-800 shadow-xl rounded-lg w-full">
                <!-- <ul class="flex justify-between divide-x divide-gray-700">
                  <li class="p-4">Booking ID</li>
                  <li class="p-4">Parking Name</li>
                  <li class="p-4">Slot Number</li>
                  <li class="p-4">From</li>
                  <li class="p-4">To</li>
                  <li class="p-4">Rate</li>
                  <li class="p-4">Cost</li>
                </ul>
                <hr class="border-gray-700" /> -->
                <ul class="divide-y divide-gray-700">
                  {#each bookings as booking}
                    {#if booking.userid == $UserStore.id}
                      <li class="">
                        <ul class="flex justify-evenly divide divide-gray-700">
                          <li class="p-4">{booking.id}</li>
                          <li class="p-4">{booking.parkname}</li>
                          <li class="p-4">{booking.slot}</li>
                          <li class="p-4">{booking.from}</li>
                          <li class="p-4">{booking.to}</li>
                          <!-- <li class="p-4">{booking.rate}</li> -->
                          <!-- <li class="p-4">Cost</li> -->
                        </ul>
                      </li>
                    {/if}
                  {/each}
                </ul>
              </div>
            </div>
          </div>
        </div>
      {/if}
    {/await}
  </div>
</section>
