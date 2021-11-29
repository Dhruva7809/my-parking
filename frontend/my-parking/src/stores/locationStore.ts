import { writable } from "svelte/store";
import { getLocations } from "../lib/api";

export const locations = writable([]);

export async function fetchLocations() {
    locations.set(await getLocations());
}