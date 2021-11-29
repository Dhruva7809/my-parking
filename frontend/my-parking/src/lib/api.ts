export type User = {
    name: string,
    email: string,
    Id: string,
}


export type LoginCreds = {
    name: string;
    email: string;
    password: string;
}

export type Booking = {
    id: string;
    location: string;
    ParkingName: string;
    slot: string;
    userId: string;
    rate: string;
    from: string;
    to: string;
}

const base_api_url = 'http://localhost:8000/api';

export async function registerUser(user: LoginCreds) {
    const response = await fetch(`${base_api_url}/register`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(user),
    });

    const res = await response.json();

    if (res.message === 'success') {
        return true;
    }
    return false;
}

export async function login(user: LoginCreds) {
    const response = await fetch(`${base_api_url}/login`, {
        method: 'POST',
        credentials: 'include',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(user),
    });

    const res = await response.json();

    if (res.message === 'success') {
        return true;
    }
    return false;
}

export async function logout() {
    const response = await fetch(`${base_api_url}/logout`, {
        method: 'POST',
        credentials: 'include',
    });

    const res = await response.json();

    if (res.message === 'success') {
        return true;
    }
    return false;
}

export async function getUser() {
    const response = await fetch(`${base_api_url}/user`, {
        method: 'GET',
        credentials: 'include',
    });

    const res = await response.json();

    if (res.message === 'Unauthenticated') {
        return null;
    }
    return res;
}

export async function getLocations() {
    const response = await fetch(`${base_api_url}/locations`, {
        method: 'GET',
        credentials: 'include',
    });

    const res = await response.json();

    return res;
}

export async function getBookings() {
    const response = await fetch(`${base_api_url}/bookings`, {
        method: 'GET',
        credentials: 'include',
    });

    const res = await response.json();

    return res;
}

export async function addBooking(booking: Booking) {
    const response = await fetch(`${base_api_url}/book`, {
        method: 'POST',
        credentials: 'include',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(booking),
    });

    const res = await response.json();

    if (res.message === 'success') {
        return true;
    }
    return false;
}