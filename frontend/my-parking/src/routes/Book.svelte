<script lang="ts">
  import { onMount } from "svelte";

  import { locations, fetchLocations } from "../stores/locationStore";
  import { UserStore } from "../stores/UserStore";
  import { Booking, addBooking, getUser } from "../lib/api";

  let selectedLocation;
  let selectedParking;
  let selectedSlot;
  let fromTime;
  let toTime;

  let locationNames = [];
  let parkingNames = [];
  let rates = [];

  let locationSelect;
  let parkingSelect;
  let slotSelect;
  let fromTimeSelect;
  let toTimeSelect;

  onMount(async () => {
    UserStore.set(await getUser());
    if ($UserStore == null) {
      window.location.href = "/#/login";
    }

    await fetchLocations();
    $locations.forEach((location) => {
      if (!locationNames.includes(location.location)) {
        locationNames.push(location.location);
      }
    });

    locationSelect = document.getElementById("locationSelect");
    parkingSelect = document.getElementById("parkingSelect");
    slotSelect = document.getElementById("slotSelect");
    fromTimeSelect = document.getElementById("from");
    toTimeSelect = document.getElementById("to");

    locationSelect.addEventListener("change", (e) => {
      selectedLocation = e.target.value;
      addParkings(selectedLocation);
    });

    parkingSelect.addEventListener("change", (e) => {
      selectedParking = e.target.value;
      addSlots(selectedParking);
    });

    slotSelect.addEventListener("change", (e) => {
      selectedSlot = e.target.value;
    });

    fromTimeSelect.addEventListener("change", (e) => {
      fromTime = e.target.value;
    });

    toTimeSelect.addEventListener("change", (e) => {
      toTime = e.target.value;
    });

    addLocations();
    addParkings(selectedLocation);
    addSlots(selectedParking);
  });

  function addLocations() {
    locationNames.forEach((location) => {
      let option = document.createElement("option");
      option.text = location;
      option.value = location;
      locationSelect.appendChild(option);
    });
    selectedLocation = locationNames[0];
  }

  function addParkings(l) {
    parkingNames = [];
    rates = [];
    $locations.forEach((location) => {
      if (location.location == l) {
        parkingNames.push(location.ParkingName);
        rates.push(location.rate);
      }
    });

    parkingNames.forEach((parking) => {
      let option = document.createElement("option");
      option.text =
        parking + " @ " + rates[parkingNames.indexOf(parking)] + " per min";
      option.value = parking;
      parkingSelect.appendChild(option);
    });
    selectedParking = parkingNames[0];
  }

  function addSlots(p) {
    let slots = 0;
    $locations.forEach((location) => {
      if (location.ParkingName == p) {
        slots = location.slots;
      }
    });

    for (let i = 1; i <= slots; i++) {
      let option = document.createElement("option");
      option.text = i.toString();
      option.value = i.toString();
      slotSelect.appendChild(option);
    }
    selectedSlot = slotSelect.value;
  }

  async function bookClicked(event: Event) {
    event.preventDefault();
    let booking: Booking = {
      id: (Math.floor(Math.random() * 1000) + 1).toString(),
      userId: $UserStore.id.toString(),
      location: selectedLocation,
      ParkingName: selectedParking,
      slot: selectedSlot.toString(),
      rate: rates[parkingNames.indexOf(selectedParking)],
      from: fromTime,
      to: toTime,
    };

    if (await addBooking(booking)) {
      alert("Booking Successful");
    } else {
      alert("Booking Failed");
    }
  }

  function getCurrentDateTime() {
    const date = new Date(Date.now());
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hours = date.getHours();
    const minutes = date.getMinutes();
    // const seconds = date.getSeconds();
    const currentDateTime = `${year}-${month}-${day}T${hours}:${
      minutes < 10 ? "0" + minutes : minutes
    }`;
    console.log(currentDateTime);
    return currentDateTime;
  }
</script>

<section class="text-gray-400 bg-gray-900 body-font flex-grow">
  <!-- Required form plugin -->
  <link
    href="https://cdn.jsdelivr.net/npm/@tailwindcss/custom-forms@0.2.1/dist/custom-forms.css"
    rel="stylesheet"
  />
  <form class="flex justify-center flex-grow pt-10" on:submit={bookClicked}>
    <div
      class="p-4 shadow-xl bg-gray-800 rounded-md text-left w-full"
      style="max-width: 400px"
    >
      <label class="block mt-4">
        <span class="">Select Location</span>
        <select
          class="form-select mt-1 block w-full"
          id="locationSelect"
          required
        />
      </label>
      <div class="flex w-full justify-between">
        <label class="block mt-4 w-3/4">
          <span class="">Select Parking</span>
          <select
            class="form-select mt-1 block w-full"
            id="parkingSelect"
            required
          />
        </label>
        <label class="block mt-4">
          <span class="">Select Slot</span>
          <select
            id="slotSelect"
            class="form-select mt-1 block w-full"
            required
          />
        </label>
      </div>
      <label class="block mt-4">
        <span class="">From: </span>
        <input
          required
          class="form-select mt-1 block w-full"
          type="datetime-local"
          id="from"
          name="from"
          min={getCurrentDateTime()}
          max={new Date(
            new Date().setDate(new Date().getDate() + 1)
          ).toISOString()}
        />
      </label>
      <label class="block mt-4">
        <span class="">To: </span>
        <input
          required
          class="form-select mt-1 block w-full"
          type="datetime-local"
          id="to"
          name="to"
          min={getCurrentDateTime()}
          max={new Date(
            new Date().setDate(new Date().getDate() + 1)
          ).toISOString()}
        />
      </label>
      <div class="flex mt-6 justify-center">
        <label class="flex items-center">
          <button
            class="text-white bg-purple-500 border-0 py-2 px-8 focus:outline-none hover:bg-purple-600 rounded text-lg"
            >Book</button
          >
        </label>
      </div>
    </div>
  </form>
</section>
